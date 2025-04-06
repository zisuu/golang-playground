package main

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("Simple add", func(t *testing.T) {
		got := Repeat("a", 5)
		want := "aaaaa"
		assertCorrectMessage(t, got, want)
	})
}

func ExampleRepeat() {
	want := Repeat("x", 3)
	fmt.Println(want)
	// Output: xxx
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", b.N)
	}
}
