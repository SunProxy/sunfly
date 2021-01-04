package entity_internal

import "github.com/sunproxy/sunfly/dragonfly/world"

// CanSolidify is a function used to check if a block affected by gravity can solidify.
var CanSolidify func(b world.Block, pos world.BlockPos, w *world.World) bool
