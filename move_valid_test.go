package knightmoves

import "testing"

func TestValidMove21(t *testing.T) {
	if !MoveValid(2, 1) {
		t.Fail()
	}
}

func TestValidMove12(t *testing.T) {
	if !MoveValid(1, 2) {
		t.Fail()
	}
}

func TestValidMoveM2M1(t *testing.T) {
	if !MoveValid(-2, -1) {
		t.Fail()
	}
}

func TestValidMoveM1M2(t *testing.T) {
	if !MoveValid(-1, -2) {
		t.Fail()
	}
}
