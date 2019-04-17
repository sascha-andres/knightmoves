package knightmoves

import "testing"

func TestValidPosition(t *testing.T) {
	pos := Position{X: 0, Y: 0}
	if !pos.Valid() {
		t.FailNow()
	}
}

func TestInvalidXLowerZero(t *testing.T) {
	pos := Position{X: -1, Y: 0}
	if pos.Valid() {
		t.FailNow()
	}
}

func TestInvalidYLowerZero(t *testing.T) {
	pos := Position{X: 0, Y: -1}
	if pos.Valid() {
		t.FailNow()
	}
}

func TestInvalidXHigherSeven(t *testing.T) {
	pos := Position{X: 8, Y: 0}
	if pos.Valid() {
		t.FailNow()
	}
}

func TestInvalidYHigherSeven(t *testing.T) {
	pos := Position{X: 0, Y: 8}
	if pos.Valid() {
		t.FailNow()
	}
}
