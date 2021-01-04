package block

import (
	"github.com/sunproxy/sunfly/dragonfly/item"
	"github.com/sunproxy/sunfly/dragonfly/item/tool"
	"math/rand"
)

// GildedBlackstone is a variant of blackstone that can drop itself or gold nuggets when mined.
type GildedBlackstone struct {
	noNBT
	solid
}

// BreakInfo ...
func (b GildedBlackstone) BreakInfo() BreakInfo {
	return BreakInfo{
		Hardness:    1.5,
		Harvestable: pickaxeHarvestable,
		Effective:   pickaxeEffective,
		Drops: func(t tool.Tool) []item.Stack {
			if rand.Float64() < 0.1 {
				return []item.Stack{item.NewStack(item.GoldNugget{}, rand.Intn(4)+2)}
			}
			return []item.Stack{item.NewStack(b, 1)}
		},
	}
}

// EncodeItem ...
func (GildedBlackstone) EncodeItem() (id int32, meta int16) {
	return -281, 0
}
