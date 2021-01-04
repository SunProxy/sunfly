package block

import (
	"github.com/sunproxy/sunfly/dragonfly/block/colour"
	"github.com/sunproxy/sunfly/dragonfly/item"
	"github.com/sunproxy/sunfly/dragonfly/world"
)

// ConcretePowder is a gravity affected block that comes in 16 different colours. When interacting with water,
// it becomes concrete.
type ConcretePowder struct {
	noNBT
	gravityAffected
	solid
	snare

	// Colour is the colour of the concrete powder.
	Colour colour.Colour
}

// CanSolidify ...
func (c ConcretePowder) CanSolidify(pos world.BlockPos, w *world.World) bool {
	_, water := w.Block(pos).(Water)
	return water
}

// NeighbourUpdateTick ...
func (c ConcretePowder) NeighbourUpdateTick(pos, _ world.BlockPos, w *world.World) {
	for i := world.Face(0); i < 6; i++ {
		if _, ok := w.Block(pos.Side(i)).(Water); ok {
			w.PlaceBlock(pos, Concrete{Colour: c.Colour})
			return
		}
	}
	c.fall(c, pos, w)
}

// BreakInfo ...
func (c ConcretePowder) BreakInfo() BreakInfo {
	return BreakInfo{
		Hardness:    0.5,
		Harvestable: alwaysHarvestable,
		Effective:   shovelEffective,
		Drops:       simpleDrops(item.NewStack(c, 1)),
	}
}

// EncodeItem ...
func (c ConcretePowder) EncodeItem() (id int32, meta int16) {
	return 237, int16(c.Colour.Uint8())
}

// EncodeBlock ...
func (c ConcretePowder) EncodeBlock() (name string, properties map[string]interface{}) {
	return "minecraft:concrete_powder", map[string]interface{}{"color": c.Colour.String()}
}

// Hash ...
func (c ConcretePowder) Hash() uint64 {
	return hashConcretePowder | (uint64(c.Colour.Uint8()) << 32)
}
