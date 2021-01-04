package block

import (
	"github.com/go-gl/mathgl/mgl64"
	"github.com/sunproxy/sunfly/dragonfly/block/fire"
	"github.com/sunproxy/sunfly/dragonfly/block/model"
	"github.com/sunproxy/sunfly/dragonfly/item"
	"github.com/sunproxy/sunfly/dragonfly/world"
)

// Lantern is a light emitting block.
type Lantern struct {
	noNBT
	transparent

	// Hanging determines if a lantern is hanging off a block.
	Hanging bool
	// Type of fire lighting the lantern.
	Type fire.Fire
}

// Model ...
func (l Lantern) Model() world.BlockModel {
	return model.Lantern{Hanging: l.Hanging}
}

// NeighbourUpdateTick ...
func (l Lantern) NeighbourUpdateTick(pos, _ world.BlockPos, w *world.World) {
	if l.Hanging {
		up := pos.Side(world.FaceUp)
		if !w.Block(up).Model().FaceSolid(up, world.FaceDown, w) {
			w.BreakBlockWithoutParticles(pos)
		}
	} else {
		down := pos.Side(world.FaceDown)
		if !w.Block(down).Model().FaceSolid(down, world.FaceUp, w) {
			w.BreakBlockWithoutParticles(pos)
		}
	}
}

// LightEmissionLevel ...
func (l Lantern) LightEmissionLevel() uint8 {
	return l.Type.LightLevel
}

// UseOnBlock ...
func (l Lantern) UseOnBlock(pos world.BlockPos, face world.Face, _ mgl64.Vec3, w *world.World, user item.User, ctx *item.UseContext) bool {
	pos, face, used := firstReplaceable(w, pos, face, l)
	if !used {
		return false
	}
	if face == world.FaceDown {
		upPos := pos.Side(world.FaceUp)
		if !w.Block(upPos).Model().FaceSolid(upPos, world.FaceDown, w) {
			face = world.FaceUp
		}
	}
	if face != world.FaceDown {
		downPos := pos.Side(world.FaceDown)
		if !w.Block(downPos).Model().FaceSolid(downPos, world.FaceUp, w) {
			return false
		}
	}
	l.Hanging = face == world.FaceDown

	place(w, pos, l, user, ctx)
	return placed(ctx)
}

// HasLiquidDrops ...
func (l Lantern) HasLiquidDrops() bool {
	return true
}

// BreakInfo ...
func (l Lantern) BreakInfo() BreakInfo {
	return BreakInfo{
		Hardness:    3.5,
		Harvestable: pickaxeHarvestable,
		Effective:   pickaxeEffective,
		Drops:       simpleDrops(item.NewStack(l, 1)),
	}
}

// EncodeItem ...
func (l Lantern) EncodeItem() (id int32, meta int16) {
	switch l.Type {
	case fire.Normal():
		return -208, 0
	case fire.Soul():
		return -269, 0
	}
	panic("invalid fire type")
}

// EncodeBlock ...
func (l Lantern) EncodeBlock() (name string, properties map[string]interface{}) {
	switch l.Type {
	case fire.Normal():
		return "minecraft:lantern", map[string]interface{}{"hanging": l.Hanging}
	case fire.Soul():
		return "minecraft:soul_lantern", map[string]interface{}{"hanging": l.Hanging}
	}
	panic("invalid fire type")
}

// Hash ...
func (l Lantern) Hash() uint64 {
	return hashLantern | (uint64(boolByte(l.Hanging)) << 32) | (uint64(l.Type.Uint8()) << 33)
}
