package utils

import (
	"fmt"
	"strings"
)

func PrintLine(s string) {
	spl := strings.Split(s, " ")
	first := strings.Title(spl[0])
	out := []string{first}
	rest := spl[1:len(spl)]
	for _, word := range rest {
		out = append(out, word)
	}
	fmt.Println(strings.Join(out, " "))
}

func StartsWithVowel(s string) bool {
	return strings.HasPrefix(s, "a") || strings.HasPrefix(s, "e") || strings.HasPrefix(s, "i") || strings.HasPrefix(s, "o") || strings.HasPrefix(s, "u")
}
