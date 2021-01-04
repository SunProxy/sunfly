package block

import (
	"github.com/sunproxy/sunfly/dragonfly/item"
)

// Bricks are decorative building blocks.
type Bricks struct {
	noNBT
	solid
	bassDrum
}

// BreakInfo ...
func (b Bricks) BreakInfo() BreakInfo {
	return BreakInfo{
		Hardness:    2,
		Harvestable: pickaxeHarvestable,
		Effective:   pickaxeEffective,
		Drops:       simpleDrops(item.NewStack(b, 1)),
	}
}

// EncodeItem ...
func (Bricks) EncodeItem() (id int32, meta int16) {
	return 45, 0
}
