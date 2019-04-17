package knightmoves

// NewGame initialises a find path game
func NewGame(initial, target string, maxDepth uint, verbose bool) (*Game, error) {
	game := &Game{
		MaxDepth: maxDepth,
		Verbose:  verbose,
	}

	initialPosition, err := NewFromString(initial)
	if err != nil {
		return nil, err
	}
	game.Initial = initialPosition

	targetPosition, err := NewFromString(target)
	if err != nil {
		return nil, err
	}
	game.Target = targetPosition

	return game, nil
}
