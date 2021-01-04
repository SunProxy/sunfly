package mcdb

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/df-mc/goleveldb/leveldb"
	"github.com/df-mc/goleveldb/leveldb/opt"
	"github.com/sandertv/gophertunnel/minecraft/nbt"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sunproxy/sunfly/dragonfly/world"
	"github.com/sunproxy/sunfly/dragonfly/world/chunk"
	"github.com/sunproxy/sunfly/dragonfly/world/difficulty"
	"github.com/sunproxy/sunfly/dragonfly/world/gamemode"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

// Provider implements a world provider for the Minecraft world format, which is based on a leveldb database.
type Provider struct {
	db  *leveldb.DB
	dir string
	d   data
}

// chunkVersion is the current version of chunks.
const chunkVersion = 19

// New creates a new provider reading and writing files to files under the path passed. If a world is present
// at the path, New will parse its data and initialise the world with it. If the data cannot be parsed, an
// error is returned.
func New(dir string) (*Provider, error) {
	_ = os.MkdirAll(filepath.Join(dir, "db"), 0777)

	p := &Provider{dir: dir}
	if _, err := os.Stat(filepath.Join(dir, "level.dat")); os.IsNotExist(err) {
		// A level.dat was not currently present for the world.
		p.initDefaultLevelDat()
	} else {
		f, err := ioutil.ReadFile(filepath.Join(dir, "level.dat"))
		if err != nil {
			return nil, fmt.Errorf("error opening level.dat file: %w", err)
		}
		// The first 8 bytes are a useless header (version and length): We don't need it.
		if len(f) < 8 {
			// The file did not have enough content, meaning it is corrupted. We return an error.
			return nil, fmt.Errorf("level.dat exists but has no data")
		}
		if err := nbt.UnmarshalEncoding(f[8:], &p.d, nbt.LittleEndian); err != nil {
			return nil, fmt.Errorf("error decoding level.dat NBT: %w", err)
		}
		p.d.WorldStartCount++
	}
	db, err := leveldb.OpenFile(filepath.Join(dir, "db"), &opt.Options{
		Compression: opt.FlateCompression,
		BlockSize:   16 * opt.KiB,
	})
	if err != nil {
		return nil, fmt.Errorf("error opening leveldb database: %w", err)
	}
	p.db = db
	return p, nil
}

// initDefaultLevelDat initialises a default level.dat file.
func (p *Provider) initDefaultLevelDat() {
	p.d.DoDayLightCycle = true
	p.d.BaseGameVersion = protocol.CurrentVersion
	p.d.LevelName = "World"
	p.d.SpawnY = 128
	p.d.GameType = 1
	p.d.StorageVersion = 8
	p.d.Generator = 1
	p.d.NetworkVersion = protocol.CurrentProtocol
	p.d.Abilities.WalkSpeed = 0.1
	p.d.PVP = true
	p.d.WorldStartCount = 1
	p.d.RandomTickSpeed = 1
	p.d.FallDamage = true
	p.d.FireDamage = true
	p.d.DrowningDamage = true
	p.d.CommandsEnabled = true
	p.d.MultiPlayerGame = true
}

// LoadTime returns the time as it was stored in the level.dat of the world loaded.
func (p *Provider) LoadTime() int64 {
	return p.d.Time
}

// SaveTime saves the time to the level.dat of the world.
func (p *Provider) SaveTime(time int64) {
	p.d.Time = time
}

// LoadTimeCycle returns whether the time is cycling or not.
func (p *Provider) LoadTimeCycle() bool {
	return p.d.DoDayLightCycle
}

// SaveTimeCycle saves the state of the time cycle, either running or stopped, to the level.dat.
func (p *Provider) SaveTimeCycle(running bool) {
	p.d.DoDayLightCycle = running
}

// WorldName returns the name of the world that the provider provides data for.
func (p *Provider) WorldName() string {
	return p.d.LevelName
}

// SetWorldName sets the name of the world to the string passed.
func (p *Provider) SetWorldName(name string) {
	p.d.LevelName = name
}

// WorldSpawn returns the spawn of the world as present in the level.dat.
func (p *Provider) WorldSpawn() world.BlockPos {
	y := p.d.SpawnY
	if p.d.SpawnY > 256 {
		// TODO: Spawn at the highest block of the world. We're currently doing a guess.
		y = 90
	}
	return world.BlockPos{int(p.d.SpawnX), int(y), int(p.d.SpawnZ)}
}

