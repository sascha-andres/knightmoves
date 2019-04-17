package knightmoves

type (
	// Game is a container to run the search
	Game struct {
		Initial  Position
		Target   Position
		MaxDepth int
		Verbose  bool

		solutions [][]Position
	}

	// Position is a position for a knight
	Position struct {
		X int
		Y int
	}
)
