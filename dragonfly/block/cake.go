package block

import (
	"github.com/go-gl/mathgl/mgl64"
	"github.com/sunproxy/sunfly/dragonfly/block/model"
	"github.com/sunproxy/sunfly/dragonfly/item"
	"github.com/sunproxy/sunfly/dragonfly/world"
)

// Cake is an edible block.
type Cake struct {
	noNBT
	transparent

	// Bites is the amount of bites taken out of the cake.
	Bites int
}

// CanDisplace ...
func (c Cake) CanDisplace(b world.Liquid) bool {
	_, water := b.(Water)
	return water
}

// SideClosed ...
func (c Cake) SideClosed(world.BlockPos, world.BlockPos, *world.World) bool {
	return false
}

// UseOnBlock ...
func (c Cake) UseOnBlock(pos world.BlockPos, face world.Face, _ mgl64.Vec3, w *world.World, user item.User, ctx *item.UseContext) bool {
	pos, _, used := firstReplaceable(w, pos, face, c)
	if !used {
		return false
	}

	if _, air := w.Block(pos.Side(world.FaceDown)).(Air); air {
		return false
	}

	place(w, pos, c, user, ctx)
	return placed(ctx)
}

// NeighbourUpdateTick ...
func (c Cake) NeighbourUpdateTick(pos, _ world.BlockPos, w *world.World) {
	if _, air := w.Block(pos.Side(world.FaceDown)).(Air); air {
		w.BreakBlock(pos)
	}
}

// Activate ...
func (c Cake) Activate(pos world.BlockPos, _ world.Face, w *world.World, u item.User) {
	if i, ok := u.(interface {
		Saturate(food int, saturation float64)
	}); ok {
		i.Saturate(2, 0.4)
		c.Bites++
		if c.Bites > 6 {
			w.BreakBlockWithoutParticles(pos)
			return
		}
		w.PlaceBlock(pos, c)
	}
}

// BreakInfo ...
func (c Cake) BreakInfo() BreakInfo {
	return BreakInfo{
		Hardness:    0.5,
		Harvestable: nothingEffective,
		Effective:   nothingEffective,
		Drops:       simpleDrops(),
	}
}

// EncodeItem ...
func (c Cake) EncodeItem() (id int32, meta int16) {
	return 354, 0
}

// EncodeBlock ...
func (c Cake) EncodeBlock() (name string, properties map[string]interface{}) {
	return "minecraft:cake", map[string]interface{}{"bite_counter": int32(c.Bites)}
}

// Hash ...
func (c Cake) Hash() uint64 {
	return hashCake | (uint64(c.Bites) << 32)
}

// Model ...
func (c Cake) Model() world.BlockModel {
	return model.Cake{Bites: c.Bites}
}