// SetWorldSpawn sets the spawn of the world to a new one.
func (p *Provider) SetWorldSpawn(pos world.BlockPos) {
	p.d.SpawnX, p.d.SpawnY, p.d.SpawnZ = int32(pos.X()), int32(pos.Y()), int32(pos.Z())
}

// LoadChunk loads a chunk at the position passed from the leveldb database. If it doesn't exist, exists is
// false. If an error is returned, exists is always assumed to be true.
func (p *Provider) LoadChunk(position world.ChunkPos) (c *chunk.Chunk, exists bool, err error) {
	data := chunk.SerialisedData{}
	key := index(position)

	// This key is where the version of a chunk resides. The chunk version has changed many times, without any
	// actual substantial changes, so we don't check this.
	_, err = p.db.Get(append(key, keyVersion), nil)
	if err == leveldb.ErrNotFound {
		return nil, false, nil
	} else if err != nil {
		return nil, true, fmt.Errorf("error reading version: %w", err)
	}

	data.Data2D, err = p.db.Get(append(key, key2DData), nil)
	if err == leveldb.ErrNotFound {
		return nil, false, nil
	} else if err != nil {
		return nil, true, fmt.Errorf("error reading 2D data: %w", err)
	}

	data.BlockNBT, err = p.db.Get(append(key, keyBlockEntities), nil)
	// Block entities aren't present when there aren't any, so it's okay if we can't find the key.
	if err != nil && err != leveldb.ErrNotFound {
		return nil, true, fmt.Errorf("error reading block entities: %w", err)
	}

	for y := byte(0); y < 16; y++ {
		data.SubChunks[y], err = p.db.Get(append(key, keySubChunkData, y), nil)
		if err == leveldb.ErrNotFound {
			// No sub chunk present at this Y level. We skip this one and move to the next, which might still
			// be present.
			continue
		} else if err != nil {
			return nil, true, fmt.Errorf("error reading 2D sub chunk %v: %w", y, err)
		}
	}
	c, err = chunk.DiskDecode(data)
	return c, true, err
}

// SaveChunk saves a chunk at the position passed to the leveldb database. Its version is written as the
// version in the chunkVersion constant.
func (p *Provider) SaveChunk(position world.ChunkPos, c *chunk.Chunk) error {
	data := chunk.DiskEncode(c, false)

	key := index(position)
	_ = p.db.Put(append(key, keyVersion), []byte{chunkVersion}, nil)
	_ = p.db.Put(append(key, key2DData), data.Data2D, nil)

	finalisation := make([]byte, 4)
	binary.LittleEndian.PutUint32(finalisation, 2)
	_ = p.db.Put(append(key, keyFinalisation), finalisation, nil)

	if len(data.BlockNBT) != 0 {
		// We only write block NBT if there actually is any.
		_ = p.db.Put(append(key, keyBlockEntities), data.BlockNBT, nil)
	}
	for y, sub := range data.SubChunks {
		if len(sub) == 0 {
			// No sub chunk here: Delete it from the database and continue.
			_ = p.db.Delete(append(key, keySubChunkData, byte(y)), nil)
			continue
		}
		_ = p.db.Put(append(key, keySubChunkData, byte(y)), sub, nil)
	}
	return nil
}

// LoadDefaultGameMode returns the default game mode stored in the level.dat.
func (p *Provider) LoadDefaultGameMode() gamemode.GameMode {
	switch p.d.GameType {
	default:
		return gamemode.Adventure{}
	case 0:
		return gamemode.Survival{}
	case 1:
		return gamemode.Creative{}
	case 2:
		return gamemode.Adventure{}
	case 3:
		return gamemode.Spectator{}
	}
}

// SaveDefaultGameMode changes the default game mode in the level.dat.
func (p *Provider) SaveDefaultGameMode(mode gamemode.GameMode) {
	switch mode.(type) {
	case gamemode.Survival:
		p.d.GameType = 0
	case gamemode.Creative:
		p.d.GameType = 1
	case gamemode.Adventure:
		p.d.GameType = 2
	case gamemode.Spectator:
		p.d.GameType = 3
	}
}

