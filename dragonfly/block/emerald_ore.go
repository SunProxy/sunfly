package block

import (
	"github.com/sunproxy/sunfly/dragonfly/item"
	"github.com/sunproxy/sunfly/dragonfly/item/tool"
)

// EmeraldOre is an ore generating exclusively under mountain biomes.
type EmeraldOre struct {
	noNBT
	solid
	bassDrum
}

// BreakInfo ...
func (e EmeraldOre) BreakInfo() BreakInfo {
	return BreakInfo{
		Hardness: 3,
		Harvestable: func(t tool.Tool) bool {
			return t.ToolType() == tool.TypePickaxe && t.HarvestLevel() >= tool.TierIron.HarvestLevel
		},
		Effective: pickaxeEffective,
		Drops:     simpleDrops(item.NewStack(item.Emerald{}, 1)), //TODO: Silk Touch
		XPDrops:   XPDropRange{3, 7},
	}
}

// EncodeItem ...
func (e EmeraldOre) EncodeItem() (id int32, meta int16) {
	return 129, 0
}
