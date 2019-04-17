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
