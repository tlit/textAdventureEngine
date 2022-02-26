package gameStructure

import (
	"bufio"
	"fmt"
	"reflect"
	"strings"
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

func (gs *GameStructure) Run() {
	utils.PrintLine(gs.CurrentScene.Print())
	var out string
	var x []string
	for _, v := range gs.CurrentScene.Actors {
		name := string(v.Name)
		if utils.StartsWithVowel(name) {
			x = append(x, "an "+name)
		} else {
			x = append(x, "a "+name)
		}
	}
	if len(x) > 0 {
		out = strings.Join(x, ", ")
	} else {
		out = ""
	}
	utils.PrintLine(out)
	fmt.Println(gs.CurrentScene.Exits)
	return
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
