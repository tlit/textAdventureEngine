package scenes

import (
	"textadventureengine/actors"
	. "textadventureengine/types"
)

var (
	DemoPit_DemoPitDarkRoom Exit
	DemoPit_DemoPitTop      Exit

	DemoPit = Scene{
		Name:        "DemoPit",
		Description: "You are in a pit. The air is dry and hot.",
		Actors: Actors{
			string(actors.Candle.Name): actors.Candle,
		},
		Exits: Exits{
			North: &DemoPit_DemoPitDarkRoom,
			Up:    &DemoPit_DemoPitTop,
		},
	}
	DemoPitDarkroom = Scene{
		Name:        "DemoPitDarkroom",
		Description: "It will be dark here once implemented",
		Actors: Actors{
			string(actors.GrapplingHook.Name): actors.GrapplingHook,
		},
		Exits: Exits{
			South: &DemoPit_DemoPitDarkRoom,
		},
	}
	DemoPitTop = Scene{
		Name:        "DemoPitTop",
		Description: "At the top of the pit",
		Actors:      Actors{},
		Exits: Exits{
			Down: &DemoPit_DemoPitTop,
		},
	}

	Scenario_DemoPit = Scenario{
		FirstScene: DemoPit,
		Scenes: map[string]Scene{
			string(DemoPit.Name):         DemoPit,
			string(DemoPitDarkroom.Name): DemoPitDarkroom,
			string(DemoPitTop.Name):      DemoPitTop,
		},
	}
)

func init() {
	DemoPit_DemoPitDarkRoom = Exit{
		Description: "dark hole",
		Destinations: Destinations{
			string(DemoPit.Name):         {"a dark hole, just large enough to squeeze through", DemoPitDarkroom, Flags{}},
			string(DemoPitDarkroom.Name): {"a hole in the wall through which a pit is visible", DemoPit, Flags{}},
		},
		Visible: true,
	}
	DemoPit_DemoPitTop = Exit{
		Description: "pit top",
		Destinations: Destinations{
			string(DemoPitTop.Name): {"a pit in the ground", DemoPit, Flags{}},
			string(DemoPit.Name):    {"light and dust both stream down from the opening above you", DemoPitTop, Flags{"climb": 20}},
		},
		Visible: true,
	}
}
