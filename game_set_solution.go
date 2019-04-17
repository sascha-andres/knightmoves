package knightmoves

import (
	"errors"
	"fmt"
)

var ErrInvalidSolution = errors.New("solution is not valid")

// SetSolution guards access for setting a possible solution
func (game *Game) SetSolution(positions []Position) error {
	logger := game.baseLogger.WithField("func", "SetSolution")

	if positions[len(positions)-1].X != game.Target.X || positions[len(positions)-1].Y != game.Target.Y {
		return ErrInvalidSolution
	}

	if len(game.Solution) == 0 || len(positions) < len(game.Solution) {
		var data string
		for _, pos := range positions {
			if len(data) == 0 {
				data = fmt.Sprintf("%s", pos.String())
			} else {
				data = fmt.Sprintf("%s,%s", data, pos.String())
			}
		}
		logger.Infof("received possible solution: %s", data)
		game.Solution = positions
	}

	return nil
}
