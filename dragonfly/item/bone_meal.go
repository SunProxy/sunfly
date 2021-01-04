package item

import (
	"github.com/go-gl/mathgl/mgl64"
	"github.com/sunproxy/sunfly/dragonfly/internal/item_internal"
	"github.com/sunproxy/sunfly/dragonfly/world"
	"github.com/sunproxy/sunfly/dragonfly/world/particle"
)

// BoneMeal is an item used to force growth in plants & crops.
type BoneMeal struct{}

// UseOnBlock ...
func (b BoneMeal) UseOnBlock(pos world.BlockPos, _ world.Face, _ mgl64.Vec3, w *world.World, _ User, ctx *UseContext) bool {
	ok := item_internal.BoneMeal(pos, w)
	if ok {
		ctx.CountSub = 1
		w.AddParticle(pos.Vec3(), particle.Bonemeal{})
	}
	return ok
}

// EncodeItem ...
func (b BoneMeal) EncodeItem() (id int32, meta int16) {
	return 351, 15
}
