package block

import (
	"github.com/sunproxy/sunfly/dragonfly/block/model"
	"github.com/sunproxy/sunfly/dragonfly/item"
	"github.com/sunproxy/sunfly/dragonfly/world"
)

// NetherBrickFence is the nether brick variant of the fence block.
type NetherBrickFence struct {
	noNBT
	transparent
}

// BreakInfo ...
func (n NetherBrickFence) BreakInfo() BreakInfo {
	return BreakInfo{
		Hardness:    2,
		Harvestable: pickaxeHarvestable,
		Effective:   pickaxeEffective,
		Drops:       simpleDrops(item.NewStack(n, 1)),
	}
}

// CanDisplace ...
func (NetherBrickFence) CanDisplace(b world.Liquid) bool {
	_, ok := b.(Water)
	return ok
}

// SideClosed ...
func (NetherBrickFence) SideClosed(world.BlockPos, world.BlockPos, *world.World) bool {
	return false
}

// Model ...
func (n NetherBrickFence) Model() world.BlockModel {
	return model.Fence{}
}

// EncodeItem ...
func (n NetherBrickFence) EncodeItem() (id int32, meta int16) {
	return 113, 0
}
