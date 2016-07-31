package myapi

import (
	"testing"
)

func TestGreeting(t *testing.T) {
	expected := GREET
	actual := Greeting()
	if expected != actual {
		t.Fatalf("Expected \"%v\" but found \"%v\"", expected, actual)
	}
}

func TestGreet(t *testing.T) {
	expected := "hello testName"
	actual := Greet("testName")
	if expected != actual {
		t.Fatalf("Expected \"%v\" but found \"%v\"", expected, actual)
	}
}
