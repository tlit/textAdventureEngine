package main

import (
	"reflect"
	"textadventureengine/actors"
	"textadventureengine/gameStructure"
	"textadventureengine/input"
	"textadventureengine/player"
	"textadventureengine/scenes"
	"textadventureengine/types"
	"textadventureengine/utils"
)

func main() {
	scenario := "Pit"
	scn := scenes.ReadScenes(scenario)
	firstScene := scn[scenario]
	ext := scenes.ReadExits(scenario)
	act := actors.ReadActors(scenario)
	req := scenes.ReadRequirements(scenario)
	Game := gameStructure.GameStructure{
		player.Player{actors.Inventory{}, types.Flags{}},
		firstScene,
		&scenes.Scene{},
		*input.NewScanner(),
		scn,
		ext,
		req,
		act,
	}

	//Main loop
	for true {
		if !reflect.DeepEqual(*Game.NextScene, scenes.Scene{}) {
			Game.CurrentScene = Game.NextScene
			Game.NextScene = &scenes.Scene{}
		}
		for k, _ := range Game.CurrentScene.Actors {
			Game.CurrentScene.Actors[k] = Game.Actors[k]
		}
		Game.Player.Flags = types.Flags{}
		for _, v := range Game.Player.Inventory {
			for flag, val := range v.Flags {
				Game.Player.Flags[flag] = val
			}
		}
		Game.CurrentScene.Run()
		utils.Prt("\nExits:\n" + Game.PrintVisibleExits())
		input.ProcessInput(&Game)
	}
}
