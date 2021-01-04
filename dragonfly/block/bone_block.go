package block

import (
	"github.com/go-gl/mathgl/mgl64"
	"github.com/sunproxy/sunfly/dragonfly/block/instrument"
	"github.com/sunproxy/sunfly/dragonfly/item"
	"github.com/sunproxy/sunfly/dragonfly/world"
)

// BoneBlock is a decorative block that can face different directions.
type BoneBlock struct {
	noNBT
	solid

	// Axis is the axis which the bone block faces.
	Axis world.Axis
}

// Instrument ...
func (b BoneBlock) Instrument() instrument.Instrument {
	return instrument.Xylophone()
}

// UseOnBlock handles the rotational placing of bone blocks.
func (b BoneBlock) UseOnBlock(pos world.BlockPos, face world.Face, _ mgl64.Vec3, w *world.World, user item.User, ctx *item.UseContext) (used bool) {
	pos, face, used = firstReplaceable(w, pos, face, b)
	if !used {
		return
	}
	b.Axis = face.Axis()

	place(w, pos, b, user, ctx)
	return placed(ctx)
}

// BreakInfo ...
func (b BoneBlock) BreakInfo() BreakInfo {
	return BreakInfo{
		Hardness:    2,
		Harvestable: pickaxeHarvestable,
		Effective:   pickaxeEffective,
		Drops:       simpleDrops(item.NewStack(b, 1)),
	}
}

// EncodeItem ...
func (b BoneBlock) EncodeItem() (id int32, meta int16) {
	return 216, 0
}

// EncodeBlock ...
func (b BoneBlock) EncodeBlock() (name string, properties map[string]interface{}) {
	return "minecraft:bone_block", map[string]interface{}{"pillar_axis": b.Axis.String(), "deprecated": int32(0)}
}

// Hash ...
func (b BoneBlock) Hash() uint64 {
	return hashBoneBlock | (uint64(b.Axis << 32))
}
