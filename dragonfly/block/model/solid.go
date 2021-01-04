package model

import (
	"github.com/go-gl/mathgl/mgl64"
	"github.com/sunproxy/sunfly/dragonfly/entity/physics"
	"github.com/sunproxy/sunfly/dragonfly/world"
)

// Solid is the model of a fully solid block. Blocks with this model, such as stone or wooden planks, have a
// 1x1x1 collision box.
type Solid struct{}

// AABB ...
func (Solid) AABB(world.BlockPos, *world.World) []physics.AABB {
	return []physics.AABB{physics.NewAABB(mgl64.Vec3{}, mgl64.Vec3{1, 1, 1})}
}

// FaceSolid ...
func (Solid) FaceSolid(world.BlockPos, world.Face, *world.World) bool {
	return true
}
