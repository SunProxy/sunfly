package block

import (
	"github.com/go-gl/mathgl/mgl64"
	"github.com/sunproxy/sunfly/dragonfly/item"
	"github.com/sunproxy/sunfly/dragonfly/world"
	"math/rand"
)

// MelonSeeds grow melon blocks.
type MelonSeeds struct {
	crop

	// direction is the direction from the stem to the melon.
	direction world.Face
}

// SameCrop ...
func (MelonSeeds) SameCrop(c Crop) bool {
	_, ok := c.(MelonSeeds)
	return ok
}

// NeighbourUpdateTick ...
func (m MelonSeeds) NeighbourUpdateTick(pos, _ world.BlockPos, w *world.World) {
	if _, ok := w.Block(pos.Side(world.FaceDown)).(Farmland); !ok {
		w.BreakBlock(pos)
	} else if m.direction != world.FaceDown {
		if _, ok := w.Block(pos.Side(m.direction)).(Melon); !ok {
			m.direction = world.FaceDown
			w.PlaceBlock(pos, m)
		}
	}
}

// RandomTick ...
func (m MelonSeeds) RandomTick(pos world.BlockPos, w *world.World, r *rand.Rand) {
	if rand.Float64() <= m.CalculateGrowthChance(pos, w) && w.Light(pos) >= 8 {
		if m.Growth < 7 {
			m.Growth++
			w.PlaceBlock(pos, m)
		} else {
			directions := world.AllDirections()
			for _, i := range directions {
				if _, ok := w.Block(pos.Side(i.Face())).(Melon); ok {
					return
				}
			}
			direction := directions[rand.Intn(len(directions))].Face()
			stemPos := pos.Side(direction)
			if _, ok := w.Block(stemPos).(Air); ok {
				switch w.Block(stemPos.Side(world.FaceDown)).(type) {
				case Farmland:
				case Dirt:
				case Grass:
					m.direction = direction
					w.PlaceBlock(pos, m)
					w.PlaceBlock(stemPos, Melon{})
				}
			}
		}
	}
}

// BoneMeal ...
func (m MelonSeeds) BoneMeal(pos world.BlockPos, w *world.World) bool {
	if m.Growth == 7 {
		return false
	}
	m.Growth = min(m.Growth+rand.Intn(4)+2, 7)
	w.PlaceBlock(pos, m)
	return true
}

// UseOnBlock ...
func (m MelonSeeds) UseOnBlock(pos world.BlockPos, face world.Face, _ mgl64.Vec3, w *world.World, user item.User, ctx *item.UseContext) bool {
	pos, _, used := firstReplaceable(w, pos, face, m)
	if !used {
		return false
	}

	if _, ok := w.Block(pos.Side(world.FaceDown)).(Farmland); !ok {
		return false
	}

	place(w, pos, m, user, ctx)
	return placed(ctx)
}

// BreakInfo ...
func (m MelonSeeds) BreakInfo() BreakInfo {
	return BreakInfo{
		Hardness:    0,
		Harvestable: alwaysHarvestable,
		Effective:   nothingEffective,
		Drops:       simpleDrops(item.NewStack(m, 1)),
	}
}

// EncodeItem ...
func (m MelonSeeds) EncodeItem() (id int32, meta int16) {
	return 362, 0
}

// EncodeBlock ...
func (m MelonSeeds) EncodeBlock() (name string, properties map[string]interface{}) {
	return "minecraft:melon_stem", map[string]interface{}{"growth": int32(m.Growth)}
}

// Hash ...
func (m MelonSeeds) Hash() uint64 {
	return hashMelonStem | (uint64(m.Growth) << 32)
}
