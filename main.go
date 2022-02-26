package main

import (
	"reflect"
	"textadventureengine/actors"
	"textadventureengine/gameStructure"
	"textadventureengine/input"
	"textadventureengine/player"
	"textadventureengine/scenes"
)

var Game gameStructure.GameStructure

func main() {
	Game = gameStructure.GameStructure{
		player.Player{actors.Inventory{}},
		*scenes.Scenes["Pit"],
		scenes.Scene{},
		input.NewScanner(),
	}

	//Main loop
	for true {
		Game.CurrentScene.Run()
		input.ProcessInput(&Game)
		if !reflect.DeepEqual(Game.NextScene, scenes.Scene{}) {
			Game.CurrentScene = Game.NextScene
		}
	}
}
