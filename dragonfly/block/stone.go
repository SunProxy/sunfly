package block

import (
	"github.com/sunproxy/sunfly/dragonfly/item"
)

type (
	// Stone is a block found underground in the world or on mountains.
	Stone struct {
		noNBT
		solid
		bassDrum

		// Smooth specifies if the stone is its smooth variant.
		Smooth bool
	}

	// Granite is a type of igneous rock.
	Granite polishable
	// Diorite is a type of igneous rock.
	Diorite polishable
	// Andesite is a type of igneous rock.
	Andesite polishable

	// polishable forms the base of blocks that may be polished.
	polishable struct {
		noNBT
		solid
		bassDrum
		// Polished specifies if the block is polished or not. When set to true, the block will represent its
		// polished variant, for example polished andesite.
		Polished bool
	}
)

var stoneBreakInfo = BreakInfo{
	Hardness:    1.5,
	Harvestable: pickaxeHarvestable,
	Effective:   pickaxeEffective,
	Drops:       simpleDrops(item.NewStack(Cobblestone{}, 1)),
}

// BreakInfo ...
func (s Stone) BreakInfo() BreakInfo {
	breakInfo := stoneBreakInfo
	if s.Smooth {
		breakInfo.Hardness = 2
		breakInfo.Drops = simpleDrops(item.NewStack(s, 1))
	}
	return breakInfo
}

// BreakInfo ...
func (g Granite) BreakInfo() BreakInfo {
	i := stoneBreakInfo
	i.Drops = simpleDrops(item.NewStack(g, 1))
	return i
}

// BreakInfo ...
func (d Diorite) BreakInfo() BreakInfo {
	i := stoneBreakInfo
	i.Drops = simpleDrops(item.NewStack(d, 1))
	return i
}

// BreakInfo ...
func (a Andesite) BreakInfo() BreakInfo {
	i := stoneBreakInfo
	i.Drops = simpleDrops(item.NewStack(a, 1))
	return i
}

// EncodeItem ...
func (s Stone) EncodeItem() (id int32, meta int16) {
	if s.Smooth {
		return -183, 0
	}
	return 1, 0
}

// EncodeItem ...
func (a Andesite) EncodeItem() (id int32, meta int16) {
	if a.Polished {
		return 1, 6
	}
	return 1, 5
}

// EncodeItem ...
func (d Diorite) EncodeItem() (id int32, meta int16) {
	if d.Polished {
		return 1, 4
	}
	return 1, 3
}

// EncodeItem ...
func (g Granite) EncodeItem() (id int32, meta int16) {
	if g.Polished {
		return 1, 2
	}
	return 1, 1
}
