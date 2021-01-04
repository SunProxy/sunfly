package session

import (
	"github.com/sunproxy/sunfly/dragonfly/world"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

// BlockPickRequestHandler handles the BlockPickRequest packet.
type BlockPickRequestHandler struct{}

// Handle ...
func (b BlockPickRequestHandler) Handle(p packet.Packet, s *Session) error {
	pk := p.(*packet.BlockPickRequest)
	s.c.PickBlock(world.BlockPos{int(pk.Position.X()), int(pk.Position.Y()), int(pk.Position.Z())})
	return nil
}
