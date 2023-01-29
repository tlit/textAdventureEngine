package player

import (
	"reflect"
	. "textadventureengine/types"
)

type Player struct {
	Inventory Actors                 `json:"inventory"`
	Flags     map[string]interface{} `json:"flags"`
}

func (p *Player) Can(f Flag) bool {
	switch reflect.TypeOf(f.Value) {
	case reflect.TypeOf(float64(0)):
		want, ok := f.Value.(float64)
		if ok {
			got, ok := p.Flags[f.Key].(float64)
			if ok {
				if got >= want {
					return true
				}
				println("you cannot climb that high")
				return false
			}
		}
	default:
	}
	return false
}
