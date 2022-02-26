package actors

import (
	"encoding/json"
	"io/ioutil"
	. "textadventureengine/types"
)

type Actor struct {
	Id          `json:"id"`
	Name        `json:"name"`
	Description `json:"description"`
	Flags       `json:"flags"`
}

type Inventory map[string]Actor

var Grapple = Actor{"grapple", "grappling hook", "a large metal hook on the end of a rope", nil}

func ReadActors(s string) map[string]*Actor {
	data := map[string]*Actor{}
	file, _ := ioutil.ReadFile("json/scenario/" + s + "/actors.json")
	_ = json.Unmarshal([]byte(file), &data)
	return data
}
