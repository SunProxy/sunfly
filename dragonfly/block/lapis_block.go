package block

import (
	"github.com/sunproxy/sunfly/dragonfly/item"
	"github.com/sunproxy/sunfly/dragonfly/item/tool"
)

// LapisBlock is a decorative mineral block that is crafted from lapis lazuli.
type LapisBlock struct {
	noNBT
	solid
}

// BreakInfo ...
func (l LapisBlock) BreakInfo() BreakInfo {
	return BreakInfo{
		Hardness: 3,
		Harvestable: func(t tool.Tool) bool {
			return t.ToolType() == tool.TypePickaxe && t.HarvestLevel() >= tool.TierStone.HarvestLevel
		},
		Effective: pickaxeEffective,
		Drops:     simpleDrops(item.NewStack(l, 1)),
	}
}

// EncodeItem ...
func (LapisBlock) EncodeItem() (id int32, meta int16) {
	return 22, 0
}
