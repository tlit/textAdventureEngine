package types

import (
	"fmt"
)

type Id string
type Name string
type Description string
type Direction int
type Flags map[string]interface{}
type Flag struct {
	Key   string
	Value interface{}
}

const (
	Up Direction = iota
	Down
	North
	South
	East
	West
)

var DirectionMap = map[string]Direction{
	"up":    Up,
	"down":  Down,
	"north": North,
	"south": South,
	"east":  East,
	"west":  West,
}

func (f Flags) Contains(s string) bool {
	return f[s] == nil
}

type Met map[bool]string

func (d Description) Print() {
	fmt.Println(d)
}
