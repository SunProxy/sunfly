package model

import (
	"github.com/go-gl/mathgl/mgl64"
	"github.com/sunproxy/sunfly/dragonfly/entity/physics"
	"github.com/sunproxy/sunfly/dragonfly/world"
)

// FenceGate is a model used by fence gates.
type FenceGate struct {
	Facing world.Direction
	Open   bool
}

// AABB ...
func (f FenceGate) AABB(pos world.BlockPos, w *world.World) []physics.AABB {
	if f.Open {
		return nil
	}
	return []physics.AABB{physics.NewAABB(mgl64.Vec3{}, mgl64.Vec3{1, 1.5, 1}).Stretch(int(f.Facing.Face().Axis()), -0.375)}
}

// FaceSolid ...
func (f FenceGate) FaceSolid(pos world.BlockPos, face world.Face, w *world.World) bool {
	return false
}
