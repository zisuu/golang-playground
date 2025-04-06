package main

import (
	"fmt"
	"strings"
)

// Repeat takes a string and an integer and returns the string repeated that many times.
func Repeat(character string, iterations int) string {
	var repeated strings.Builder
	for i := 0; i < iterations; i++ {
		repeated.WriteString(character)
	}
	return repeated.String()
}

func main() {
	fmt.Println(Repeat("a", 5))
}
