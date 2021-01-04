package entity

import (
	"fmt"
	"github.com/go-gl/mathgl/mgl64"
	"github.com/sunproxy/sunfly/dragonfly/entity/physics"
	"github.com/sunproxy/sunfly/dragonfly/entity/state"
	"github.com/sunproxy/sunfly/dragonfly/internal/entity_internal"
	"github.com/sunproxy/sunfly/dragonfly/internal/item_internal"
	"github.com/sunproxy/sunfly/dragonfly/item"
	"github.com/sunproxy/sunfly/dragonfly/world"
	"sync/atomic"
)

// FallingBlock is the entity form of a block that appears when a gravity-affected block loses its support.
type FallingBlock struct {
	block         world.Block
	velocity, pos atomic.Value

	*MovementComputer
}

// NewFallingBlock ...
func NewFallingBlock(block world.Block, pos mgl64.Vec3) *FallingBlock {
	f := &FallingBlock{block: block, MovementComputer: &MovementComputer{gravity: 0.04, dragBeforeGravity: true}}
	f.pos.Store(pos)
	f.velocity.Store(mgl64.Vec3{})
	return f
}

// Block ...
func (f *FallingBlock) Block() world.Block {
	return f.block
}

// Tick ...
func (f *FallingBlock) Tick(_ int64) {
	f.pos.Store(f.tickMovement(f))

	pos := world.BlockPosFromVec3(f.Position())
	if f.OnGround() || entity_internal.CanSolidify(f.block, pos, f.World()) {
		if item_internal.Replaceable(f.World(), pos, f.block) {
			f.World().PlaceBlock(pos, f.block)
		} else {
			if i, ok := f.block.(world.Item); ok {
				itemEntity := NewItem(item.NewStack(i, 1), f.Position())
				itemEntity.SetVelocity(mgl64.Vec3{})
				f.World().AddEntity(itemEntity)
			}
		}

		_ = f.Close()
	}
}

// Close ...
func (f *FallingBlock) Close() error {
	if f.World() != nil {
		f.World().RemoveEntity(f)
	}
	return nil
}

// Name ...
func (f *FallingBlock) Name() string {
	return fmt.Sprintf("%T", f.block)
}

// AABB ...
func (f *FallingBlock) AABB() physics.AABB {
	return physics.NewAABB(mgl64.Vec3{-0.49, 0, -0.49}, mgl64.Vec3{0.49, 0.98, 0.49})
}

// Position ...
func (f *FallingBlock) Position() mgl64.Vec3 {
	return f.pos.Load().(mgl64.Vec3)
}

// World ...
func (f *FallingBlock) World() *world.World {
	w, _ := world.OfEntity(f)
	return w
}

// Yaw ...
func (f *FallingBlock) Yaw() float64 {
	return 0
}

// Pitch ...
func (f *FallingBlock) Pitch() float64 {
	return 0
}

// State ...
func (f *FallingBlock) State() []state.State {
	return nil
}

// Velocity ...
func (f *FallingBlock) Velocity() mgl64.Vec3 {
	return f.velocity.Load().(mgl64.Vec3)
}

// SetVelocity ...
func (f *FallingBlock) SetVelocity(v mgl64.Vec3) {
	f.velocity.Store(v)
}

// EncodeEntity ...
func (f *FallingBlock) EncodeEntity() string {
	return "minecraft:falling_block"
}
