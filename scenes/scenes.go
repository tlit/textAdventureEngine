package scenes

import (
	"encoding/json"
	"io/ioutil"
	"strings"
	"textadventureengine/actors"
	. "textadventureengine/types"
)

type DestinationMap map[string]string

type Scene struct {
	Id          `json:"id"`
	Name        `json:"name"`
	Description `json:"description"`
	Actors      map[string]*actors.Actor `json:"actors"`
	Exits       map[string]string        `json:"exits"`
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
	Description           string `json:"description"`
	Destination           string `json:"destination"`
	ExitRequirement       string `json:"exit_requirement"`
	VisibilityRequirement string `json:"visibility_requirement"`
}

type Actors map[string]string

var Room = Scene{
	"1",
	"empty room",
	"You find yourself in an empty, windowless room.",
	nil,
	nil,
}

//"grappleUp", []actors.Actor{}, actors.Inventory{"grappling hook": actors.Grapple}, "You swing the grappling hook over the rim of the pit and climb the rope.", "You cannot climb up to the rim of the pit."
//var Exits = json.Marshal()
//var Scenes =

func (s Scene) GetDestination(dest string) string {
	return s.Exits[dest]
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

func (s Scene) Print() string {
	return string(s.Description) + "\n"
}

func ReadScenes(s string) map[string]*Scene {
	data := map[string]*Scene{}
	file, _ := ioutil.ReadFile("json/scenario/" + strings.ToLower(s) + "/scenes.json")
	_ = json.Unmarshal([]byte(file), &data)
	return data
}

func ReadExits(s string) map[string]*Exit {
	data := map[string]*Exit{}
	file, _ := ioutil.ReadFile("json/scenario/" + strings.ToLower(s) + "/exits.json")
	_ = json.Unmarshal([]byte(file), &data)
	return data
}

func ReadRequirements(s string) map[string]*Requirement {
	data := map[string]*Requirement{}
	file, _ := ioutil.ReadFile("json/scenario/" + strings.ToLower(s) + "/requirements.json")
	_ = json.Unmarshal([]byte(file), &data)
	return data
}
