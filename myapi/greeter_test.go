package myapi

import (
	"testing"
)

func TestGreeting(t *testing.T) {
	expected := "hello"
	actual := Greeting()
	if expected != actual {
		t.Fatalf("Expected \"%v\" but found \"%v\"", expected, actual)
	}
}
