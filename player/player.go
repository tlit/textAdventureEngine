package player

import (
	. "textadventureengine/types"
)

type Player struct {
	Inventory Actors                 `json:"inventory"`
	Flags     map[string]interface{} `json:"flags"`
}

func (p *Player) Can(f Flag) bool {
	want := convertToFloat(f.Value)
	got := convertToFloat(f.Value)
	if got >= want {
		return true
	}
	{
		println("you cannot climb that high")
		return false
	}
	return true
}

func convertToFloat(v interface{}) float64 {
	var ok bool
	switch v.(type) {
	case int:
		if _, ok = v.(int); ok {
			return float64(v.(int))
		}
	case float64:
		return v.(float64)
	default:
	}
	return v.(float64)
}
