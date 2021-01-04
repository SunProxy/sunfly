package item

import (
	"github.com/sunproxy/sunfly/dragonfly/item/armour"
	"github.com/sunproxy/sunfly/dragonfly/world"
)

// Boots are a defensive item that may be equipped in the boots armour slot. They come in several tiers, like
// leather, gold, chain, iron and diamond.
type Boots struct {
	// Tier is the tier of the boots.
	Tier armour.Tier
}

// Use handles the auto-equipping of boots in the armour slot when using it.
func (b Boots) Use(_ *world.World, user User, _ *UseContext) bool {
	if armoured, ok := user.(Armoured); ok {
		currentEquipped := armoured.Armour().Boots()

		right, left := user.HeldItems()
		armoured.Armour().SetBoots(right)
		user.SetHeldItems(currentEquipped, left)
	}
	return false
}

// MaxCount always returns 1.
func (b Boots) MaxCount() int {
	return 1
}

// DurabilityInfo ...
func (b Boots) DurabilityInfo() DurabilityInfo {
	return DurabilityInfo{
		MaxDurability: int(b.Tier.BaseDurability + b.Tier.BaseDurability/5.5),
		BrokenItem:    simpleItem(Stack{}),
	}
}

// DefencePoints ...
func (b Boots) DefencePoints() float64 {
	switch b.Tier {
	case armour.TierLeather, armour.TierGold, armour.TierChain:
		return 1
	case armour.TierIron:
		return 2
	case armour.TierDiamond, armour.TierNetherite:
		return 3
	}
	panic("invalid boots tier")
}

// KnockBackResistance ...
func (b Boots) KnockBackResistance() float64 {
	return b.Tier.KnockBackResistance
}

// EncodeItem ...
func (b Boots) EncodeItem() (id int32, meta int16) {
	switch b.Tier {
	case armour.TierLeather:
		return 301, 0
	case armour.TierGold:
		return 317, 0
	case armour.TierChain:
		return 305, 0
	case armour.TierIron:
		return 309, 0
	case armour.TierDiamond:
		return 313, 0
	case armour.TierNetherite:
		return 751, 0
	}
	panic("invalid boots tier")
}
