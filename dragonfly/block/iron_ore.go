package block

import (
	"github.com/sunproxy/sunfly/dragonfly/item"
	"github.com/sunproxy/sunfly/dragonfly/item/tool"
)

// IronOre is a mineral block found underground.
type IronOre struct {
	noNBT
	solid
	bassDrum
}

// BreakInfo ...
func (i IronOre) BreakInfo() BreakInfo {
	return BreakInfo{
		Hardness: 3,
		Harvestable: func(t tool.Tool) bool {
			return t.ToolType() == tool.TypePickaxe && t.HarvestLevel() >= tool.TierStone.HarvestLevel
		},
		Effective: pickaxeEffective,
		Drops:     simpleDrops(item.NewStack(i, 1)),
	}
}

// EncodeItem ...
func (i IronOre) EncodeItem() (id int32, meta int16) {
	return 15, 0
}
