package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Dave")

	got := buffer.String()
	want := "Hello, Dave!"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
