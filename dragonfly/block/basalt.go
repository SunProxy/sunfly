package block

import (
	"github.com/go-gl/mathgl/mgl64"
	"github.com/sunproxy/sunfly/dragonfly/item"
	"github.com/sunproxy/sunfly/dragonfly/world"
)

// Basalt is a type of igneous rock found in the Nether.
type Basalt struct {
	noNBT
	solid
	bassDrum

	// Polished specifies if the basalt is its polished variant.
	Polished bool
	// Axis is the axis which the basalt faces.
	Axis world.Axis
}

// UseOnBlock ...
func (b Basalt) UseOnBlock(pos world.BlockPos, face world.Face, _ mgl64.Vec3, w *world.World, user item.User, ctx *item.UseContext) (used bool) {
	pos, face, used = firstReplaceable(w, pos, face, b)
	if !used {
		return
	}
	b.Axis = face.Axis()

	place(w, pos, b, user, ctx)
	return placed(ctx)
}

// BreakInfo ...
func (b Basalt) BreakInfo() BreakInfo {
	return BreakInfo{
		Hardness:    1.25,
		Harvestable: pickaxeHarvestable,
		Effective:   pickaxeEffective,
		Drops:       simpleDrops(item.NewStack(b, 1)),
	}
}

// EncodeItem ...
func (b Basalt) EncodeItem() (id int32, meta int16) {
	if b.Polished {
		return -235, 0
	}
	return -234, 0
}

// EncodeBlock ...
func (b Basalt) EncodeBlock() (name string, properties map[string]interface{}) {
	if b.Polished {
		return "minecraft:polished_basalt", map[string]interface{}{"pillar_axis": b.Axis.String()}
	}
	return "minecraft:basalt", map[string]interface{}{"pillar_axis": b.Axis.String()}
}

// Hash ...
func (b Basalt) Hash() uint64 {
	return hashBasalt | (uint64(boolByte(b.Polished)) << 32) | (uint64(b.Axis) << 33)
}
