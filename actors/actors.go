package actors

import . "textadventureengine/types"

type Actor struct {
	Id          `json:"id"`
	Name        `json:"name"`
	Description `json:"description"`
	Flags       `json:"flags"`
}

type Inventory map[string]Actor

var Grapple = Actor{"grapple", "grappling hook", "a large metal hook on the end of a rope", nil}
