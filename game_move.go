package knightmoves

import (
	"errors"
)

var ErrMaxDepthReached = errors.New("error reached max depth")

var moveDescription = []int{2, 1}

// Move starts to test positions
func (game *Game) Move(depth uint, positions ...Position) error {
	logger := game.baseLogger.WithField("func", "Move")
	logger.Debugf("depth := %d positions := %d positions := %#v", depth, len(positions), positions)

	if depth >= game.MaxDepth {
		return ErrMaxDepthReached
	}

	if positions[len(positions)-1].Equals(game.Target) {
		if nil == game.Solution || len(game.Solution) >= len(positions) {
			game.Solution = positions
		}
	} else {
		var err error
		for _, firstLength := range moveDescription {
			lastPosition := positions[len(positions)-1]
			secondLength := 1
			if firstLength == 1 {
				secondLength = 2
			}
			p, e := lastPosition.Move(firstLength, secondLength)
			if e == nil && !positionsContain(p, positions...) && err == nil {
				err = game.Move(depth+1, append(positions, *p)...)
			}
			p, e = lastPosition.Move(firstLength, secondLength*-1)
			if e == nil && !positionsContain(p, positions...) && err == nil {
				err = game.Move(depth+1, append(positions, *p)...)
			}
			p, e = lastPosition.Move(firstLength*-1, secondLength)
			if e == nil && !positionsContain(p, positions...) && err == nil {
				err = game.Move(depth+1, append(positions, *p)...)
			}
			p, e = lastPosition.Move(firstLength*-1, secondLength*-1)
			if e == nil && !positionsContain(p, positions...) && err == nil {
				err = game.Move(depth+1, append(positions, *p)...)
			}
		}
	}

	return nil
}

func positionsContain(newPosition *Position, positions ...Position) bool {
	for _, p := range positions {
		if p.Equals(newPosition) {
			return true
		}
	}
	return false
}
