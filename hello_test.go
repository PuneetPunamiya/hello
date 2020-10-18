package main

import "testing"

func TestHelloWorld(t *testing.T) {
	expected := "Hello World, learn Go the hard way"

	actual := hello()
	if expected != actual {
		t.Fatalf("expected %v got %v", expected, actual)
	}
}
