package block_internal

import (
	"github.com/sunproxy/sunfly/dragonfly/entity"
	"github.com/sunproxy/sunfly/dragonfly/entity/damage"
	"github.com/sunproxy/sunfly/dragonfly/world"
)

// FireDamage deals fire damage to the entity.
func FireDamage(e world.Entity, amount float64) {
	if l, ok := e.(entity.Living); ok && !l.AttackImmune() {
		l.Hurt(amount, damage.SourceFire{})
	}
}

// LavaDamage deals lava damage to the entity.
func LavaDamage(e world.Entity, amount float64) {
	if l, ok := e.(entity.Living); ok && !l.AttackImmune() {
		l.Hurt(amount, damage.SourceLava{})
	}
}
