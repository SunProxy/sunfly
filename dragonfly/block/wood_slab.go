package block

import (
	"github.com/go-gl/mathgl/mgl64"
	"github.com/sunproxy/sunfly/dragonfly/block/model"
	"github.com/sunproxy/sunfly/dragonfly/block/wood"
	"github.com/sunproxy/sunfly/dragonfly/entity/physics"
	"github.com/sunproxy/sunfly/dragonfly/item"
	"github.com/sunproxy/sunfly/dragonfly/item/tool"
	"github.com/sunproxy/sunfly/dragonfly/world"
)

// WoodSlab is a half block that allows entities to walk up blocks without jumping.
type WoodSlab struct {
	noNBT
	bass

	// Wood is the type of wood of the slabs. This field must have one of the values found in the material
	// package.
	Wood wood.Wood
	// Top specifies if the slab is in the top part of the block.
	Top bool
	// Double specifies if the slab is a double slab. These double slabs can be made by placing another slab
	// on an existing slab.
	Double bool
}

// FlammabilityInfo ...
func (s WoodSlab) FlammabilityInfo() FlammabilityInfo {
	if !s.Wood.Flammable() {
		return FlammabilityInfo{}
	}
	return FlammabilityInfo{
		Encouragement: 5,
		Flammability:  20,
		LavaFlammable: true,
	}
}

// Model ...
func (s WoodSlab) Model() world.BlockModel {
	return model.Slab{Double: s.Double, Top: s.Top}
}

// UseOnBlock handles the placement of slabs with relation to them being upside down or not and handles slabs
// being turned into double slabs.
func (s WoodSlab) UseOnBlock(pos world.BlockPos, face world.Face, clickPos mgl64.Vec3, w *world.World, user item.User, ctx *item.UseContext) (used bool) {
	clickedBlock := w.Block(pos)
	if clickedSlab, ok := clickedBlock.(WoodSlab); ok && !s.Double {
		if (face == world.FaceUp && !clickedSlab.Double && clickedSlab.Wood == s.Wood && !clickedSlab.Top) ||
			(face == world.FaceDown && !clickedSlab.Double && clickedSlab.Wood == s.Wood && clickedSlab.Top) {
			// A half slab of the same type was clicked at the top, so we can make it full.
			clickedSlab.Double = true

			place(w, pos, clickedSlab, user, ctx)
			return placed(ctx)
		}
	}
	if sideSlab, ok := w.Block(pos.Side(face)).(WoodSlab); ok && !replaceableWith(w, pos, s) && !s.Double {
		// The block on the side of the one clicked was a slab and the block clicked was not replaceableWith, so
		// the slab on the side must've been half and may now be filled if the wood types are the same.
		if !sideSlab.Double && sideSlab.Wood == s.Wood {
			sideSlab.Double = true

			place(w, pos.Side(face), sideSlab, user, ctx)
			return placed(ctx)
		}
	}
	pos, face, used = firstReplaceable(w, pos, face, s)
	if !used {
		return
	}
	if face == world.FaceDown || (clickPos[1] > 0.5 && face != world.FaceUp) {
		s.Top = true
	}

	place(w, pos, s, user, ctx)
	return placed(ctx)
}

// BreakInfo ...
func (s WoodSlab) BreakInfo() BreakInfo {
	return BreakInfo{
		Hardness:    2,
		Harvestable: alwaysHarvestable,
		Effective:   axeEffective,
		Drops: func(t tool.Tool) []item.Stack {
			if s.Double {
				s.Double = false
				// If the slab is double, it should drop two single slabs.
				return []item.Stack{item.NewStack(s, 2)}
			}
			return []item.Stack{item.NewStack(s, 1)}
		},
	}
}

// LightDiffusionLevel returns 0 if the slab is a half slab, or 15 if it is double.
func (s WoodSlab) LightDiffusionLevel() uint8 {
	if s.Double {
		return 15
	}
	return 0
}

// AABB ...
func (s WoodSlab) AABB(world.BlockPos, *world.World) []physics.AABB {
	if s.Double {
		return []physics.AABB{physics.NewAABB(mgl64.Vec3{}, mgl64.Vec3{1, 1, 1})}
	}
	if s.Top {
		return []physics.AABB{physics.NewAABB(mgl64.Vec3{0, 0.5, 0}, mgl64.Vec3{1, 1, 1})}
	}
	return []physics.AABB{physics.NewAABB(mgl64.Vec3{}, mgl64.Vec3{1, 0.5, 1})}
}

// EncodeItem ...
func (s WoodSlab) EncodeItem() (id int32, meta int16) {
	switch s.Wood {
	case wood.Oak():
		if s.Double {
			return 157, 0
		}
		return 158, 0
	case wood.Spruce():
		if s.Double {
			return 157, 1
		}
		return 158, 1
	case wood.Birch():
		if s.Double {
			return 157, 2
		}
		return 158, 2
	case wood.Jungle():
		if s.Double {
			return 157, 3
		}
		return 158, 3
	case wood.Acacia():
		if s.Double {
			return 157, 4
		}
		return 158, 4
	case wood.DarkOak():
		if s.Double {
			return 157, 5
		}
		return 158, 5
	case wood.Crimson():
		if s.Double {
			return -266, 0
		}
		return -264, 0
	case wood.Warped():
		if s.Double {
			return -267, 0
		}
		return -265, 0
	}
	panic("invalid wood type")
}

// EncodeBlock ...
func (s WoodSlab) EncodeBlock() (name string, properties map[string]interface{}) {
	if s.Double {
		if s.Wood == wood.Crimson() || s.Wood == wood.Warped() {
			return "minecraft:" + s.Wood.String() + "_double_slab", map[string]interface{}{"top_slot_bit": s.Top}
		}
		return "minecraft:double_wooden_slab", map[string]interface{}{"top_slot_bit": s.Top, "wood_type": s.Wood.String()}
	}
	if s.Wood == wood.Crimson() || s.Wood == wood.Warped() {
		return "minecraft:" + s.Wood.String() + "_slab", map[string]interface{}{"top_slot_bit": s.Top}
	}
	return "minecraft:wooden_slab", map[string]interface{}{"top_slot_bit": s.Top, "wood_type": s.Wood.String()}
}

// Hash ...
func (s WoodSlab) Hash() uint64 {
	return hashWoodSlab | (uint64(boolByte(s.Top)) << 32) | (uint64(boolByte(s.Double)) << 33) | (uint64(s.Wood.Uint8()) << 34)
}

// CanDisplace ...
func (s WoodSlab) CanDisplace(b world.Liquid) bool {
	_, ok := b.(Water)
	return !s.Double && ok
}

// SideClosed ...
func (s WoodSlab) SideClosed(pos, side world.BlockPos, _ *world.World) bool {
	// Only returns true if the side is below the slab and if the slab is not upside down.
	return !s.Top && side[1] == pos[1]-1
}
