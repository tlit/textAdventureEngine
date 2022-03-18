package gameStructure

import (
	"bufio"
	"reflect"
	"textadventureengine/actors"
	"textadventureengine/player"
	"textadventureengine/scenes"
	"textadventureengine/types"
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
	//req := gs.Requirements[exit.Requirements]
	//if req == nil {
	//	req = &scenes.Requirement{}
	//}
	//if gs.Meets(*req) {
	//	utils.Prt(req.SuccessMessage + "\n")
	//} else {
	//	utils.Prt(req.FailMessage + "\n")
	//}
	canExit := true
	for k, v := range exit.Requirements {
		if !gs.Player.Can(types.Flag{k, v}) {
			canExit = false
		}
	}
	if !canExit {
		utils.Prt("No can do.")
	} else {
		gs.NextScene = gs.Scenes[exit.Destination]
	}
}

func (gs *GameStructure) GetObject(o string) *actors.Actor {
	if obj, ok := gs.CurrentScene.Actors[o]; ok {
		return obj
	}
	return nil
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

func (gs *GameStructure) UseObject(o string) {

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

func (gs GameStructure) PrintVisibleExits() (output string) {
	var out string
	for k, v := range gs.CurrentScene.Exits {
		x := gs.Exits[v]
		if x.Visible {
			out = out + "\t" + k + ": " + x.Description
		}
	}
	return out
}
