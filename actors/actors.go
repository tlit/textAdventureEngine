package actors

import (
	"encoding/json"
	"io/ioutil"
	. "textadventureengine/types"
)

func ReadActors(s string) map[string]*Actor {
	data := map[string]*Actor{}
	file, _ := ioutil.ReadFile("json/scenario/" + s + "/actors.json")
	_ = json.Unmarshal([]byte(file), &data)
	return data
}

var (
	GrapplingHook = Actor{
		"grapplingHook",
		"grappling hook",
		"Three splayed iron hooks attacked to the end of a rope",
		Flags{
			"climb": 20,
		},
	}
	Candle = Actor{
		"candle",
		"candle",
		"A wax candle",
		Flags{
			"illuminate": true,
		},
	}
)
