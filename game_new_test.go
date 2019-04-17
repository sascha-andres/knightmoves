package knightmoves

import "testing"

func TestGameA1H8(t *testing.T) {
	g, err := NewGame("A1", "H8", 20, true)

	if err != nil {
		t.Logf("expected no error, got [%s]", err)
		t.FailNow()
	}

	if g.Initial.X != 0 || g.Initial.Y != 0 {
		t.Logf("expected 0-0 for initial position, got %d-%d", g.Initial.X, g.Initial.Y)
		t.FailNow()
	}

	if g.Target.X != 7 || g.Target.Y != 7 {
		t.Logf("expected 7-7 for target position, got %d-%d", g.Target.X, g.Target.Y)
		t.FailNow()
	}
}

func TestGameInvalidInitialPosition(t *testing.T) {
	_, err := NewGame("Z1", "H8", 20, true)

	if err != ErrInvalidFormat {
		t.Logf("expected ErrInvalidFormat, got [%s]", err)
		t.FailNow()
	}
}

func TestGameInvalidTargetPosition(t *testing.T) {
	_, err := NewGame("A1", "Z8", 20, true)

	if err != ErrInvalidFormat {
		t.Logf("expected ErrInvalidFormat, got [%s]", err)
		t.FailNow()
	}
}
