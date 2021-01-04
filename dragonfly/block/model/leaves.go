package model

import (
	"github.com/go-gl/mathgl/mgl64"
	"github.com/sunproxy/sunfly/dragonfly/entity/physics"
	"github.com/sunproxy/sunfly/dragonfly/world"
)

// Leaves is a model for leaves-like blocks. These blocks have a full collision box, but none of their faces
// are solid.
type Leaves struct{}

// AABB ...
func (Leaves) AABB(world.BlockPos, *world.World) []physics.AABB {
	return []physics.AABB{physics.NewAABB(mgl64.Vec3{}, mgl64.Vec3{1, 1, 1})}
}

// FaceSolid ...
func (Leaves) FaceSolid(world.BlockPos, world.Face, *world.World) bool {
	return false
}
