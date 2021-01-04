package model

import (
	"github.com/go-gl/mathgl/mgl64"
	"github.com/sunproxy/sunfly/dragonfly/entity/physics"
	"github.com/sunproxy/sunfly/dragonfly/world"
)

// CocoaBean is a model used by cocoa bean blocks.
type CocoaBean struct {
	Facing world.Direction
	Age    int
}

// AABB ...
func (c CocoaBean) AABB(pos world.BlockPos, w *world.World) []physics.AABB {
	return []physics.AABB{physics.NewAABB(mgl64.Vec3{}, mgl64.Vec3{1, 1, 1}).
		Stretch(int(c.Facing.Rotate90().Face().Axis()), -((6-float64(c.Age))/16)).
		ExtendTowards(int(world.FaceDown), -0.25).
		ExtendTowards(int(world.FaceUp), -((7-float64(c.Age)*2)/16)).
		ExtendTowards(int(c.Facing.Face()), -0.0625).
		ExtendTowards(int(c.Facing.Opposite().Face()), -((11 - float64(c.Age)*2) / 16))}
}

// FaceSolid ...
func (c CocoaBean) FaceSolid(pos world.BlockPos, face world.Face, w *world.World) bool {
	return false
}
