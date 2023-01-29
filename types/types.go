package types

import (
	"fmt"
	"strings"
	"textadventureengine/utils"
)

type Id string
type Name string
type Description string
type Flags map[string]interface{}
type Flag struct {
	Key   string
	Value interface{}
}
type Operator int
type CompassDirection struct {
	int
	Name string
}

var (
	North = CompassDirection{0, "North"}
	East  = CompassDirection{1, "East"}
	South = CompassDirection{2, "South"}
	West  = CompassDirection{3, "West"}
	Up    = CompassDirection{4, "Up"}
	Down  = CompassDirection{5, "Down"}
)

var DirectionMap = map[string]CompassDirection{
	"north": North,
	"east":  East,
	"south": South,
	"west":  West,
	"up":    Up,
	"down":  Down,
}

const (
	Equals Operator = iota
	GreaterThan
	LessThan
)

var Operators = map[string]Operator{
	"=": Equals,
	">": GreaterThan,
	"<": LessThan,
}

func (f Flags) Contains(s string) bool {
	return f[s] == nil
}

type Met map[bool]string

type Destination struct {
	Description
	Scene
	Requirements Flags `json:"requirements"`
}

type Destinations map[string]*Destination
type Exit struct {
	Id                    `json:"id"`
	Description           `json:"description"`
	Destinations          `json:"destinations"`
	VisibilityRequirement string `json:"visibility_requirement"`
	Visible               bool
}
type Exits map[CompassDirection]*Exit

type Scenario struct {
	FirstScene Scene
	Scenes     map[string]Scene
}
type Scene struct {
	Id          `json:"id"`
	Name        `json:"name"`
	Description `json:"description"`
	Actors      `json:"actors"`
	Exits       `json:"exits"`
}
type Requirement struct {
	Id             `json:"id"`
	SceneActors    Actors `json:"scene_actors"`
	Inventory      `json:"inventory"`
	SuccessMessage string `json:"successmessage"`
	FailMessage    string `json:"failmessage"`
}

type DestinationMap map[string]string

type Actor struct {
	Id          `json:"id"`
	Name        `json:"name"`
	Description `json:"description"`
	Flags       `default: Flags{}`
}

type Actors map[string]Actor

type Inventory Actors

func Print(s string) {
	fmt.Println(s)
}

func (s *Scene) Run() {
	utils.Prt(s.PrintDescription())
	utils.Prt("You see here:")
	utils.Prt(s.PrintActors())
	return
}

func (s Scene) PrintActors() string {
	var out string
	var x []string
	for _, v := range s.Actors {
		name := string(v.Name)
		if utils.StartsWithVowel(name) {
			x = append(x, "\tan "+name)
		} else {
			x = append(x, "\ta "+name)
		}
	}
	if len(x) > 0 {
		out = strings.Join(x, ",\n")
	} else {
		out = ""
	}
	return out
}

func (s Scene) GetDestinationName(dir CompassDirection) string {
	return string(s.Exits[dir].Id)
}

func (s Scene) PrintDescription() string {
	return string(s.Description) + "\n"
}

func (d Description) Print() string {
	return string(d)
}

func (d CompassDirection) Print() string {
	direction := d
	return direction.Name
}
