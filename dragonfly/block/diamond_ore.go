package block

import (
	"github.com/sunproxy/sunfly/dragonfly/item"
	"github.com/sunproxy/sunfly/dragonfly/item/tool"
)

// DiamondOre is a rare ore that generates underground.
type DiamondOre struct {
	noNBT
	solid
	bassDrum
}

// BreakInfo ...
func (d DiamondOre) BreakInfo() BreakInfo {
	return BreakInfo{
		Hardness: 3,
		Harvestable: func(t tool.Tool) bool {
			return t.ToolType() == tool.TypePickaxe && t.HarvestLevel() >= tool.TierIron.HarvestLevel
		},
		Effective: pickaxeEffective,
		Drops:     simpleDrops(item.NewStack(item.Diamond{}, 1)), //TODO: Silk Touch
		XPDrops:   XPDropRange{3, 7},
	}
}

// EncodeItem ...
func (d DiamondOre) EncodeItem() (id int32, meta int16) {
	return 56, 0
}
