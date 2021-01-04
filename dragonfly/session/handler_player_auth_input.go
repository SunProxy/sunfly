package session

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/go-gl/mathgl/mgl64"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
	"github.com/sunproxy/sunfly/dragonfly/world"
)

// PlayerAuthInputHandler handles the PlayerAuthInput packet.
type PlayerAuthInputHandler struct{}

// Handle ...
func (h PlayerAuthInputHandler) Handle(p packet.Packet, s *Session) error {
	pk := p.(*packet.PlayerAuthInput)
	pk.Position = pk.Position.Sub(mgl32.Vec3{0, 1.62}) // Subtract the base offset of players from the pos.

	newPos := vec32To64(pk.Position)
	deltaPos, deltaYaw, deltaPitch := newPos.Sub(s.c.Position()), float64(pk.Yaw)-s.c.Yaw(), float64(pk.Pitch)-s.c.Pitch()
	if mgl64.FloatEqual(deltaPos.Len(), 0) && mgl64.FloatEqual(deltaYaw, 0) && mgl64.FloatEqual(deltaPitch, 0) {
		// The PlayerAuthInput packet is sent every tick, so don't do anything if the position and rotation
		// were unchanged.
		return nil
	}

	s.teleportMu.Lock()
	teleportPos := s.teleportPos
	s.teleportMu.Unlock()
	if teleportPos != nil {
		if newPos.Sub(*teleportPos).Len() > 2 {
			// The player has moved before it received the teleport packet. Ignore this movement entirely and
			// wait for the client to sync itself back to the server. Once we get a movement that is close
			// enough to the teleport position, we'll allow the player to move around again.
			return nil
		}
		s.teleportMu.Lock()
		s.teleportPos = nil
		s.teleportMu.Unlock()
	}

	_, submergedBefore := s.c.World().Liquid(world.BlockPosFromVec3(s.c.Position().Add(mgl64.Vec3{0, s.c.EyeHeight()})))

	s.c.Move(deltaPos)
	s.c.Rotate(deltaYaw, deltaPitch)

	_, submergedAfter := s.c.World().Liquid(world.BlockPosFromVec3(s.c.Position().Add(mgl64.Vec3{0, s.c.EyeHeight()})))

	if submergedBefore != submergedAfter {
		// Player wasn't either breathing before and no longer isn't, or wasn't breathing before and now is,
		// so send the updated metadata.
		s.ViewEntityState(s.c, s.c.State())
	}

	s.chunkLoader.Move(s.c.Position())
	s.writePacket(&packet.NetworkChunkPublisherUpdate{
		Position: protocol.BlockPos{int32(pk.Position[0]), int32(pk.Position[1]), int32(pk.Position[2])},
		Radius:   uint32(s.chunkRadius * 16),
	})
	return nil
}
