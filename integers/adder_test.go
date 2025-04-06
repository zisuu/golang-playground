package main

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	t.Run("Simple add", func(t *testing.T) {
		got := Add(2, 2)
		want := 4
		assertCorrectMessage(t, got, want)
	})
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}

func assertCorrectMessage(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
