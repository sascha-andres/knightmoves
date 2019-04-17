package knightmoves

import "testing"

func TestZeroZero(t *testing.T) {
	pos := Position{X: 0, Y: 0}
	result := pos.String()
	if result != "A1" {
		t.Logf("Expected [A1], got %s", result)
		t.FailNow()
	}
}

func TestSevenTwo(t *testing.T) {
	pos := Position{X: 7, Y: 2}
	result := pos.String()
	if result != "H3" {
		t.Logf("Expected [H3], got %s", result)
		t.FailNow()
	}
}