// LoadDifficulty loads the difficulty stored in the level.dat.
func (p *Provider) LoadDifficulty() difficulty.Difficulty {
	switch p.d.Difficulty {
	default:
		return difficulty.Normal{}
	case 0:
		return difficulty.Peaceful{}
	case 1:
		return difficulty.Easy{}
	case 3:
		return difficulty.Hard{}
	}
}

// SaveDifficulty saves the difficulty passed to the level.dat.
func (p *Provider) SaveDifficulty(d difficulty.Difficulty) {
	switch d.(type) {
	case difficulty.Peaceful:
		p.d.Difficulty = 0
	case difficulty.Easy:
		p.d.Difficulty = 1
	case difficulty.Normal:
		p.d.Difficulty = 2
	case difficulty.Hard:
		p.d.Difficulty = 3
	}
}

// LoadEntities loads all entities from the chunk position passed.
func (p *Provider) LoadEntities(world.ChunkPos) ([]world.Entity, error) {
	// TODO: Implement entities.
	return nil, nil
}

// SaveEntities saves all entities to the chunk position passed.
func (p *Provider) SaveEntities(world.ChunkPos, []world.Entity) error {
	// TODO: Implement entities.
	return nil
}

// LoadBlockNBT loads all block entities from the chunk position passed.
func (p *Provider) LoadBlockNBT(position world.ChunkPos) ([]map[string]interface{}, error) {
	data, err := p.db.Get(append(index(position), keyBlockEntities), nil)
	if err != leveldb.ErrNotFound && err != nil {
		return nil, err
	}
	var a []map[string]interface{}

	buf := bytes.NewBuffer(data)
	dec := nbt.NewDecoderWithEncoding(buf, nbt.LittleEndian)

	for buf.Len() != 0 {
		var m map[string]interface{}
		if err := dec.Decode(&m); err != nil {
			return nil, fmt.Errorf("error decoding block NBT: %w", err)
		}
		a = append(a, m)
	}
	return a, nil
}

// SaveBlockNBT saves all block NBT data to the chunk position passed.
func (p *Provider) SaveBlockNBT(position world.ChunkPos, data []map[string]interface{}) error {
	if len(data) == 0 {
		return p.db.Delete(append(index(position), keyBlockEntities), nil)
	}
	buf := bytes.NewBuffer(nil)
	enc := nbt.NewEncoderWithEncoding(buf, nbt.LittleEndian)
	for _, d := range data {
		if err := enc.Encode(d); err != nil {
			return fmt.Errorf("error encoding block NBT: %w", err)
		}
	}
	return p.db.Put(append(index(position), keyBlockEntities), buf.Bytes(), nil)
}

// Close closes the provider, saving any file that might need to be saved, such as the level.dat.
func (p *Provider) Close() error {
	p.d.LastPlayed = time.Now().Unix()

	f, err := os.OpenFile(filepath.Join(p.dir, "level.dat"), os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("error opening level.dat file: %w", err)
	}

	buf := bytes.NewBuffer(nil)
	_ = binary.Write(buf, binary.LittleEndian, int32(3))
	nbtData, err := nbt.MarshalEncoding(p.d, nbt.LittleEndian)
	if err != nil {
		return fmt.Errorf("error encoding level.dat to NBT: %w", err)
	}
	_ = binary.Write(buf, binary.LittleEndian, int32(len(nbtData)))
	_, _ = buf.Write(nbtData)

	_, _ = f.Write(buf.Bytes())

	if err := f.Close(); err != nil {
		return fmt.Errorf("error closing level.dat: %w", err)
	}
	//noinspection SpellCheckingInspection
	if err := ioutil.WriteFile(filepath.Join(p.dir, "levelname.txt"), []byte(p.d.LevelName), 0644); err != nil {
		return fmt.Errorf("error writing levelname.txt: %w", err)
	}
	return p.db.Close()
}

// index returns a byte buffer holding the written index of the chunk position passed.
func index(position world.ChunkPos) []byte {
	x, z := uint32(position[0]), uint32(position[1])
	return []byte{
		byte(x), byte(x >> 8), byte(x >> 16), byte(x >> 24),
		byte(z), byte(z >> 8), byte(z >> 16), byte(z >> 24),
	}
}
