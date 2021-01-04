package block

import (
	"github.com/sunproxy/sunfly/dragonfly/item"
	"github.com/sunproxy/sunfly/dragonfly/world"
)

// Sand is a block affected by gravity. It can come in a red variant.
type Sand struct {
	noNBT
	gravityAffected
	solid
	snare

	// Red toggles the red sand variant.
	Red bool
}

// NeighbourUpdateTick ...
func (s Sand) NeighbourUpdateTick(pos, _ world.BlockPos, w *world.World) {
	s.fall(s, pos, w)
}

// BreakInfo ...
func (s Sand) BreakInfo() BreakInfo {
	return BreakInfo{
		Hardness:    0.5,
		Harvestable: alwaysHarvestable,
		Effective:   shovelEffective,
		Drops:       simpleDrops(item.NewStack(s, 1)),
	}
}

// EncodeItem ...
func (s Sand) EncodeItem() (id int32, meta int16) {
	if s.Red {
		return 12, 1
	}
	return 12, 0
}
