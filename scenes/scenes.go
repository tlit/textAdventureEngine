package scenes

import (
	"fmt"
	"strings"
	"textadventureengine/actors"
	. "textadventureengine/types"
	"textadventureengine/utils"
)

type Scene struct {
	Id          `json:"id"`
	Name        `json:"name"`
	Description `json:"description"`
	Actors      map[string]actors.Actor `json:"actors"`
	Exits       DestinationMap          `json:"exits"`
}
type Requirement struct {
	Id               `json:"id"`
	SceneActors      []actors.Actor `json:"scene_actors"`
	actors.Inventory `json:"inventory"`
	SuccessMessage   string `json:"successmessage"`
	FailMessage      string `json:"failmessage"`
}

type Exit struct {
	Id                    `json:"id"`
	Description           string      `json:"description"`
	Destination           string      `json:"destination"`
	ExitRequirement       Requirement `json:"exit_requirement"`
	VisibilityRequirement Requirement `json:"visibility_requirement"`
}

type SceneMap map[string]*Scene
type ExitMap map[string]Exit
type DestinationMap map[string]string

var Room = Scene{
	"1",
	"empty room",
	"You find yourself in an empty, windowless room.",
	nil,
	DestinationMap{},
}

var Exits = ExitMap{
	"PitExit": Exit{
		"1",
		"Light streams in from above the rim of the pit.",
		"AbovePit",
		Requirement{"grappleUp", []actors.Actor{}, actors.Inventory{"grappling hook": actors.Grapple}, "You swing the grappling hook over the rim of the pit and climb the rope.", "You cannot climb up to the rim of the pit."},
		Requirement{},
	},
	"PitEntrance": Exit{
		"2",
		"A dark pit.",
		"Pit",
		Requirement{},
		Requirement{},
	},
}
var Scenes = SceneMap{
	"AbovePit": &Scene{
		"3",
		"above the pit",
		"You are in the open desert. The sun beats down on your skin. The air is dry",
		map[string]actors.Actor{},
		DestinationMap{"down": "PitEntrance"},
	},
	"Pit": &Scene{
		"2",
		"empty pit",
		"You find yourself at the bottom of a pit. The air is cool and humid. You can see no way to climb out.",
		map[string]actors.Actor{"grappling hook": actors.Grapple},
		DestinationMap{"up": "PitExit"},
	},
}

func (s Scene) GetDestination(dest string) string {
	return s.Exits[dest]
}

func (s Scene) Run() {
	s.print()
	fmt.Println(s.Exits)
}

func (s Scene) GetNextScene(dir string) string {
	var dest string
	switch dir {
	case "up", "down", "north", "south", "east", "west":
		dest = s.GetDestination(dir)
		break
	default:
		dest = ""
	}
	return dest
}

func (mp SceneMap) Get(s string) *Scene {
	return mp[s]
}

func (mp ExitMap) Get(s string) Exit {
	return mp[s]
}

func (s Scene) print() {
	utils.PrintLine(string(s.Description))
	var actorNames []string
	for _, x := range s.Actors {
		if utils.StartsWithVowel(string(x.Description)) {
			actorNames = append(actorNames, "an "+string(x.Name))
		} else {
			actorNames = append(actorNames, "a "+string(x.Name))
		}
	}
	if len(s.Actors) > 0 {
		utils.PrintLine("You see here; " + strings.Join(actorNames, ", "))
	}
}
