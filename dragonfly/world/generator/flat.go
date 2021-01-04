package generator

import (
	"github.com/sunproxy/sunfly/dragonfly/block"
	"github.com/sunproxy/sunfly/dragonfly/world"
	"github.com/sunproxy/sunfly/dragonfly/world/chunk"
)

// Flat is the flat generator of World. It generates flat worlds (like those in vanilla) with no other
// decoration.
type Flat struct{}

var (
	grassBlock, _ = world.BlockRuntimeID(block.Grass{})
	dirt, _       = world.BlockRuntimeID(block.Dirt{})
	bedrock, _    = world.BlockRuntimeID(block.Bedrock{})
)

// GenerateChunk ...
func (Flat) GenerateChunk(_ world.ChunkPos, chunk *chunk.Chunk) {
	for x := uint8(0); x < 16; x++ {
		for z := uint8(0); z < 16; z++ {
			chunk.SetRuntimeID(x, 0, z, 0, bedrock)
			chunk.SetRuntimeID(x, 1, z, 0, dirt)
			chunk.SetRuntimeID(x, 2, z, 0, dirt)
			chunk.SetRuntimeID(x, 3, z, 0, grassBlock)
		}
	}
}
