package gameStructure

import (
	"bufio"
	"fmt"
	"reflect"
	"textadventureengine/player"
	"textadventureengine/scenes"
	"textadventureengine/utils"
)

type GameStructure struct {
	player.Player
	CurrentScene *scenes.Scene
	NextScene    *scenes.Scene
	Input        bufio.Scanner
	Scenes       *scenes.SceneMap
	Exits        *scenes.ExitMap
}

func (gs *GameStructure) GoDirection(d string) {
	scn := gs.CurrentScene.GetNextScene(d)
	exit := gs.Exits.Get(scn)
	if gs.Meets(exit.ExitRequirement) {
		x := gs.Scenes.Get(exit.Destination)
		gs.NextScene = x
		utils.PrintLine(exit.ExitRequirement.SuccessMessage)
	} else {
		utils.PrintLine(exit.ExitRequirement.FailMessage)
	}
	fmt.Println()
}

func (gs *GameStructure) TakeObject(o string) {
	//TODO check requirements are met
	obj := gs.CurrentScene.Actors[o]
	gs.Player.Inventory[o] = obj
	delete(gs.CurrentScene.Actors, o)
}

func (gs *GameStructure) DropObject(o string) {
	//TODO check requirements are met
	obj := gs.Player.Inventory[o]
	gs.CurrentScene.Actors[o] = obj
	delete(gs.Inventory, o)
}

func (gs GameStructure) Meets(req scenes.Requirement) bool {

	sceneHasRequiredActors := len(req.SceneActors) == 0
	if !sceneHasRequiredActors {
		for _, ra := range req.SceneActors {
			for _, sa := range gs.CurrentScene.Actors {
				if reflect.DeepEqual(ra, sa) {
					sceneHasRequiredActors = true
				}
			}
		}
	}
	inventoryHasRequiredActors := len(req.Inventory) == 0
	if !inventoryHasRequiredActors {
		for _, ri := range req.Inventory {
			for _, pi := range gs.Player.Inventory {
				if reflect.DeepEqual(ri, pi) {
					inventoryHasRequiredActors = true
				}
			}
		}
	}
	return sceneHasRequiredActors && inventoryHasRequiredActors
}
