package knightmoves

// Start kicks off search for solution
func (game *Game) Start() error {
	logger := game.baseLogger.WithField("func", "Start")
	logger.Debug("Start search")

	err := game.Move(0, *game.Initial)
	if err != nil {
		return err
	}

	return nil
}
