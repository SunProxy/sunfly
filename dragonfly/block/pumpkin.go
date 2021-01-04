package block

import (
	"github.com/go-gl/mathgl/mgl64"
	"github.com/sunproxy/sunfly/dragonfly/block/instrument"
	"github.com/sunproxy/sunfly/dragonfly/item"
	"github.com/sunproxy/sunfly/dragonfly/world"
)

// Pumpkin is a crop block. Interacting with shears results in the carved variant.
type Pumpkin struct {
	noNBT
	solid

	// Carved is whether the pumpkin is carved.
	Carved bool
	// Facing is the direction the pumpkin is facing.
	Facing world.Direction
}

// Instrument ...
func (p Pumpkin) Instrument() instrument.Instrument {
	if !p.Carved {
		return instrument.Didgeridoo()
	}
	return instrument.Piano()
}

// UseOnBlock ...
func (p Pumpkin) UseOnBlock(pos world.BlockPos, face world.Face, _ mgl64.Vec3, w *world.World, user item.User, ctx *item.UseContext) (used bool) {
	pos, _, used = firstReplaceable(w, pos, face, p)
	if !used {
		return
	}
	p.Facing = user.Facing().Opposite()

	place(w, pos, p, user, ctx)
	return placed(ctx)
}

// BreakInfo ...
func (p Pumpkin) BreakInfo() BreakInfo {
	return BreakInfo{
		Hardness:    1,
		Harvestable: alwaysHarvestable,
		Effective:   axeEffective,
		Drops:       simpleDrops(item.NewStack(p, 1)),
	}
}

// EncodeItem ...
func (p Pumpkin) EncodeItem() (id int32, meta int16) {
	if p.Carved {
		return -155, 0
	}
	return 86, 0
}

// EncodeBlock ...
func (p Pumpkin) EncodeBlock() (name string, properties map[string]interface{}) {
	direction := 2
	switch p.Facing {
	case world.South:
		direction = 0
	case world.West:
		direction = 1
	case world.East:
		direction = 3
	}

	if p.Carved {
		return "minecraft:carved_pumpkin", map[string]interface{}{"direction": int32(direction)}
	}
	return "minecraft:pumpkin", map[string]interface{}{"direction": int32(direction)}
}

// Hash ...
func (p Pumpkin) Hash() uint64 {
	return hashPumpkin | (uint64(boolByte(p.Carved)) << 32) | (uint64(p.Facing) << 33)
}
