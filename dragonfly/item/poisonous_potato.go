package item

import (
	"github.com/sunproxy/sunfly/dragonfly/entity/effect"
	"github.com/sunproxy/sunfly/dragonfly/world"
	"math/rand"
	"time"
)

// PoisonousPotato is a type of potato that can poison the player.
type PoisonousPotato struct{}

// AlwaysConsumable ...
func (p PoisonousPotato) AlwaysConsumable() bool {
	return false
}

// ConsumeDuration ...
func (p PoisonousPotato) ConsumeDuration() time.Duration {
	return DefaultConsumeDuration
}

// Consume ...
func (p PoisonousPotato) Consume(_ *world.World, c Consumer) Stack {
	c.Saturate(2, 1.2)
	if rand.Float64() < 0.6 {
		c.AddEffect(effect.Poison{}.WithSettings(time.Duration(5)*time.Second, 1, false))
	}
	return Stack{}
}

// EncodeItem ...
func (p PoisonousPotato) EncodeItem() (id int32, meta int16) {
	return 394, 0
}
