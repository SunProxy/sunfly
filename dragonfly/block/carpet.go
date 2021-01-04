package block

import (
	"github.com/go-gl/mathgl/mgl64"
	"github.com/sunproxy/sunfly/dragonfly/block/colour"
	"github.com/sunproxy/sunfly/dragonfly/item"
	"github.com/sunproxy/sunfly/dragonfly/world"
)

// Carpet is a colourful block that can be obtained by killing/shearing sheep, or crafted using four string.
type Carpet struct {
	noNBT
	carpet
	transparent

	// Colour is the colour of the carpet.
	Colour colour.Colour
}

// FlammabilityInfo ...
func (c Carpet) FlammabilityInfo() FlammabilityInfo {
	return FlammabilityInfo{
		Encouragement: 30,
		Flammability:  60,
		LavaFlammable: true,
	}
}

// CanDisplace ...
func (Carpet) CanDisplace(b world.Liquid) bool {
	_, water := b.(Water)
	return water
}

// SideClosed ...
func (Carpet) SideClosed(world.BlockPos, world.BlockPos, *world.World) bool {
	return false
}

// BreakInfo ...
func (c Carpet) BreakInfo() BreakInfo {
	return BreakInfo{
		Hardness:    0.1,
		Harvestable: alwaysHarvestable,
		Effective:   nothingEffective,
		Drops:       simpleDrops(item.NewStack(c, 1)),
	}
}

// EncodeItem ...
func (c Carpet) EncodeItem() (id int32, meta int16) {
	return 171, int16(c.Colour.Uint8())
}

// EncodeBlock ...
func (c Carpet) EncodeBlock() (name string, properties map[string]interface{}) {
	return "minecraft:carpet", map[string]interface{}{"color": c.Colour.String()}
}

// Hash ...
func (c Carpet) Hash() uint64 {
	return hashCarpet | (uint64(c.Colour.Uint8()) << 32)
}

// HasLiquidDrops ...
func (Carpet) HasLiquidDrops() bool {
	return true
}

// NeighbourUpdateTick ...
func (Carpet) NeighbourUpdateTick(pos, _ world.BlockPos, w *world.World) {
	if _, ok := w.Block(pos.Add(world.BlockPos{0, -1})).(Air); ok {
		w.BreakBlockWithoutParticles(pos)
	}
}

// UseOnBlock handles not placing carpets on top of air blocks.
func (c Carpet) UseOnBlock(pos world.BlockPos, face world.Face, _ mgl64.Vec3, w *world.World, user item.User, ctx *item.UseContext) (used bool) {
	pos, _, used = firstReplaceable(w, pos, face, c)
	if !used {
		return
	}

	if _, ok := w.Block((world.BlockPos{pos.X(), pos.Y() - 1, pos.Z()})).(Air); ok {
		return
	}

	place(w, pos, c, user, ctx)
	return placed(ctx)
}
