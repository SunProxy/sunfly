package block

import (
	"github.com/sunproxy/sunfly/dragonfly/block/instrument"
	"github.com/sunproxy/sunfly/dragonfly/item"
	"github.com/sunproxy/sunfly/dragonfly/item/tool"
)

// EmeraldBlock is a precious mineral block crafted using 9 emeralds.
type EmeraldBlock struct {
	noNBT
	solid
}

// Instrument ...
func (e EmeraldBlock) Instrument() instrument.Instrument {
	return instrument.Bit()
}

// BreakInfo ...
func (e EmeraldBlock) BreakInfo() BreakInfo {
	return BreakInfo{
		Hardness: 5,
		Harvestable: func(t tool.Tool) bool {
			return t.ToolType() == tool.TypePickaxe && t.HarvestLevel() >= tool.TierIron.HarvestLevel
		},
		Effective: pickaxeEffective,
		Drops:     simpleDrops(item.NewStack(e, 1)),
	}
}

// PowersBeacon ...
func (EmeraldBlock) PowersBeacon() bool {
	return true
}

// EncodeItem ...
func (EmeraldBlock) EncodeItem() (id int32, meta int16) {
	return 133, 0
}
