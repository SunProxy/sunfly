package block

import (
	"github.com/sunproxy/sunfly/dragonfly/item"
)

// Terracotta is a block formed from clay, with a hardness and blast resistance comparable to stone. For colouring it,
// take a look at StainedTerracotta.
type Terracotta struct {
	noNBT
	solid
	bassDrum
}

// BreakInfo ...
func (t Terracotta) BreakInfo() BreakInfo {
	return BreakInfo{
		Hardness:    1.25,
		Harvestable: pickaxeHarvestable,
		Effective:   pickaxeEffective,
		Drops:       simpleDrops(item.NewStack(t, 1)),
	}
}

// EncodeItem ...
func (t Terracotta) EncodeItem() (id int32, meta int16) {
	return 172, meta
}
