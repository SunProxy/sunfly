package block

import (
	"github.com/sunproxy/sunfly/dragonfly/item"
	"math/rand"
)

// NetherGoldOre is a variant of gold ore found exclusively in The Nether.
type NetherGoldOre struct {
	noNBT
	solid
}

// BreakInfo ...
func (n NetherGoldOre) BreakInfo() BreakInfo {
	return BreakInfo{
		Hardness:    3,
		Harvestable: pickaxeHarvestable,
		Effective:   pickaxeEffective,
		Drops:       simpleDrops(item.NewStack(item.GoldNugget{}, rand.Intn(4)+2)), //TODO: Silk Touch
		XPDrops:     XPDropRange{0, 1},
	}
}

// EncodeItem ...
func (n NetherGoldOre) EncodeItem() (id int32, meta int16) {
	return -288, 0
}
