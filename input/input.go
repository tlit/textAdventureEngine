package input

import (
	"bufio"
	"os"
	"strings"
	"textadventureengine/gameStructure"
	"textadventureengine/utils"
)

func NewScanner() *bufio.Scanner {
	return bufio.NewScanner(os.Stdin)
}

func scanText(s bufio.Scanner) string {
	if s.Scan() {
		return s.Text()
	}
	return ""
}

func ProcessInput(gs *gameStructure.GameStructure) {
	var input string
	input = scanText(gs.Input)
	input = strings.ToLower(input)
	words := strings.Split(input, " ")
	verb := words[0]
	switch verb {
	case "go":
		var direction string
		if len(words) > 1 {
			direction = words[1]
		} else {
			utils.Prt("Which way?")
			direction = scanText(gs.Input)
		}
		if _, ok := gs.CurrentScene.Exits[direction]; ok {
			gs.GoDirection(direction)
		} else {
			utils.Prt("You cannot go " + direction)
		}
		return
	case "look":
		var direction string
		if len(words) > 1 {
			direction = words[1]
		} else {
			utils.Prt("Which way?")
			direction = scanText(gs.Input)
		}
		if exitId, ok := gs.CurrentScene.Exits[direction]; ok {
			exit := gs.Exits[exitId]
			utils.Prt(exit.Description)
		} else {
			utils.Prt("You cannot see anything in that direction.")
		}
		return
	case "get", "grab", "take":
		var object string
		if len(words) > 1 {
			object = strings.Join(words[1:len(words)], " ")
		} else {
			utils.Prt(verb + " what?")
			object = scanText(gs.Input)
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
		return
	case "drop":
		var object string
		if len(words) > 1 {
			object = strings.Join(words[1:len(words)], " ")
		} else {
			utils.Prt(verb + " what?")
			object = scanText(gs.Input)
		}
		if o, ok := gs.Player.Inventory[object]; ok {
			gs.DropObject(object)
			utils.Prt("You drop the " + string(o.Name))
		} else {
			utils.Prt("I don't understand " + object)
		}
		return
	case "i", "inv", "inventory":
		utils.Prt("You are carrying:")
		if len(gs.Player.Inventory) == 0 {
			utils.Prt("	Nothing")
		}
		for _, o := range gs.Player.Inventory {
			utils.Prt("	" + string(o.Name))
		}
		return
	default:
		utils.Prt("I don't know \"" + verb + "\"")
	}
}
