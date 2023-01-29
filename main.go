package main

import (
	"reflect"
	"textadventureengine/gameStructure"
	"textadventureengine/input"
	"textadventureengine/player"
	"textadventureengine/scenes"
	. "textadventureengine/types"
	"textadventureengine/utils"
)

func main() {
	var scenario = scenes.Scenario_DemoPit
	//scn := scenes.ReadScenes(scenario)
	//firstScene := scn[scenario]
	//ext := scenes.ReadExits(scenario)
	//act := actors.ReadActors(scenario)
	//req := scenes.ReadRequirements(scenario)
	Game := gameStructure.GameStructure{
		player.Player{Actors{}, Flags{}},
		scenario,
		&scenario.FirstScene,
		&Scene{},
		*input.NewScanner(),
		scenario.FirstScene.Actors,
	}

	//Main loop
	for true {
		if !reflect.DeepEqual(*Game.NextScene, Scene{}) {
			Game.CurrentScene = Game.NextScene
			Game.NextScene = &Scene{}
		}
		//for k, _ := range Game.CurrentScene.Actors {
		//	Game.CurrentScene.Actors[k] = Game.Actors[k]
		//}
		Game.Player.Flags = Flags{}
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
