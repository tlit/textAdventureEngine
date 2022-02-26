package player

import "textadventureengine/actors"

type Player struct {
	actors.Inventory `json:"inventory"`
}
