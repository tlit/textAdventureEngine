package gameStructure

import (
	"bufio"
	"reflect"
	"textadventureengine/actors"
	"textadventureengine/player"
	"textadventureengine/scenes"
	"textadventureengine/utils"
)

type GameStructure struct {
	player.Player
	CurrentScene *scenes.Scene
	NextScene    *scenes.Scene
	Input        bufio.Scanner

	Scenes       map[string]*scenes.Scene
	Exits        map[string]*scenes.Exit
	Requirements map[string]*scenes.Requirement
	Actors       map[string]*actors.Actor
}

func (gs *GameStructure) GoDirection(d string) {
	scn := gs.CurrentScene.GetNextScene(d)

	exit := gs.Exits[scn]
	req := gs.Requirements[exit.ExitRequirement]
	if req == nil {
		req = &scenes.Requirement{}
	}
	if gs.Meets(*req) {
		utils.PrintLine(req.SuccessMessage + "\n")

	} else {
		utils.PrintLine(req.FailMessage + "\n")
	}
	gs.NextScene = gs.Scenes[exit.Destination]
}

func (gs *GameStructure) TakeObject(o string) *actors.Actor {
	//TODO check requirements are met
	obj := gs.CurrentScene.Actors[o]
	gs.Player.Inventory[o] = *obj
	delete(gs.CurrentScene.Actors, o)
	return obj
}

func (gs *GameStructure) DropObject(o string) {
	//TODO check requirements are met
	obj := gs.Player.Inventory[o]
	gs.CurrentScene.Actors[o] = &obj
	delete(gs.Player.Inventory, o)
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
		for x, _ := range req.Inventory {
			obj := gs.Actors[x]
			for _, pi := range gs.Player.Inventory {
				if reflect.DeepEqual(obj, &pi) {
					inventoryHasRequiredActors = true
				}
			}
		}
	}
	return sceneHasRequiredActors && inventoryHasRequiredActors
}
