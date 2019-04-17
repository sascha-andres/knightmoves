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
		err := game.SetSolution(positions)
		if err != nil {
			logger.Warnf("provided wrong solution: %#v", positions)
		}
	} else {
		var err error
		for _, firstLength := range moveDescription {
			if !(len(game.Solution) == 0) && len(positions) >= len(game.Solution) {
				continue
			}
			lastPosition := positions[len(positions)-1]
			secondLength := 1
			if firstLength == 1 {
				secondLength = 2
			}
			game.makeMove(lastPosition, firstLength, secondLength, positions, err, depth)
			game.makeMove(lastPosition, firstLength, secondLength*-1, positions, err, depth)
			game.makeMove(lastPosition, firstLength*-1, secondLength, positions, err, depth)
			game.makeMove(lastPosition, firstLength*-1, secondLength*-1, positions, err, depth)
		}
	}

	return nil
}
