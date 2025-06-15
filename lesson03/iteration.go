package iteration

import (
	"fmt"
	"strings"
)

const counter = 5

func Repeat(character string) string {
	var repeated strings.Builder
	for i := 0; i < counter; i++ {
		repeated.WriteString(character)
	}
	return repeated.String()
}

func CompareStrings(character01, character02 string) int {
	return strings.Compare(character01, character02)
}

func CutPrefixString(c, character string) bool  {
	after, found := strings.CutPrefix(c, character)
	fmt.Printf("CutPrefix(%q, %q) = %q, %v\n", c, character, after, found)
	return found
}