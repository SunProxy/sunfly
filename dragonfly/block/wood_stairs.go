package block

import (
	"github.com/sunproxy/sunfly/dragonfly/block/model"
	"github.com/sunproxy/sunfly/dragonfly/block/wood"
	"github.com/sunproxy/sunfly/dragonfly/item"
	"github.com/sunproxy/sunfly/dragonfly/world"
	"github.com/go-gl/mathgl/mgl64"
)

// WoodStairs are blocks that allow entities to walk up blocks without jumping. They are crafted using planks.
type WoodStairs struct {
	noNBT
	transparent
	bass

	// Wood is the type of wood of the stairs. This field must have one of the values found in the material
	// package.
	Wood wood.Wood
	// UpsideDown specifies if the stairs are upside down. If set to true, the full side is at the top part
	// of the block.
	UpsideDown bool
	// Facing is the direction that the full side of the stairs is facing.
	Facing world.Direction
}

// FlammabilityInfo ...
func (f WoodStairs) FlammabilityInfo() FlammabilityInfo {
	if !f.Wood.Flammable() {
		return FlammabilityInfo{}
	}
	return FlammabilityInfo{
		Encouragement: 5,
		Flammability:  20,
		LavaFlammable: true,
	}
}

// UseOnBlock handles the directional placing of stairs and makes sure they are properly placed upside down
// when needed.
func (s WoodStairs) UseOnBlock(pos world.BlockPos, face world.Face, clickPos mgl64.Vec3, w *world.World, user item.User, ctx *item.UseContext) (used bool) {
	pos, face, used = firstReplaceable(w, pos, face, s)
	if !used {
		return
	}
	s.Facing = user.Facing()
	if face == world.FaceDown || (clickPos[1] > 0.5 && face != world.FaceUp) {
		s.UpsideDown = true
	}

	place(w, pos, s, user, ctx)
	return placed(ctx)
}

// Model ...
func (s WoodStairs) Model() world.BlockModel {
	return model.Stair{Facing: s.Facing, UpsideDown: s.UpsideDown}
}

// BreakInfo ...
func (s WoodStairs) BreakInfo() BreakInfo {
	return BreakInfo{
		Hardness:    2,
		Harvestable: alwaysHarvestable,
		Effective:   axeEffective,
		Drops:       simpleDrops(item.NewStack(s, 1)),
	}
}

// EncodeItem ...
func (s WoodStairs) EncodeItem() (id int32, meta int16) {
	switch s.Wood {
	case wood.Oak():
		return 53, 0
	case wood.Spruce():
		return 134, 0
	case wood.Birch():
		return 135, 0
	case wood.Jungle():
		return 136, 0
	case wood.Acacia():
		return 163, 0
	case wood.DarkOak():
		return 164, 0
	case wood.Crimson():
		return -254, 0
	case wood.Warped():
		return -255, 0
	}
	panic("invalid wood type")
}

// EncodeBlock ...
func (s WoodStairs) EncodeBlock() (name string, properties map[string]interface{}) {
	return "minecraft:" + s.Wood.String() + "_stairs", map[string]interface{}{"upside_down_bit": s.UpsideDown, "weirdo_direction": toStairsDirection(s.Facing)}
}

// Hash ...
func (s WoodStairs) Hash() uint64 {
	return hashWoodStairs | (uint64(boolByte(s.UpsideDown)) << 32) | (uint64(s.Facing) << 33) | (uint64(s.Wood.Uint8()) << 35)
}

// toStairDirection converts a facing to a stairs direction for Minecraft.
func toStairsDirection(v world.Direction) int32 {
	return int32(3 - v)
}

// CanDisplace ...
func (WoodStairs) CanDisplace(b world.Liquid) bool {
	_, ok := b.(Water)
	return ok
}

// SideClosed ...
func (s WoodStairs) SideClosed(pos, side world.BlockPos, w *world.World) bool {
	return s.Model().FaceSolid(pos, pos.Face(side), w)
}
