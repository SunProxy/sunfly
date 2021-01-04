package block

import (
	"github.com/sunproxy/sunfly/dragonfly/item"
)

// SoulSoil is a block naturally found only in the soul sand valley.
type SoulSoil struct {
	noNBT
	solid
}

// BreakInfo ...
func (s SoulSoil) BreakInfo() BreakInfo {
	return BreakInfo{
		Hardness:    0.5,
		Harvestable: alwaysHarvestable,
		Effective:   shovelEffective,
		Drops:       simpleDrops(item.NewStack(s, 1)),
	}
}

// EncodeItem ...
func (SoulSoil) EncodeItem() (id int32, meta int16) {
	return -236, 0
}
