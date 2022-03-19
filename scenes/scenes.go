package scenes

import (
	"encoding/json"
	"io/ioutil"
	"strings"
	"textadventureengine/actors"
	. "textadventureengine/types"
	"textadventureengine/utils"
)

type DestinationMap map[string]string

type Scene struct {
	Id          `json:"id"`
	Name        `json:"name"`
	Description `json:"description"`
	Actors      map[string]*actors.Actor `json:"actors"`
	Exits       map[Direction]string     `json:"exits"`
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
	Requirements          Flags  `json:"requirements"`
	VisibilityRequirement string `json:"visibility_requirement"`
	Visible               bool
}

type Actors map[string]string

func (s *Scene) Run() {
	utils.Prt(s.PrintDescription())
	utils.Prt("You see here:")
	utils.Prt(s.PrintActors())
	return
}

func (s Scene) GetDestination(dir Direction) string {
	return s.Exits[dir]
}

func (s Scene) GetNextScene(dir string) string {
	var dest string
	if v, ok := DirectionMap[dir]; ok {
		dest = s.GetDestination(v)
	}
	return dest
}

func (s Scene) PrintDescription() string {
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

func (s Scene) PrintActors() string {
	var out string
	var x []string
	for _, v := range s.Actors {
		name := string(v.Name)
		if utils.StartsWithVowel(name) {
			x = append(x, "an "+name)
		} else {
			x = append(x, "a "+name)
		}
	}
	if len(x) > 0 {
		out = strings.Join(x, ",\n")
		out = "\t" + out
	} else {
		out = ""
	}
	return out
}
