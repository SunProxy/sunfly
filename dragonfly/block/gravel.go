package block

import (
	"github.com/sunproxy/sunfly/dragonfly/item"
	"github.com/sunproxy/sunfly/dragonfly/item/tool"
	"github.com/sunproxy/sunfly/dragonfly/world"
	"math/rand"
)

// Gravel is a block affected by gravity. It has a 10% chance of dropping flint instead of itself on break.
type Gravel struct {
	noNBT
	gravityAffected
	solid
	snare
}

// NeighbourUpdateTick ...
func (g Gravel) NeighbourUpdateTick(pos, _ world.BlockPos, w *world.World) {
	g.fall(g, pos, w)
}

// BreakInfo ...
func (g Gravel) BreakInfo() BreakInfo {
	return BreakInfo{
		Hardness:    0.6,
		Harvestable: alwaysHarvestable,
		Effective:   shovelEffective,
		Drops: func(t tool.Tool) []item.Stack {
			if rand.Float64() < 0.1 {
				return []item.Stack{item.NewStack(item.Flint{}, 1)}
			}
			return []item.Stack{item.NewStack(g, 1)}
		},
	}
}

// EncodeItem ...
func (g Gravel) EncodeItem() (id int32, meta int16) {
	return 13, 0
}
