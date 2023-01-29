package gameStructure

import (
	"bufio"
	"textadventureengine/player"
	. "textadventureengine/types"
	"textadventureengine/utils"
)

type GameStructure struct {
	player.Player
	Scenario
	CurrentScene *Scene
	NextScene    *Scene
	Input        bufio.Scanner
	Actors
}

func (gs *GameStructure) GoDirection(d CompassDirection) {
	exit := gs.CurrentScene.Exits[d]
	dest := exit.Destinations[string(gs.CurrentScene.Name)]

	canExit := true

	for k, v := range dest.Requirements {
		if !gs.Player.Can(Flag{k, v}) {
			canExit = false
		}
	}

	if !canExit {
		utils.Prt("No can do.")
	} else {
		gs.NextScene = &dest.Scene
	}
}

func (gs *GameStructure) TakeObject(x Actor) bool {
	//TODO check requirements are met
	gs.Player.Inventory[string(x.Name)] = x
	delete(gs.Actors, string(x.Name))
	return true
}

func (gs *GameStructure) DropObject(x Actor) bool {
	//TODO check requirements are met
	obj := gs.Player.Inventory[string(x.Name)]
	gs.CurrentScene.Actors[string(x.Name)] = obj
	delete(gs.Player.Inventory, string(x.Name))
	return true
}

func (gs *GameStructure) UseObject(o string) {

}

func (gs GameStructure) PrintVisibleExits() (output string) {
	sceneName := string(gs.CurrentScene.Name)
	out := ""
	for k, v := range gs.CurrentScene.Exits {
		if v.Visible {
			out = out + "\t" + k.Print() + "\t"
			out = out + "\t" + v.Destinations[sceneName].Description.Print() + "\n"
		}
	}
	return out
}
