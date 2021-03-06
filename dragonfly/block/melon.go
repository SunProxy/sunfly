package block

import (
	"github.com/sunproxy/sunfly/dragonfly/item"
	"math/rand"
)

// Melon is a fruit block that grows from melon stems.
type Melon struct {
	noNBT
	solid
}

// BreakInfo ...
func (m Melon) BreakInfo() BreakInfo {
	return BreakInfo{
		Hardness:    1,
		Harvestable: alwaysHarvestable,
		Effective:   axeEffective,
		Drops:       simpleDrops(item.NewStack(item.MelonSlice{}, rand.Intn(5)+3)), //TODO: Silk Touch
	}
}

// EncodeItem ...
func (m Melon) EncodeItem() (id int32, meta int16) {
	return 103, 0
}
