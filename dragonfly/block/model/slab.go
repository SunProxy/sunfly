package model

import (
	"github.com/go-gl/mathgl/mgl64"
	"github.com/sunproxy/sunfly/dragonfly/entity/physics"
	"github.com/sunproxy/sunfly/dragonfly/world"
)

// Slab is the model of a slab-like block, which is either a half block or a full block, depending on if the
// slab is double.
type Slab struct {
	Double, Top bool
}

// AABB ...
func (s Slab) AABB(world.BlockPos, *world.World) []physics.AABB {
	if s.Double {
		return []physics.AABB{physics.NewAABB(mgl64.Vec3{}, mgl64.Vec3{1, 1, 1})}
	}
	if s.Top {
		return []physics.AABB{physics.NewAABB(mgl64.Vec3{0, 0.5, 0}, mgl64.Vec3{1, 1, 1})}
	}
	return []physics.AABB{physics.NewAABB(mgl64.Vec3{}, mgl64.Vec3{1, 0.5, 1})}
}

// FaceSolid ...
func (s Slab) FaceSolid(_ world.BlockPos, face world.Face, _ *world.World) bool {
	if s.Double {
		return true
	} else if s.Top {
		return face == world.FaceUp
	}
	return face == world.FaceDown
}
