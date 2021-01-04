package model

import (
	"github.com/sunproxy/sunfly/dragonfly/entity/physics"
	"github.com/sunproxy/sunfly/dragonfly/world"
)

// Empty is a model that is completely empty. It has no collision boxes or solid faces.
type Empty struct{}

// AABB ...
func (Empty) AABB(world.BlockPos, *world.World) []physics.AABB {
	return nil
}

// FaceSolid ...
func (Empty) FaceSolid(world.BlockPos, world.Face, *world.World) bool {
	return false
}
