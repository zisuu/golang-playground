package main

import (
	"fmt"

	"github.com/zisuu/golang-playground/internal/dictionary"
)

func main() {
	dict := dictionary.Dictionary{"test": "this is just a test"}
	fmt.Println(dict.Search("test"))
}
