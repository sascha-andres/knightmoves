package knightmoves

import (
	"fmt"
	"testing"
)

func TestInvalidPositionZ1(t *testing.T) {
	_, err := NewFromString("Z1")
	if err != ErrInvalidFormat {
		t.Logf("expected ErrIvalidFormat, got [%s]", err)
		t.FailNow()
	}
}

func TestValidPositionB2(t *testing.T) {
	p, err := NewFromString("B2")
	if err != nil {
		t.Logf("expected no error, got [%s]", err)
		t.FailNow()
	}
	if nil == p {
		t.Logf("expected position, got nil")
		t.FailNow()
	}
	if p.X != 1 || p.Y != 1 {
		t.Logf("Expected X,Y := 1, got X=%d, Y=%d", p.X, p.Y)
		t.FailNow()
	}
}

func TestValidPositionA1(t *testing.T) {
	p, err := NewFromString("A1")
	if err != nil {
		t.Logf("expected no error, got [%s]", err)
		t.FailNow()
	}
	if nil == p {
		t.Logf("expected position, got nil")
		t.FailNow()
	}
	if p.X != 0 || p.Y != 0 {
		t.Logf("Expected X,Y := 1, got X=%d, Y=%d", p.X, p.Y)
		t.FailNow()
	}
}

func TestValidPositionHColumn(t *testing.T) {
	for column := 1; column < 9; column++ {
		positionName := fmt.Sprintf("H%d", column)
		t.Run(positionName, func(t *testing.T) {
			p, err := NewFromString(positionName)
			if err != nil {
				t.Logf("expected no error, got [%s]", err)
				t.FailNow()
			}
			if nil == p {
				t.Logf("expected position, got nil")
				t.FailNow()
			}
			if p.X != 7 || p.Y != column-1 {
				t.Logf("Expected X=7, Y=%d, got X=%d, Y=%d", column-1, p.X, p.Y)
				t.FailNow()
			}
		})
	}
}

func TestValidPosition3Row(t *testing.T) {
	for idx, char := range "ABCDEFGH" {
		positionName := fmt.Sprintf("%s3", string(char))
		t.Run(positionName, func(t *testing.T) {
			p, err := NewFromString(positionName)
			if err != nil {
				t.Logf("expected no error, got [%s]", err)
				t.FailNow()
			}
			if nil == p {
				t.Logf("expected position, got nil")
				t.FailNow()
			}
			if p.X != idx || p.Y != 2 {
				t.Logf("Expected X=%d, Y=2, got X=%d, Y=%d", idx, p.X, p.Y)
				t.FailNow()
			}
		})
	}
}
