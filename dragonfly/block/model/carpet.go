package model

import (
	"github.com/go-gl/mathgl/mgl64"
	"github.com/sunproxy/sunfly/dragonfly/entity/physics"
	"github.com/sunproxy/sunfly/dragonfly/world"
)

// Carpet is a model for carpet-like extremely thin blocks.
type Carpet struct{}

// AABB ...
func (Carpet) AABB(world.BlockPos, *world.World) []physics.AABB {
	return []physics.AABB{physics.NewAABB(mgl64.Vec3{}, mgl64.Vec3{1, 0.0625, 1})}
}

// FaceSolid ...
func (Carpet) FaceSolid(world.BlockPos, world.Face, *world.World) bool {
	return false
}
