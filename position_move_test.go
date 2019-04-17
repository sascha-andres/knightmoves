package knightmoves

import "testing"

func TestInvalidMove(t *testing.T) {
	pos := &Position{0, 0}
	_, err := pos.Move(-1, -3)

	if err != ErrMoveInvalid {
		t.Logf("expected ErrMoveInvalid, got %s", err)
		t.Fail()
	}
}

func TestInvalidTarget(t *testing.T) {
	pos := &Position{0, 0}
	_, err := pos.Move(-1, -2)

	if err != ErrTargetPositionInvalid {
		t.Logf("expected ErrTargetPositionInvalid, got %s", err)
		t.Fail()
	}
}

func TestValidMove(t *testing.T) {
	pos := &Position{0, 0}
	newPosition, err := pos.Move(1, 2)
	if err != nil {
		t.Logf("got error %s, expected none", err)
		t.FailNow()
	}
	if !(newPosition.X == 1 && newPosition.Y == 2) {
		t.Logf("Expected [B:3], got %s", newPosition)
	}
}
