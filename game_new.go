package knightmoves

import "github.com/sirupsen/logrus"

// NewGame initialises a find path game
func NewGame(initial, target string, maxDepth uint, verbose bool) (*Game, error) {
	game := &Game{
		MaxDepth:   maxDepth,
		baseLogger: logrus.WithField("package", "knightmoves"),
	}

	logrus.SetLevel(logrus.InfoLevel)
	if verbose {
		logrus.SetLevel(logrus.DebugLevel)
	}

	logger := game.baseLogger.WithField("func", "NewGame")
	logger.Infof("initial := [%s] target := [%s], maxDepth := %d, verbose := %t", initial, target, maxDepth, verbose)

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
