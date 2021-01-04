package block

import (
	"github.com/sunproxy/sunfly/dragonfly/block/colour"
	"github.com/sunproxy/sunfly/dragonfly/block/fire"
	"github.com/sunproxy/sunfly/dragonfly/block/wood"
	"github.com/sunproxy/sunfly/dragonfly/internal/entity_internal"
	"github.com/sunproxy/sunfly/dragonfly/internal/item_internal"
	"github.com/sunproxy/sunfly/dragonfly/world"
	_ "unsafe" // Imported for compiler directives.
)

// init registers all blocks implemented by Dragonfly.
func init() {
	_ = world.RegisterBlock(Air{}, world.BlockState{Name: "minecraft:air"})
	_ = world.RegisterBlock(Stone{}, world.BlockState{Name: "minecraft:stone", Properties: map[string]interface{}{"stone_type": "stone"}})
	_ = world.RegisterBlock(Stone{Smooth: true}, world.BlockState{Name: "minecraft:smooth_stone"})
	_ = world.RegisterBlock(Granite{}, world.BlockState{Name: "minecraft:stone", Properties: map[string]interface{}{"stone_type": "granite"}})
	_ = world.RegisterBlock(Granite{Polished: true}, world.BlockState{Name: "minecraft:stone", Properties: map[string]interface{}{"stone_type": "granite_smooth"}})
	_ = world.RegisterBlock(Diorite{}, world.BlockState{Name: "minecraft:stone", Properties: map[string]interface{}{"stone_type": "diorite"}})
	_ = world.RegisterBlock(Diorite{Polished: true}, world.BlockState{Name: "minecraft:stone", Properties: map[string]interface{}{"stone_type": "diorite_smooth"}})
	_ = world.RegisterBlock(Andesite{}, world.BlockState{Name: "minecraft:stone", Properties: map[string]interface{}{"stone_type": "andesite"}})
	_ = world.RegisterBlock(Andesite{Polished: true}, world.BlockState{Name: "minecraft:stone", Properties: map[string]interface{}{"stone_type": "andesite_smooth"}})
	_ = world.RegisterBlock(Grass{}, world.BlockState{Name: "minecraft:grass"})
	_ = world.RegisterBlock(GrassPath{}, world.BlockState{Name: "minecraft:grass_path"})
	_ = world.RegisterBlock(Dirt{}, world.BlockState{Name: "minecraft:dirt", Properties: map[string]interface{}{"dirt_type": "normal"}})
	_ = world.RegisterBlock(Dirt{Coarse: true}, world.BlockState{Name: "minecraft:dirt", Properties: map[string]interface{}{"dirt_type": "coarse"}})
	_ = world.RegisterBlock(Cobblestone{}, world.BlockState{Name: "minecraft:cobblestone"})
	_ = world.RegisterBlock(Cobblestone{Mossy: true}, world.BlockState{Name: "minecraft:mossy_cobblestone"})
	_ = world.RegisterBlock(Bedrock{}, world.BlockState{Name: "minecraft:bedrock", Properties: map[string]interface{}{"infiniburn_bit": false}})
	_ = world.RegisterBlock(Bedrock{InfiniteBurning: true}, world.BlockState{Name: "minecraft:bedrock", Properties: map[string]interface{}{"infiniburn_bit": true}})
	_ = world.RegisterBlock(Obsidian{}, world.BlockState{Name: "minecraft:obsidian"})
	_ = world.RegisterBlock(DiamondBlock{}, world.BlockState{Name: "minecraft:diamond_block"})
	_ = world.RegisterBlock(Glass{}, world.BlockState{Name: "minecraft:glass"})
	_ = world.RegisterBlock(Glowstone{}, world.BlockState{Name: "minecraft:glowstone"})
	_ = world.RegisterBlock(EmeraldBlock{}, world.BlockState{Name: "minecraft:emerald_block"})
	_ = world.RegisterBlock(EndBricks{}, world.BlockState{Name: "minecraft:end_bricks"})
	_ = world.RegisterBlock(GoldBlock{}, world.BlockState{Name: "minecraft:gold_block"})
	_ = world.RegisterBlock(NetheriteBlock{}, world.BlockState{Name: "minecraft:netherite_block"})
	_ = world.RegisterBlock(IronBlock{}, world.BlockState{Name: "minecraft:iron_block"})
	_ = world.RegisterBlock(CoalBlock{}, world.BlockState{Name: "minecraft:coal_block"})
	_ = world.RegisterBlock(Beacon{}, world.BlockState{Name: "minecraft:beacon"})
	_ = world.RegisterBlock(Sponge{}, world.BlockState{Name: "minecraft:sponge", Properties: map[string]interface{}{"sponge_type": "dry"}})
	_ = world.RegisterBlock(Sponge{Wet: true}, world.BlockState{Name: "minecraft:sponge", Properties: map[string]interface{}{"sponge_type": "wet"}})
	_ = world.RegisterBlock(LapisBlock{}, world.BlockState{Name: "minecraft:lapis_block"})
	_ = world.RegisterBlock(Terracotta{}, world.BlockState{Name: "minecraft:hardened_clay"})
	_ = world.RegisterBlock(GlassPane{}, world.BlockState{Name: "minecraft:glass_pane"})
	_ = world.RegisterBlock(IronBars{}, world.BlockState{Name: "minecraft:iron_bars"})
	_ = world.RegisterBlock(NetherBrickFence{}, world.BlockState{Name: "minecraft:nether_brick_fence"})
	_ = world.RegisterBlock(EndStone{}, world.BlockState{Name: "minecraft:end_stone"})
	_ = world.RegisterBlock(Netherrack{}, world.BlockState{Name: "minecraft:netherrack"})
	_ = world.RegisterBlock(Quartz{Smooth: true}, world.BlockState{Name: "minecraft:quartz_block", Properties: map[string]interface{}{"chisel_type": "smooth", "pillar_axis": "x"}})
	_ = world.RegisterBlock(Quartz{Smooth: false}, world.BlockState{Name: "minecraft:quartz_block", Properties: map[string]interface{}{"chisel_type": "default", "pillar_axis": "x"}})
	_ = world.RegisterBlock(ChiseledQuartz{}, world.BlockState{Name: "minecraft:quartz_block", Properties: map[string]interface{}{"chisel_type": "chiseled", "pillar_axis": "x"}})
	_ = world.RegisterBlock(QuartzBricks{}, world.BlockState{Name: "minecraft:quartz_bricks"})
	_ = world.RegisterBlock(Clay{}, world.BlockState{Name: "minecraft:clay"})
	_ = world.RegisterBlock(Lantern{Type: fire.Normal(), Hanging: true}, world.BlockState{Name: "minecraft:lantern", Properties: map[string]interface{}{"hanging": true}})
	_ = world.RegisterBlock(Lantern{Type: fire.Normal(), Hanging: true}, world.BlockState{Name: "minecraft:soul_lantern", Properties: map[string]interface{}{"hanging": true}})
	_ = world.RegisterBlock(Lantern{Type: fire.Normal(), Hanging: true}, world.BlockState{Name: "minecraft:lantern", Properties: map[string]interface{}{"hanging": false}})
	_ = world.RegisterBlock(Lantern{Type: fire.Normal(), Hanging: true}, world.BlockState{Name: "minecraft:soul_lantern", Properties: map[string]interface{}{"hanging": false}})
	_ = world.RegisterBlock(AncientDebris{}, world.BlockState{Name: "minecraft:ancient_debris"})
	_ = world.RegisterBlock(EmeraldOre{}, world.BlockState{Name: "minecraft:emerald_ore"})
	_ = world.RegisterBlock(DiamondOre{}, world.BlockState{Name: "minecraft:diamond_ore"})
	_ = world.RegisterBlock(LapisOre{}, world.BlockState{Name: "minecraft:lapis_ore"})
	_ = world.RegisterBlock(NetherGoldOre{}, world.BlockState{Name: "minecraft:nether_gold_ore"})
	_ = world.RegisterBlock(GoldOre{}, world.BlockState{Name: "minecraft:gold_ore"})
	_ = world.RegisterBlock(IronOre{}, world.BlockState{Name: "minecraft:iron_ore"})
	_ = world.RegisterBlock(CoalOre{}, world.BlockState{Name: "minecraft:coal_ore"})
	_ = world.RegisterBlock(NetherQuartzOre{}, world.BlockState{Name: "minecraft:quartz_ore"})
	_ = world.RegisterBlock(Melon{}, world.BlockState{Name: "minecraft:melon_block"})
	_ = world.RegisterBlock(Sand{}, world.BlockState{Name: "minecraft:sand", Properties: map[string]interface{}{"sand_type": "normal"}})
	_ = world.RegisterBlock(Sand{Red: true}, world.BlockState{Name: "minecraft:sand", Properties: map[string]interface{}{"sand_type": "red"}})
	_ = world.RegisterBlock(Gravel{}, world.BlockState{Name: "minecraft:gravel"})
	_ = world.RegisterBlock(Bricks{}, world.BlockState{Name: "minecraft:brick_block"})
	_ = world.RegisterBlock(SoulSand{}, world.BlockState{Name: "minecraft:soul_sand"})
	_ = world.RegisterBlock(Barrier{}, world.BlockState{Name: "minecraft:barrier"})
	_ = world.RegisterBlock(CryingObsidian{}, world.BlockState{Name: "minecraft:crying_obsidian"})
	_ = world.RegisterBlock(SeaLantern{}, world.BlockState{Name: "minecraft:seaLantern"})
	_ = world.RegisterBlock(SoulSoil{}, world.BlockState{Name: "minecraft:soul_soil"})
	_ = world.RegisterBlock(BlueIce{}, world.BlockState{Name: "minecraft:blue_ice"})
	_ = world.RegisterBlock(GildedBlackstone{}, world.BlockState{Name: "minecraft:gilded_blackstone"})
	_ = world.RegisterBlock(Shroomlight{}, world.BlockState{Name: "minecraft:shroomlight"})
	_ = world.RegisterBlock(InvisibleBedrock{}, world.BlockState{Name: "minecraft:invisibleBedrock"})
	_ = world.RegisterBlock(NoteBlock{}, world.BlockState{Name: "minecraft:noteblock"})
	_ = world.RegisterBlock(DragonEgg{}, world.BlockState{Name: "minecraft:dragon_egg"})
	_ = world.RegisterBlock(BoneBlock{Axis: world.X}, world.BlockState{Name: "minecraft:bone_block", Properties: map[string]interface{}{"pillar_axis": world.X.String(), "deprecated": int32(0)}})
	_ = world.RegisterBlock(BoneBlock{Axis: world.Y}, world.BlockState{Name: "minecraft:bone_block", Properties: map[string]interface{}{"pillar_axis": world.Y.String(), "deprecated": int32(0)}})
	_ = world.RegisterBlock(BoneBlock{Axis: world.Z}, world.BlockState{Name: "minecraft:bone_block", Properties: map[string]interface{}{"pillar_axis": world.Z.String(), "deprecated": int32(0)}})

	// Colour block implementations
	for _, c := range colour.All() {
		_ = world.RegisterBlock(Carpet{Colour: c}, world.BlockState{Name: "minecraft:carpet", Properties: map[string]interface{}{"color": c.String()}})
		_ = world.RegisterBlock(Concrete{Colour: c}, world.BlockState{Name: "minecraft:concrete", Properties: map[string]interface{}{"color": c.String()}})
		_ = world.RegisterBlock(ConcretePowder{Colour: c}, world.BlockState{Name: "minecraft:concretePowder", Properties: map[string]interface{}{"color": c.String()}})
		_ = world.RegisterBlock(Wool{Colour: c}, world.BlockState{Name: "minecraft:wool", Properties: map[string]interface{}{"color": c.String()}})
		colourName := c.String()
		if c == colour.LightGrey() {
			// Light grey is actually called "silver" in the block state. Mojang pls.
			colourName = "silver"
		}

		for _, d := range world.AllDirections() {
			_ = world.RegisterBlock(GlazedTerracotta{Colour: c}, world.BlockState{Name: "minecraft:" + colourName + "_glazed_terracotta", Properties: map[string]interface{}{"facing_direction": int32(d)}})
		}

		_ = world.RegisterBlock(StainedGlass{Colour: c}, world.BlockState{Name: "minecraft:stained_glass", Properties: map[string]interface{}{"color": colourName}})
		_ = world.RegisterBlock(StainedGlassPane{Colour: c}, world.BlockState{Name: "minecraft:stained_glass_pane", Properties: map[string]interface{}{"color": colourName}})
		_ = world.RegisterBlock(StainedTerracotta{Colour: c}, world.BlockState{Name: "minecraft:stained_hardened_clay", Properties: map[string]interface{}{"color": colourName}})
	}

	// Wood implementation
	for _, w := range wood.All() {
		if w == wood.Acacia() || w == wood.DarkOak() {
			_ = world.RegisterBlock(WoodFence{Wood: w}, world.BlockState{Name: "minecraft:fence", Properties: map[string]interface{}{"wood_type": w.String()}})
			for i := 0; i < 3; i++ {
				_ = world.RegisterBlock(Log{Wood: w, Stripped: false, Axis: world.Axis(i)}, world.BlockState{Name: "minecraft:log2", Properties: map[string]interface{}{"new_log_type": w.String(), "pillar_axis": world.Axis(i).String()}, Version: 17825808})
				_ = world.RegisterBlock(Log{Wood: w, Stripped: true, Axis: world.Axis(i)}, world.BlockState{Name: "minecraft:stripped_" + w.String() + "_log", Properties: map[string]interface{}{"pillar_axis": world.Axis(i).String()}, Version: 17825808})
			}
			_ = world.RegisterBlock(Planks{Wood: w}, world.BlockState{Name: "minecraft:planks", Properties: map[string]interface{}{"wood_type": w.String()}, Version: 17825808})
			for i := 0; i < 2; i++ {
				for j := 0; j < 2; j++ {
					_ = world.RegisterBlock(Leaves{Wood: w, Persistent: i != 0, shouldUpdate: j != 0}, world.BlockState{Name: "minecraft:leaves2", Properties: map[string]interface{}{"new_leaf_type": w.String(), "persistent_bit": i != 0, "update_bit": j != 0}, Version: 17825808})
				}
			}
		} else if w == wood.Crimson() || w == wood.Warped() {
			_ = world.RegisterBlock(WoodFence{Wood: w}, world.BlockState{Name: "minecraft:" + w.String() + "_fence"})
			// TODO (Civiled): Register warped and crimson wood blocks
			// TODO (Civiled): Register doors
		} else {
			_ = world.RegisterBlock(WoodFence{Wood: w}, world.BlockState{Name: "minecraft:fence", Properties: map[string]interface{}{"wood_type": w.String()}})
			for i := 0; i < 3; i++ {
				_ = world.RegisterBlock(Log{Wood: w, Stripped: false, Axis: world.Axis(i)}, world.BlockState{Name: "minecraft:log", Properties: map[string]interface{}{"old_log_type": w.String(), "pillar_axis": world.Axis(i).String()}, Version: 17825808})
				_ = world.RegisterBlock(Log{Wood: w, Stripped: true, Axis: world.Axis(i)}, world.BlockState{Name: "minecraft:stripped_" + w.String() + "_log", Properties: map[string]interface{}{"pillar_axis": world.Axis(i).String()}, Version: 17825808})
			}
			_ = world.RegisterBlock(Planks{Wood: w}, world.BlockState{Name: "minecraft:planks", Properties: map[string]interface{}{"wood_type": w.String()}, Version: 17825808})
			for i := 0; i < 2; i++ {
				for j := 0; j < 2; j++ {
					_ = world.RegisterBlock(Leaves{Wood: w, Persistent: i != 0, shouldUpdate: j != 0}, world.BlockState{Name: "minecraft:leaves", Properties: map[string]interface{}{"old_leaf_type": w.String(), "persistent_bit": i != 0, "update_bit": j != 0}, Version: 17825808})
				}
			}
		}

	}

	// Directional block implementation
	for _, d := range world.AllDirections() {
		for i := 0; i < 3; i++ {

			_ = world.RegisterBlock(CocoaBean{Age: i, Facing: d}, world.BlockState{Name: "minecraft:cocoa", Properties: map[string]interface{}{"age": i, "direction": int32(d)}})
		}
		_ = world.RegisterBlock(LitPumpkin{Facing: d}, world.BlockState{Name: "minecraft:lit_pumpkin", Properties: map[string]interface{}{"direction": int32(d)}})
		_ = world.RegisterBlock(Chest{Facing: d}, world.BlockState{Name: "minecraft:chest", Properties: map[string]interface{}{"facing_direction": 2 + int32(d)}, Version: 17825808})
		_ = world.RegisterBlock(Pumpkin{Facing: d}, world.BlockState{Name: "minecraft:pumpkin", Properties: map[string]interface{}{"direction": int32(d)}})
		_ = world.RegisterBlock(Pumpkin{Facing: d}, world.BlockState{Name: "minecraft:carved_pumpkin", Properties: map[string]interface{}{"direction": int32(d)}})
	}

	// Fire implementation
	for f := 0; f < 16; f++ {
		_ = world.RegisterBlock(Fire{Age: f, Type: fire.Normal()}, world.BlockState{Name: "minecraft:fire", Properties: map[string]interface{}{"age": f}})
		_ = world.RegisterBlock(Fire{Age: f, Type: fire.Normal()}, world.BlockState{Name: "minecraft:soul_fire", Properties: map[string]interface{}{"age": f}})
	}

	// Beetroot/Carrot/Melon/Potato/Wheat implementation
	for g := 0; g < 8; g++ {
		_ = world.RegisterBlock(BeetrootSeeds{}, world.BlockState{Name: "minecraft:beetroot", Properties: map[string]interface{}{"growth": g}})
		_ = world.RegisterBlock(Carrot{}, world.BlockState{Name: "minecraft:carrot", Properties: map[string]interface{}{"growth": g}})
		_ = world.RegisterBlock(MelonSeeds{}, world.BlockState{Name: "minecraft:melon_stem", Properties: map[string]interface{}{"growth": g}})
		_ = world.RegisterBlock(Potato{}, world.BlockState{Name: "minecraft:potatoes", Properties: map[string]interface{}{"growth": int32(g)}})
		_ = world.RegisterBlock(WheatSeeds{}, world.BlockState{Name: "minecraft:wheat", Properties: map[string]interface{}{"growth": int32(g)}})
	}

	// Cake/Farmland implementation
	for b := 0; b < 7; b++ {
		_ = world.RegisterBlock(Cake{Bites: b}, world.BlockState{Name: "minecraft:cake", Properties: map[string]interface{}{"bite_counter": b}})
		_ = world.RegisterBlock(Farmland{Hydration: b}, world.BlockState{Name: "minecraft:farmland", Properties: map[string]interface{}{"moisturized_amount": b}})
	}

	// Kelp implementation
	for k := 0; k < 26; k++ {
		_ = world.RegisterBlock(Kelp{Age: k}, world.BlockState{Name: "minecraft:kelp", Properties: map[string]interface{}{"kelp_age": k}})
	}

	// Chiseled quartz implementation
	for dir := 0; dir < 3; dir++ {
		_ = world.RegisterBlock(QuartzPillar{Axis: world.Axis(dir)}, world.BlockState{Name: "minecraft:quartz_block", Properties: map[string]interface{}{"pillar_axis": world.Axis(dir), "chisel_type": "lines"}})
		_ = world.RegisterBlock(Basalt{Axis: world.Axis(dir), Polished: true}, world.BlockState{Name: "minecraft:polished_basalt", Properties: map[string]interface{}{"pillar_axis": world.Axis(dir)}})
		_ = world.RegisterBlock(Basalt{Axis: world.Axis(dir), Polished: false}, world.BlockState{Name: "minecraft:basalt", Properties: map[string]interface{}{"pillar_axis": world.Axis(dir)}})
	}

	// Nether Wart implementation
	for n := 0; n < 4; n++ {
		_ = world.RegisterBlock(NetherWart{Age: n}, world.BlockState{Name: "minecraft:nether_wart", Properties: map[string]interface{}{"age": n}})
	}
}

