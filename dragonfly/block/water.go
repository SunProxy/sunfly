package block

import (
	"github.com/sunproxy/sunfly/dragonfly/entity"
	"github.com/sunproxy/sunfly/dragonfly/event"
	"github.com/sunproxy/sunfly/dragonfly/world"
	"github.com/sunproxy/sunfly/dragonfly/world/sound"
	"time"
)

// Water is a natural fluid that generates abundantly in the world.
type Water struct {
	noNBT
	empty
	replaceable

	// Still makes the water appear as if it is not flowing.
	Still bool
	// Depth is the depth of the water. This is a number from 1-8, where 8 is a source block and 1 is the
	// smallest possible water block.
	Depth int
	// Falling specifies if the water is falling. Falling water will always appear as a source block, but its
	// behaviour differs when it starts spreading.
	Falling bool
}

// EntityCollide ...
func (w Water) EntityCollide(e world.Entity) {
	if flammable, ok := e.(entity.Flammable); ok {
		flammable.Extinguish()
	}
}

// LiquidDepth returns the depth of the water.
func (w Water) LiquidDepth() int {
	return w.Depth
}

// SpreadDecay returns 1 - The amount of levels decreased upon spreading.
func (Water) SpreadDecay() int {
	return 1
}

// WithDepth returns the water with the depth passed.
func (w Water) WithDepth(depth int, falling bool) world.Liquid {
	w.Depth = depth
	w.Falling = falling
	w.Still = false
	return w
}

// LiquidFalling returns Water.Falling.
func (w Water) LiquidFalling() bool {
	return w.Falling
}

// HasLiquidDrops ...
func (Water) HasLiquidDrops() bool {
	return false
}

// LightDiffusionLevel ...
func (Water) LightDiffusionLevel() uint8 {
	return 2
}

// ScheduledTick ...
func (w Water) ScheduledTick(pos world.BlockPos, wo *world.World) {
	if w.Depth == 7 {
		// Attempt to form new water source blocks.
		count := 0
		pos.Neighbours(func(neighbour world.BlockPos) {
			if neighbour[1] == pos[1] {
				if liquid, ok := wo.Liquid(neighbour); ok {
					if water, ok := liquid.(Water); ok && water.Depth == 8 && !water.Falling {
						count++
					}
				}
			}
		})
		if count >= 2 {
			func() {
				if canFlowInto(w, wo, pos.Side(world.FaceDown), true) {
					return
				}
				// Only form a new source block if there either is no water below this block, or if the water
				// below this is not falling (full source block).
				wo.SetLiquid(pos, Water{Depth: 8, Still: true})
			}()
		}
	}
	tickLiquid(w, pos, wo)
}

// NeighbourUpdateTick ...
func (Water) NeighbourUpdateTick(pos, _ world.BlockPos, wo *world.World) {
	wo.ScheduleBlockUpdate(pos, time.Second/4)
}

// LiquidType ...
func (Water) LiquidType() string {
	return "water"
}

// Harden hardens the water if lava flows into it.
func (w Water) Harden(pos world.BlockPos, wo *world.World, flownIntoBy *world.BlockPos) bool {
	if flownIntoBy == nil {
		return false
	}
	if lava, ok := wo.Block(pos.Side(world.FaceUp)).(Lava); ok {
		ctx := event.C()
		wo.Handler().HandleLiquidHarden(ctx, pos, w, lava, Stone{})
		ctx.Continue(func() {
			wo.PlaceBlock(pos, Stone{})
			wo.PlaySound(pos.Vec3Centre(), sound.Fizz{})
		})
		return true
	} else if lava, ok := wo.Block(*flownIntoBy).(Lava); ok {
		ctx := event.C()
		wo.Handler().HandleLiquidHarden(ctx, pos, w, lava, Cobblestone{})
		ctx.Continue(func() {
			wo.PlaceBlock(*flownIntoBy, Cobblestone{})
			wo.PlaySound(pos.Vec3Centre(), sound.Fizz{})
		})
		return true
	}
	return false
}

// EncodeBlock ...
func (w Water) EncodeBlock() (name string, properties map[string]interface{}) {
	if w.Depth < 1 || w.Depth > 8 {
		panic("invalid water depth, must be between 1 and 8")
	}
	v := 8 - w.Depth
	if w.Falling {
		v += 8
	}
	if w.Still {
		return "minecraft:water", map[string]interface{}{"liquid_depth": int32(v)}
	}
	return "minecraft:flowing_water", map[string]interface{}{"liquid_depth": int32(v)}
}

// Hash ...
func (w Water) Hash() uint64 {
	return hashWater | (uint64(boolByte(w.Falling)) << 32) | (uint64(boolByte(w.Still)) << 33) | (uint64(w.Depth) << 34)
}
