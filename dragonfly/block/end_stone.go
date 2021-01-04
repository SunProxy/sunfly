package block

import "github.com/sunproxy/sunfly/dragonfly/item"

// EndStone is a block found in The End.
type EndStone struct {
	noNBT
	solid
	bassDrum
}

// BreakInfo ...
func (e EndStone) BreakInfo() BreakInfo {
	return BreakInfo{
		Hardness:    3,
		Harvestable: pickaxeHarvestable,
		Effective:   pickaxeEffective,
		Drops:       simpleDrops(item.NewStack(e, 1)),
	}
}

// EncodeItem ...
func (e EndStone) EncodeItem() (id int32, meta int16) {
	return 121, 0
}