func init() {
	world.RegisterItem("minecraft:air", Air{})
	world.RegisterItem("minecraft:stone", Stone{})
	world.RegisterItem("minecraft:smooth_stone", Stone{Smooth: true})
	world.RegisterItem("minecraft:stone", Granite{})
	world.RegisterItem("minecraft:stone", Granite{Polished: true})
	world.RegisterItem("minecraft:stone", Diorite{})
	world.RegisterItem("minecraft:stone", Diorite{Polished: true})
	world.RegisterItem("minecraft:stone", Andesite{})
	world.RegisterItem("minecraft:stone", Andesite{Polished: true})
	world.RegisterItem("minecraft:grass", Grass{})
	world.RegisterItem("minecraft:grass_path", GrassPath{})
	world.RegisterItem("minecraft:dirt", Dirt{})
	world.RegisterItem("minecraft:dirt", Dirt{Coarse: true})
	world.RegisterItem("minecraft:cobblestone", Cobblestone{})
	world.RegisterItem("minecraft:bedrock", Bedrock{})
	world.RegisterItem("minecraft:kelp", Kelp{})
	world.RegisterItem("minecraft:log", Log{Wood: wood.Oak()})
	world.RegisterItem("minecraft:log", Log{Wood: wood.Spruce()})
	world.RegisterItem("minecraft:log", Log{Wood: wood.Birch()})
	world.RegisterItem("minecraft:log", Log{Wood: wood.Jungle()})
	world.RegisterItem("minecraft:leaves", Leaves{Wood: wood.Oak(), Persistent: true})
	world.RegisterItem("minecraft:leaves", Leaves{Wood: wood.Spruce(), Persistent: true})
	world.RegisterItem("minecraft:leaves", Leaves{Wood: wood.Birch(), Persistent: true})
	world.RegisterItem("minecraft:leaves", Leaves{Wood: wood.Jungle(), Persistent: true})
	world.RegisterItem("minecraft:chest", Chest{})
	world.RegisterItem("minecraft:mossy_cobblestone", Cobblestone{Mossy: true})
	world.RegisterItem("minecraft:leaves2", Leaves{Wood: wood.Acacia(), Persistent: true})
	world.RegisterItem("minecraft:leaves2", Leaves{Wood: wood.DarkOak(), Persistent: true})
	world.RegisterItem("minecraft:log2", Log{Wood: wood.Acacia()})
	world.RegisterItem("minecraft:log2", Log{Wood: wood.DarkOak()})
	world.RegisterItem("minecraft:stripped_spruce_log", Log{Wood: wood.Spruce(), Stripped: true})
	world.RegisterItem("minecraft:stripped_birch_log", Log{Wood: wood.Birch(), Stripped: true})
	world.RegisterItem("minecraft:stripped_jungle_log", Log{Wood: wood.Jungle(), Stripped: true})
	world.RegisterItem("minecraft:stripped_acacia_log", Log{Wood: wood.Acacia(), Stripped: true})
	world.RegisterItem("minecraft:stripped_dark_oak_log", Log{Wood: wood.DarkOak(), Stripped: true})
	world.RegisterItem("minecraft:stripped_oak_log", Log{Wood: wood.Oak(), Stripped: true})
	for _, c := range colour.All() {
		world.RegisterItem("minecraft:concrete", Concrete{Colour: c})
		world.RegisterItem("minecraft:concrete_powder", ConcretePowder{Colour: c})
		world.RegisterItem("minecraft:stained_hardened_clay", StainedTerracotta{Colour: c})
		world.RegisterItem("minecraft:carpet", Carpet{Colour: c})
		world.RegisterItem("minecraft:wool", Wool{Colour: c})
		world.RegisterItem("minecraft:stained_glass", StainedGlass{Colour: c})
		world.RegisterItem("minecraft:stained_glass_pane", StainedGlassPane{Colour: c})

		colourName := c.String()
		if c == colour.LightGrey() {
			colourName = "silver"
		}

		world.RegisterItem("minecraft:"+colourName+"_glazed_terracotta", GlazedTerracotta{Colour: c})
	}
	for _, b := range allLight() {
		world.RegisterItem("minecraft:light_block", b.(world.Item))
	}
	for _, w := range wood.All() {
		if w == wood.Crimson() || w == wood.Warped() {
			world.RegisterItem("minecraft:"+w.String()+"_planks", Planks{Wood: w})
		} else {
			world.RegisterItem("minecraft:planks", Planks{Wood: w})
		}
	}
	world.RegisterItem("minecraft:oak_stairs", WoodStairs{Wood: wood.Oak()})
	world.RegisterItem("minecraft:spruce_stairs", WoodStairs{Wood: wood.Spruce()})
	world.RegisterItem("minecraft:birch_stairs", WoodStairs{Wood: wood.Birch()})
	world.RegisterItem("minecraft:jungle_stairs", WoodStairs{Wood: wood.Jungle()})
	world.RegisterItem("minecraft:acacia_stairs", WoodStairs{Wood: wood.Acacia()})
	world.RegisterItem("minecraft:dark_oak_stairs", WoodStairs{Wood: wood.DarkOak()})
	world.RegisterItem("minecraft:crimson_stairs", WoodStairs{Wood: wood.Crimson()})
	world.RegisterItem("minecraft:warped_stairs", WoodStairs{Wood: wood.Warped()})
	world.RegisterItem("minecraft:wooden_slab", WoodSlab{Wood: wood.Oak()})
	world.RegisterItem("minecraft:wooden_slab", WoodSlab{Wood: wood.Spruce()})
	world.RegisterItem("minecraft:wooden_slab", WoodSlab{Wood: wood.Birch()})
	world.RegisterItem("minecraft:wooden_slab", WoodSlab{Wood: wood.Jungle()})
	world.RegisterItem("minecraft:wooden_slab", WoodSlab{Wood: wood.Acacia()})
	world.RegisterItem("minecraft:wooden_slab", WoodSlab{Wood: wood.DarkOak()})
	world.RegisterItem("minecraft:crimson_slab", WoodSlab{Wood: wood.Crimson()})
	world.RegisterItem("minecraft:warped_slab", WoodSlab{Wood: wood.Warped()})
	world.RegisterItem("minecraft:double_wooden_slab", WoodSlab{Wood: wood.Oak(), Double: true})
	world.RegisterItem("minecraft:double_wooden_slab", WoodSlab{Wood: wood.Spruce(), Double: true})
	world.RegisterItem("minecraft:double_wooden_slab", WoodSlab{Wood: wood.Birch(), Double: true})
	world.RegisterItem("minecraft:double_wooden_slab", WoodSlab{Wood: wood.Jungle(), Double: true})
	world.RegisterItem("minecraft:double_wooden_slab", WoodSlab{Wood: wood.Acacia(), Double: true})
	world.RegisterItem("minecraft:double_wooden_slab", WoodSlab{Wood: wood.DarkOak(), Double: true})
	world.RegisterItem("minecraft:crimson_double_slab", WoodSlab{Wood: wood.Crimson(), Double: true})
	world.RegisterItem("minecraft:warped_double_slab", WoodSlab{Wood: wood.Warped(), Double: true})
	world.RegisterItem("minecraft:obsidian", Obsidian{})
	world.RegisterItem("minecraft:diamond_block", DiamondBlock{})
	world.RegisterItem("minecraft:glass", Glass{})
	world.RegisterItem("minecraft:glowstone", Glowstone{})
	world.RegisterItem("minecraft:emerald_block", EmeraldBlock{})
	world.RegisterItem("minecraft:end_bricks", EndBricks{})
	world.RegisterItem("minecraft:end_brick_stairs", EndBrickStairs{})
	world.RegisterItem("minecraft:netherite_block", NetheriteBlock{})
	world.RegisterItem("minecraft:gold_block", GoldBlock{})
	world.RegisterItem("minecraft:iron_block", IronBlock{})
	world.RegisterItem("minecraft:coal_block", CoalBlock{})
	world.RegisterItem("minecraft:beacon", Beacon{})
	world.RegisterItem("minecraft:sponge", Sponge{})
	world.RegisterItem("minecraft:wet_sponge", Sponge{Wet: true})
	world.RegisterItem("minecraft:lapis_block", LapisBlock{})
	world.RegisterItem("minecraft:hardened_clay", Terracotta{})
	world.RegisterItem("minecraft:quartz_block", Quartz{})
	world.RegisterItem("minecraft:quartz_block", Quartz{Smooth: true})
	world.RegisterItem("minecraft:quartz_block", ChiseledQuartz{})
	world.RegisterItem("minecraft:quartz_block", QuartzPillar{})
	world.RegisterItem("minecraft:quartz_bricks", QuartzBricks{})
	world.RegisterItem("minecraft:glass_pane", GlassPane{})
	world.RegisterItem("minecraft:iron_bars", IronBars{})
	world.RegisterItem("minecraft:fence", WoodFence{Wood: wood.Oak()})
	world.RegisterItem("minecraft:fence", WoodFence{Wood: wood.Spruce()})
	world.RegisterItem("minecraft:fence", WoodFence{Wood: wood.Birch()})
	world.RegisterItem("minecraft:fence", WoodFence{Wood: wood.Jungle()})
	world.RegisterItem("minecraft:fence", WoodFence{Wood: wood.Acacia()})
	world.RegisterItem("minecraft:fence", WoodFence{Wood: wood.DarkOak()})
	world.RegisterItem("minecraft:crimson_fence", WoodFence{Wood: wood.Crimson()})
	world.RegisterItem("minecraft:warped_fence", WoodFence{Wood: wood.Warped()})
	world.RegisterItem("minecraft:nether_brick_fence", NetherBrickFence{})
	world.RegisterItem("minecraft:fence_gate", WoodFenceGate{Wood: wood.Oak()})
	world.RegisterItem("minecraft:spruce_fence_gate", WoodFenceGate{Wood: wood.Spruce()})
	world.RegisterItem("minecraft:birch_fence_gate", WoodFenceGate{Wood: wood.Birch()})
	world.RegisterItem("minecraft:jungle_fence_gate", WoodFenceGate{Wood: wood.Jungle()})
	world.RegisterItem("minecraft:acacia_fence_gate", WoodFenceGate{Wood: wood.Acacia()})
	world.RegisterItem("minecraft:dark_oak_fence_gate", WoodFenceGate{Wood: wood.DarkOak()})
	world.RegisterItem("minecraft:crimson_fence_gate", WoodFenceGate{Wood: wood.Crimson()})
	world.RegisterItem("minecraft:warped_fence_gate", WoodFenceGate{Wood: wood.Warped()})
	world.RegisterItem("minecraft:trapdoor", WoodTrapdoor{Wood: wood.Oak()})
	world.RegisterItem("minecraft:spruce_trapdoor", WoodTrapdoor{Wood: wood.Spruce()})
	world.RegisterItem("minecraft:birch_trapdoor", WoodTrapdoor{Wood: wood.Birch()})
	world.RegisterItem("minecraft:jungle_trapdoor", WoodTrapdoor{Wood: wood.Jungle()})
	world.RegisterItem("minecraft:acacia_trapdoor", WoodTrapdoor{Wood: wood.Acacia()})
	world.RegisterItem("minecraft:dark_oak_trapdoor", WoodTrapdoor{Wood: wood.DarkOak()})
	world.RegisterItem("minecraft:crimson_trapdoor", WoodTrapdoor{Wood: wood.Crimson()})
	world.RegisterItem("minecraft:warped_trapdoor", WoodTrapdoor{Wood: wood.Warped()})
	world.RegisterItem("minecraft:wooden_door", WoodDoor{Wood: wood.Oak()})
	world.RegisterItem("minecraft:spruce_door", WoodDoor{Wood: wood.Spruce()})
	world.RegisterItem("minecraft:birch_door", WoodDoor{Wood: wood.Birch()})
	world.RegisterItem("minecraft:jungle_door", WoodDoor{Wood: wood.Jungle()})
	world.RegisterItem("minecraft:acacia_door", WoodDoor{Wood: wood.Acacia()})
	world.RegisterItem("minecraft:dark_oak_door", WoodDoor{Wood: wood.DarkOak()})
	world.RegisterItem("minecraft:crimson_door", WoodDoor{Wood: wood.Crimson()})
	world.RegisterItem("minecraft:warped_door", WoodDoor{Wood: wood.Warped()})
	for _, c := range allCoral() {
		world.RegisterItem("minecraft:coral", c.(world.Item))
	}
	for _, c := range allCoralBlocks() {
		world.RegisterItem("minecraft:coral_block", c.(world.Item))
	}
	world.RegisterItem("minecraft:pumpkin", Pumpkin{})
	world.RegisterItem("minecraft:lit_pumpkin", LitPumpkin{})
	world.RegisterItem("minecraft:carved_pumpkin", Pumpkin{Carved: true})
	world.RegisterItem("minecraft:end_stone", EndStone{})
	world.RegisterItem("minecraft:netherrack", Netherrack{})
	world.RegisterItem("minecraft:clay", Clay{})
	world.RegisterItem("minecraft:bone_block", BoneBlock{})
	world.RegisterItem("minecraft:lantern", Lantern{Type: fire.Normal()})
	world.RegisterItem("minecraft:soul_lantern", Lantern{Type: fire.Soul()})
	world.RegisterItem("minecraft:ancient_debris", AncientDebris{})
	world.RegisterItem("minecraft:emerald_ore", EmeraldOre{})
	world.RegisterItem("minecraft:diamond_ore", DiamondOre{})
	world.RegisterItem("minecraft:lapis_ore", LapisOre{})
	world.RegisterItem("minecraft:nether_gold_ore", NetherGoldOre{})
	world.RegisterItem("minecraft:gold_ore", GoldOre{})
	world.RegisterItem("minecraft:iron_ore", IronOre{})
	world.RegisterItem("minecraft:coal_ore", CoalOre{})
	world.RegisterItem("minecraft:quartz_ore", NetherQuartzOre{})
	world.RegisterItem("minecraft:dye", CocoaBean{})
	world.RegisterItem("minecraft:wheat_seeds", WheatSeeds{})
	world.RegisterItem("minecraft:beetroot_seeds", BeetrootSeeds{})
	world.RegisterItem("minecraft:potato", Potato{})
	world.RegisterItem("minecraft:carrot", Carrot{})
	world.RegisterItem("minecraft:pumpkin_seeds", PumpkinSeeds{})
	world.RegisterItem("minecraft:melon_seeds", MelonSeeds{})
	world.RegisterItem("minecraft:melon_block", Melon{})
	world.RegisterItem("minecraft:sand", Sand{})
	world.RegisterItem("minecraft:sand", Sand{Red: true})
	world.RegisterItem("minecraft:gravel", Gravel{})
	world.RegisterItem("minecraft:brick_block", Bricks{})
	world.RegisterItem("minecraft:soul_sand", SoulSand{})
	world.RegisterItem("minecraft:barrier", Barrier{})
	world.RegisterItem("minecraft:basalt", Basalt{})
	world.RegisterItem("minecraft:polished_basalt", Basalt{Polished: true})
	world.RegisterItem("minecraft:crying_obsidian", CryingObsidian{})
	world.RegisterItem("minecraft:seaLantern", SeaLantern{})
	world.RegisterItem("minecraft:soul_soil", SoulSoil{})
	world.RegisterItem("minecraft:blue_ice", BlueIce{})
	world.RegisterItem("minecraft:gilded_blackstone", GildedBlackstone{})
	world.RegisterItem("minecraft:shroomlight", Shroomlight{})
	world.RegisterItem("minecraft:torch", Torch{Type: fire.Normal()})
	world.RegisterItem("minecraft:soul_torch", Torch{Type: fire.Soul()})
	world.RegisterItem("minecraft:cake", Cake{})
	world.RegisterItem("minecraft:nether_wart", NetherWart{})
	world.RegisterItem("minecraft:invisibleBedrock", InvisibleBedrock{})
	world.RegisterItem("minecraft:noteblock", NoteBlock{})
	world.RegisterItem("minecraft:dragon_egg", DragonEgg{})
}

