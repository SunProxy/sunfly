package block

import (
	"github.com/sunproxy/sunfly/dragonfly/world"
)

// Air is the block present in otherwise empty space.
type Air struct {
	noNBT
	empty
	replaceable
	transparent
}

// CanDisplace ...
func (Air) CanDisplace(world.Liquid) bool {
	return true
}

// HasLiquidDrops ...
func (Air) HasLiquidDrops() bool {
	return false
}

// EncodeItem ...
func (Air) EncodeItem() (id int32, meta int16) {
	return 0, 0
}
