package player

import (
	"reflect"
	"textadventureengine/actors"
	"textadventureengine/types"
)

type Player struct {
	actors.Inventory `json:"inventory"`
	Flags            map[string]interface{} `json:"flags"`
}

func (p *Player) Can(f types.Flag) bool {
	switch reflect.TypeOf(f.Value) {
	case reflect.TypeOf(float64(0)):
		want, ok := f.Value.(float64)
		if ok {
			got, ok := p.Flags[f.Key].(float64)
			if ok {
				return got >= want
			}
		}
	default:
	}
	return false
}
