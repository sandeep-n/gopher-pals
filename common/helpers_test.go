package common

import "testing"

func TestGreet(t *testing.T) {
	got := greet("Sandeep")
	want := "Hello, Sandeep"
	if got != want {
		t.Errorf("Failed to greet.")
	}
}