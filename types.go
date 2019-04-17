package knightmoves

import "github.com/sirupsen/logrus"

type (
	// Game is a container to run the search
	Game struct {
		Initial  *Position
		Target   *Position
		MaxDepth uint
		Solution []Position

		baseLogger *logrus.Entry
	}

	// Position is a position for a knight
	Position struct {
		X int
		Y int
	}
)
