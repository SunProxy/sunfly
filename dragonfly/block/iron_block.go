package block

import (
	"github.com/sunproxy/sunfly/dragonfly/block/instrument"
	"github.com/sunproxy/sunfly/dragonfly/item"
	"github.com/sunproxy/sunfly/dragonfly/item/tool"
)

// IronBlock is a precious metal block made from 9 iron ingots.
type IronBlock struct {
	noNBT
	solid
}

// Instrument ...
func (i IronBlock) Instrument() instrument.Instrument {
	return instrument.IronXylophone()
}

// BreakInfo ...
func (i IronBlock) BreakInfo() BreakInfo {
	return BreakInfo{
		Hardness: 5,
		Harvestable: func(t tool.Tool) bool {
			return t.ToolType() == tool.TypePickaxe && t.HarvestLevel() >= tool.TierStone.HarvestLevel
		},
		Effective: pickaxeEffective,
		Drops:     simpleDrops(item.NewStack(i, 1)),
	}
}

// PowersBeacon ...
func (IronBlock) PowersBeacon() bool {
	return true
}

// EncodeItem ...
func (IronBlock) EncodeItem() (id int32, meta int16) {
	return 42, 0
}
