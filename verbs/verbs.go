package verbs

import (
	"strings"
	"textadventureengine/gameStructure"
	"textadventureengine/utils"
)

type Verb map[string]func(*gameStructure.GameStructure, ...string)

var Verbs = Verb{
	"go":_go,
	"look":_look,
	"get":_get,
	"grab":_get,
	"take":_get,
	"drop":_drop,
	"inv":_inv,
	"inventory":_inv,
	"climb":_na,
	"carry":_na,
}

func _go(gs *gameStructure.GameStructure, words... string) {
	var direction string
	if len(words) > 1 {
		direction = words[1]
	} else {
		utils.Prt("Which way?")
		if gs.Input.Scan() {
			direction = func() string {
				return gs.Input.Text()
			}()
		}
	}
	if _, ok := gs.CurrentScene.Exits[direction]; ok {
		gs.GoDirection(direction)
	} else {
		utils.Prt("You cannot go " + direction)
	}
}

func _look(gs *gameStructure.GameStructure, words... string) {
	var direction string
	if len(words) > 1 {
		direction = words[1]
	} else {
		utils.Prt("Which way?")
		if gs.Input.Scan() {
			direction = func() string {
				return gs.Input.Text()
			}()
		}
	}
	if exitId, ok := gs.CurrentScene.Exits[direction]; ok {
		exit := gs.Exits[exitId]
		utils.Prt(exit.Description)
	} else {
		utils.Prt("You cannot see anything in that direction.")
	}
}
func _get(gs *gameStructure.GameStructure, words... string) {
	var object string
	if len(words) > 1 {
		object = strings.Join(words[1:len(words)], " ")
	} else {
		utils.Prt(words[0] + " what?")
		if gs.Input.Scan() {
			object = func() string {
				return gs.Input.Text()
			}()
		}
	}
	if obj, ok := gs.CurrentScene.Actors[object]; ok {
		if _, ok := obj.Flags["portable"]; ok {
			gs.TakeObject(object)
			utils.Prt("You take the " + string(obj.Name))
		} else {
			utils.Prt("You cannot take the " + string(obj.Name))
		}
	} else {
		utils.Prt("I don't understand " + object)
	}
}
func _drop(gs *gameStructure.GameStructure, words... string) {
	var object string
	if len(words) > 1 {
		object = strings.Join(words[1:len(words)], " ")
	} else {
		utils.Prt(words[0] + " what?")
		if gs.Input.Scan() {
			object = func() string {
				return gs.Input.Text()
			}()
		}
	}
	if o, ok := gs.Player.Inventory[object]; ok {
		gs.DropObject(object)
		utils.Prt("You drop the " + string(o.Name))
	} else {
		utils.Prt("I don't understand " + object)
	}
}
func _inv(gs *gameStructure.GameStructure, words... string) {
	utils.Prt("You are carrying:")
	if len(gs.Player.Inventory) == 0 {
		utils.Prt("	Nothing")
	}
	for _, o := range gs.Player.Inventory {
		utils.Prt("	" + string(o.Name))
	}
}

func _na(gs *gameStructure.GameStructure, words... string) {
	utils.Prt("not yet implemented")
}
