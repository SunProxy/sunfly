package block

import (
	"github.com/sunproxy/sunfly/dragonfly/item"
	"math/rand"
)

// SeaLantern is an underwater light sources that appear in ocean monuments and underwater ruins.
type SeaLantern struct {
	noNBT
	transparent
	solid
	clicksAndSticks
}

// LightEmissionLevel ...
func (SeaLantern) LightEmissionLevel() uint8 {
	return 15
}

// BreakInfo ...
func (s SeaLantern) BreakInfo() BreakInfo {
	return BreakInfo{
		Hardness:    0.3,
		Harvestable: alwaysHarvestable,
		Effective:   nothingEffective,
		Drops:       simpleDrops(item.NewStack(item.PrismarineCrystals{}, rand.Intn(2)+2)), //TODO: Silk Touch
	}
}

// EncodeItem ...
func (SeaLantern) EncodeItem() (id int32, meta int16) {
	return 169, 0
}