func init() {
	item_internal.Air = Air{}
	item_internal.Grass = Grass{}
	item_internal.GrassPath = GrassPath{}
	item_internal.Farmland = Farmland{Hydration: 0}
	item_internal.Dirt = Dirt{}
	item_internal.IsUnstrippedLog = func(b world.Block) bool {
		l, ok := b.(Log)
		return ok && !l.Stripped
	}
	item_internal.StripLog = func(b world.Block) world.Block {
		l := b.(Log)
		l.Stripped = true
		return l
	}
	item_internal.IsCarvedPumpkin = func(b world.Item) bool {
		p, ok := b.(Pumpkin)
		return ok && p.Carved
	}
	item_internal.IsUncarvedPumpkin = func(b world.Block) bool {
		p, ok := b.(Pumpkin)
		return ok && !p.Carved
	}
	item_internal.CarvePumpkin = func(b world.Block, face world.Face) world.Block {
		return Pumpkin{Carved: true, Facing: face.Direction()}
	}
	item_internal.Lava = Lava{Depth: 8, Still: true}
	item_internal.Water = Water{Depth: 8, Still: true}
	item_internal.IsWater = func(b world.Block) bool {
		_, ok := b.(Water)
		return ok
	}
	item_internal.IsWaterSource = func(b world.Block) bool {
		water, ok := b.(Water)
		return ok && water.Depth == 8
	}
	item_internal.BoneMeal = func(pos world.BlockPos, w *world.World) bool {
		b := w.Block(pos)
		if bonemealAffected, ok := b.(BoneMealAffected); ok {
			return bonemealAffected.BoneMeal(pos, w)
		}
		return false
	}
	item_internal.Replaceable = replaceableWith
	entity_internal.CanSolidify = func(b world.Block, pos world.BlockPos, w *world.World) bool {
		gravity, ok := b.(GravityAffected)
		if !ok {
			return false
		}
		return gravity.CanSolidify(pos, w)
	}
	item_internal.Fire = Fire{}
}

// readSlice reads an interface slice from a map at the key passed.
//noinspection GoCommentLeadingSpace
func readSlice(m map[string]interface{}, key string) []interface{} {
	//lint:ignore S1005 Double assignment is done explicitly to prevent panics.
	v, _ := m[key]
	b, _ := v.([]interface{})
	return b
}

// readString reads a string from a map at the key passed.
//noinspection GoCommentLeadingSpace
func readString(m map[string]interface{}, key string) string {
	//lint:ignore S1005 Double assignment is done explicitly to prevent panics.
	v, _ := m[key]
	b, _ := v.(string)
	return b
}

// readInt32 reads an int32 from a map at the key passed.
//noinspection GoCommentLeadingSpace
func readInt32(m map[string]interface{}, key string) int32 {
	//lint:ignore S1005 Double assignment is done explicitly to prevent panics.
	v, _ := m[key]
	b, _ := v.(int32)
	return b
}

// readByte reads a byte from a map at the key passed.
//noinspection GoCommentLeadingSpace
func readByte(m map[string]interface{}, key string) byte {
	//lint:ignore S1005 Double assignment is done explicitly to prevent panics.
	v, _ := m[key]
	b, _ := v.(byte)
	return b
}
