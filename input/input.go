package input

import (
	"bufio"
	"os"
	"strings"
	"textadventureengine/gameStructure"
	"textadventureengine/utils"
	v "textadventureengine/verbs"
)

func NewScanner() *bufio.Scanner {
	return bufio.NewScanner(os.Stdin)
}

func ProcessInput(gs *gameStructure.GameStructure) {
	var input string
	if gs.Input.Scan() {
		input = func() string {
			return gs.Input.Text()
		}()
	}

	input = strings.ToLower(input)
	words := strings.Split(input, " ")
	if f, ok := v.Verbs[words[0]]; ok {
		f(gs, words...)
	} else {
		utils.Prt("I don't know \"" + words[0] + "\"")
	}
}
