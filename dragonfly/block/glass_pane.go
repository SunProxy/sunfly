package block

import (
	"github.com/sunproxy/sunfly/dragonfly/item/tool"
	"github.com/sunproxy/sunfly/dragonfly/world"
)

// GlassPane is a transparent block that can be used as a more efficient alternative to glass blocks.
type GlassPane struct {
	noNBT
	transparent
	thin
	clicksAndSticks
}

// CanDisplace ...
func (p GlassPane) CanDisplace(b world.Liquid) bool {
	_, water := b.(Water)
	return water
}

// SideClosed ...
func (p GlassPane) SideClosed(world.BlockPos, world.BlockPos, *world.World) bool {
	return false
}

// BreakInfo ...
func (p GlassPane) BreakInfo() BreakInfo {
	return BreakInfo{
		Hardness: 0.3,
		Harvestable: func(t tool.Tool) bool {
			return true // TODO(lhochbaum): Glass panes can be silk touched, implement silk touch.
		},
		Effective: nothingEffective,
		Drops:     simpleDrops(),
	}
}

// EncodeItem ...
func (p GlassPane) EncodeItem() (id int32, meta int16) {
	return 102, meta
}
