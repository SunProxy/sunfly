package block

import (
	"github.com/sunproxy/sunfly/dragonfly/entity"
	"github.com/sunproxy/sunfly/dragonfly/entity/physics"
	"github.com/sunproxy/sunfly/dragonfly/event"
	"github.com/sunproxy/sunfly/dragonfly/internal/block_internal"
	"github.com/sunproxy/sunfly/dragonfly/world"
	"github.com/sunproxy/sunfly/dragonfly/world/sound"
	"math/rand"
	"time"
)

// Lava is a light-emitting fluid block that causes fire damage.
type Lava struct {
	noNBT
	empty
	replaceable

	// Still makes the lava not spread whenever it is updated. Still lava cannot be acquired in the game
	// without world editing.
	Still bool
	// Depth is the depth of the water. This is a number from 1-8, where 8 is a source block and 1 is the
	// smallest possible lava block.
	Depth int
	// Falling specifies if the lava is falling. Falling lava will always appear as a source block, but its
	// behaviour differs when it starts spreading.
	Falling bool
}

// neighboursLavaFlammable returns true if one a block adjacent to the passed position is flammable.
func neighboursLavaFlammable(pos world.BlockPos, w *world.World) bool {
	for i := world.Face(0); i < 6; i++ {
		if flammable, ok := w.Block(pos.Side(i)).(Flammable); ok && flammable.FlammabilityInfo().LavaFlammable {
			return true
		}
	}
	return false
}

// EntityCollide ...
func (l Lava) EntityCollide(e world.Entity) {
	if flammable, ok := e.(entity.Flammable); ok {
		block_internal.LavaDamage(e, 4)
		flammable.SetOnFire(time.Duration(15) * time.Second)
	}
}

// RandomTick ...
func (l Lava) RandomTick(pos world.BlockPos, w *world.World, r *rand.Rand) {
	i := r.Intn(3)
	if i > 0 {
		for j := 0; j < i; j++ {
			pos = pos.Add(world.BlockPos{r.Intn(3) - 1, 1, r.Intn(3) - 1})
			if _, ok := w.Block(pos).(Air); ok {
				if neighboursLavaFlammable(pos, w) {
					w.PlaceBlock(pos, Fire{})
				}
			}
		}
	} else {
		for j := 0; j < 3; j++ {
			pos = pos.Add(world.BlockPos{r.Intn(3) - 1, 0, r.Intn(3) - 1})
			if _, ok := w.Block(pos.Side(world.FaceUp)).(Air); ok {
				if flammable, ok := w.Block(pos).(Flammable); ok && flammable.FlammabilityInfo().LavaFlammable && flammable.FlammabilityInfo().Encouragement > 0 {
					w.PlaceBlock(pos, Fire{})
				}
			}
		}
	}
}

// AABB returns no boxes.
func (Lava) AABB(world.BlockPos, *world.World) []physics.AABB {
	return nil
}

// HasLiquidDrops ...
func (Lava) HasLiquidDrops() bool {
	return false
}

// LightDiffusionLevel always returns 2.
func (Lava) LightDiffusionLevel() uint8 {
	return 2
}

// LightEmissionLevel returns 15.
func (Lava) LightEmissionLevel() uint8 {
	return 15
}

// NeighbourUpdateTick ...
func (l Lava) NeighbourUpdateTick(pos, _ world.BlockPos, w *world.World) {
	if !l.Harden(pos, w, nil) {
		w.ScheduleBlockUpdate(pos, time.Second*3/2)
	}
}

// ScheduledTick ...
func (l Lava) ScheduledTick(pos world.BlockPos, w *world.World) {
	if !l.Harden(pos, w, nil) {
		tickLiquid(l, pos, w)
	}
}

// LiquidDepth returns the depth of the lava.
func (l Lava) LiquidDepth() int {
	return l.Depth
}

// SpreadDecay always returns 2.
func (Lava) SpreadDecay() int {
	return 2
}

// WithDepth returns a new Lava block with the depth passed and falling if set to true.
func (l Lava) WithDepth(depth int, falling bool) world.Liquid {
	l.Depth = depth
	l.Falling = falling
	l.Still = false
	return l
}

// LiquidFalling checks if the lava is falling.
func (l Lava) LiquidFalling() bool {
	return l.Falling
}

// LiquidType returns "lava" as a unique identifier for the lava liquid.
func (Lava) LiquidType() string {
	return "lava"
}

// Harden handles the hardening logic of lava.
func (l Lava) Harden(pos world.BlockPos, w *world.World, flownIntoBy *world.BlockPos) bool {
	var ok bool
	var water, b world.Block

	if flownIntoBy == nil {
		var water, b world.Block
		_, soulSoilFound := w.Block(pos.Side(world.FaceDown)).(SoulSoil)
		pos.Neighbours(func(neighbour world.BlockPos) {
			if b != nil || neighbour[1] == pos[1]-1 {
				return
			}
			if _, ok := w.Block(neighbour).(BlueIce); ok {
				if soulSoilFound {
					b = Basalt{}
				}
				return
			}
			if waterBlock, ok := w.Block(neighbour).(Water); ok {
				water = waterBlock
				if l.Depth == 8 && !l.Falling {
					b = Obsidian{}
					return
				}
				b = Cobblestone{}
			}
		})
		if b != nil {
			ctx := event.C()
			w.Handler().HandleLiquidHarden(ctx, pos, l, water, b)
			ctx.Continue(func() {
				w.PlaySound(pos.Vec3Centre(), sound.Fizz{})
				w.PlaceBlock(pos, b)
			})
			return true
		}
		return false
	}
	water, ok = w.Block(*flownIntoBy).(Water)
	if !ok {
		return false
	}

	if l.Depth == 8 && !l.Falling {
		b = Obsidian{}
	} else {
		b = Cobblestone{}
	}
	ctx := event.C()
	w.Handler().HandleLiquidHarden(ctx, pos, l, water, b)
	ctx.Continue(func() {
		w.PlaceBlock(pos, b)
		w.PlaySound(pos.Vec3Centre(), sound.Fizz{})
	})
	return true
}

// EncodeBlock ...
func (l Lava) EncodeBlock() (name string, properties map[string]interface{}) {
	if l.Depth < 1 || l.Depth > 8 {
		panic("invalid lava depth, must be between 1 and 8")
	}
	v := 8 - l.Depth
	if l.Falling {
		v += 8
	}
	if l.Still {
		return "minecraft:lava", map[string]interface{}{"liquid_depth": int32(v)}
	}
	return "minecraft:flowing_lava", map[string]interface{}{"liquid_depth": int32(v)}
}

// Hash ...
func (l Lava) Hash() uint64 {
	return hashLava | (uint64(boolByte(l.Falling)) << 32) | (uint64(boolByte(l.Still)) << 33) | (uint64(l.Depth) << 34)
}
