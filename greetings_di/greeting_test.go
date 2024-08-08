package greetings_di

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Arjesh")

	got := buffer.String()
	want := "Hello, Arjesh"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
