package verbs

import (
	"strings"
	"textadventureengine/gameStructure"
	"textadventureengine/types"
	"textadventureengine/utils"
)

type Verb map[string]func(*gameStructure.GameStructure, ...string)

var Verbs = Verb{
	"carry":     na,
	"climb":     climb,
	"drop":      drop,
	"get":       get,
	"go":        travel,
	"grab":      get,
	"inv":       inv,
	"inventory": inv,
	"look":      look,
	"take":      get,
}

func travel(gs *gameStructure.GameStructure, words ...string) {
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
	if _, okDir := types.DirectionMap[direction]; okDir {
		gs.GoDirection(direction)
	} else {
		utils.Prt("You cannot go " + direction)
	}
}

func look(gs *gameStructure.GameStructure, words ...string) {
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
	if dir, okDir := types.DirectionMap[direction]; okDir {
		if exitId, ok := gs.CurrentScene.Exits[dir]; ok {
			exit := gs.Exits[exitId]
			utils.Prt(exit.Description)
		} else {
			utils.Prt("You cannot see anything in that direction.")
		}
	}
}
func get(gs *gameStructure.GameStructure, words ...string) {
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
func drop(gs *gameStructure.GameStructure, words ...string) {
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
func inv(gs *gameStructure.GameStructure, words ...string) {
	utils.Prt("You are carrying:")
	if len(gs.Player.Inventory) == 0 {
		utils.Prt("	Nothing")
	}
	for _, o := range gs.Player.Inventory {
		utils.Prt("	" + string(o.Name))
	}
}

func climb(gs *gameStructure.GameStructure, words ...string) {
	if x, y := gs.Player.Flags["climb"]; y {
		var object string
		if x != nil {
			if len(words) > 1 {
				object = strings.Join(words[1:], " ")
			} else {
				utils.Prt(words[0] + " what or where?")
				if gs.Input.Scan() {
					object = func() string {
						return gs.Input.Text()
					}()
				}
			}
			//climb in direction
			if _, okDir := types.DirectionMap[object]; okDir {
				if _, okFlag := gs.Player.Flags["climb"]; okFlag {
					travel(gs, words[1:]...)
				}
			}
		}
	} else {
		utils.Prt("you cannot climb")
	}
}

func na(gs *gameStructure.GameStructure, words ...string) {
	utils.Prt("you don't know how to " + words[0])
}
