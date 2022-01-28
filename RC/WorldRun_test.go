package main

import (
	"testing"
)
func TestMoveBoat(t *testing.T) {
	if MoveBoat("West") != "East" {
		t.error("From West to East, not this way:)")
	} 
	else if MoveBoat("East") != "West" {
		t.error("From East to West, not this way:)")
	}
}