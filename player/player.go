package player

import (
	"textadventureengine/actors"
	"textadventureengine/types"
)

type Player struct {
	actors.Inventory `json:"inventory"`
	types.Flags       `json:"flags"`
}

func(p *Player) can(property string) bool {
	switch property {
	case "climb":
		if p.Flags.Contains("climb"){
			return true
		}
	default:
		return false
	}
	return false
}