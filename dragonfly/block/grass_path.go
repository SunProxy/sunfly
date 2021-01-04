package block

import (
	"github.com/sunproxy/sunfly/dragonfly/item"
	"github.com/sunproxy/sunfly/dragonfly/world"
)

// GrassPath is a decorative block that can be created by using a shovel on a grass block.
type GrassPath struct {
	noNBT
	tilledGrass
	transparent
}

// NeighbourUpdateTick handles the turning from grass path into dirt if a block is placed on top of it.
func (p GrassPath) NeighbourUpdateTick(pos, _ world.BlockPos, w *world.World) {
	up := pos.Add(world.BlockPos{0, 1})
	if w.Block(up).Model().FaceSolid(up, world.FaceDown, w) {
		// A block with a solid side at the bottom was placed onto this one.
		w.SetBlock(pos, Dirt{})
	}
}

// BreakInfo ...
func (p GrassPath) BreakInfo() BreakInfo {
	return BreakInfo{
		Hardness:    0.6,
		Harvestable: alwaysHarvestable,
		Effective:   shovelEffective,
		Drops:       simpleDrops(item.NewStack(Dirt{}, 1)),
	}
}

// EncodeItem ...
func (p GrassPath) EncodeItem() (id int32, meta int16) {
	return 198, 0
}
