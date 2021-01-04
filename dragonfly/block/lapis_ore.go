package block

import (
	"github.com/sunproxy/sunfly/dragonfly/item"
	"github.com/sunproxy/sunfly/dragonfly/item/tool"
	"math/rand"
)

// LapisOre is an ore block from which lapis lazuli is obtained.
type LapisOre struct {
	noNBT
	solid
	bassDrum
}

// BreakInfo ...
func (l LapisOre) BreakInfo() BreakInfo {
	return BreakInfo{
		Hardness: 3,
		Harvestable: func(t tool.Tool) bool {
			return t.ToolType() == tool.TypePickaxe && t.HarvestLevel() >= tool.TierStone.HarvestLevel
		},
		Effective: pickaxeEffective,
		Drops:     simpleDrops(item.NewStack(item.LapisLazuli{}, rand.Intn(5)+4)), //TODO: Silk Touch
		XPDrops:   XPDropRange{2, 5},
	}
}

// EncodeItem ...
func (l LapisOre) EncodeItem() (id int32, meta int16) {
	return 21, 0
}
