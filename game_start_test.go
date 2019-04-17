package knightmoves

import "testing"

func TestGameOneMove(t *testing.T) {
	g, _ := NewGame("A1", "B3", 10, true)
	err := g.Start()
	if err != nil {
		t.Logf("expected no error, got [%s]", err)
		t.FailNow()
	}

	if nil == g.Solution {
		t.Log("expected solution")
		t.FailNow()
	}

	if len(g.Solution) != 2 {
		t.Log("there is a one step solution")
		t.Log(g.Solution)
		t.FailNow()
	}
}

func TestGameTargetEqualsInitial(t *testing.T) {
	g, _ := NewGame("A1", "A1", 10, true)
	err := g.Start()
	if err != nil {
		t.Logf("expected no error, got [%s]", err)
		t.FailNow()
	}

	if nil == g.Solution {
		t.Log("expected solution")
		t.FailNow()
	}

	if len(g.Solution) != 1 {
		t.Log("there is a zero step solution")
		t.FailNow()
	}
}
