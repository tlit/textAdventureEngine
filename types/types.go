package types

import (
	"fmt"
)

type Id string
type Name string
type Description string
type Direction string
type Flags map[string]interface{}

func (f Flags) Contains(s string) bool {
	return f[s] == nil
}

type Met map[bool]string

func (d Description) Print() {
	fmt.Println(d)
}
