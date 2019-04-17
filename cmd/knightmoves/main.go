package main

import (
	"fmt"
	"github.com/integrii/flaggy"
	"github.com/sirupsen/logrus"
	"livingit.de/code/knightmoves"
	"os"
)

func main() {
	var (
		initialPosition string
		targetPosition  string
		maxDepth        uint
		verbose         bool
	)

	flaggy.String(&initialPosition, "i", "initial", "initial position")
	flaggy.String(&targetPosition, "t", "target", "target position")
	flaggy.UInt(&maxDepth, "d", "depth", "maximum depth to analyse")
	flaggy.Bool(&verbose, "v", "verbose", "change log level")

	flaggy.Parse()

	g, err := knightmoves.NewGame(initialPosition, targetPosition, maxDepth, verbose)
	if err != nil {
		logrus.Warnf("error creating engine: %d", err)
		os.Exit(1)
	}

	err = g.Start()
	if err != nil {
		logrus.Warnf("error running analysis: %d", err)
		os.Exit(1)
	}

	if len(g.Solution) > 0 {
		fmt.Printf("%d moves:\n", len(g.Solution))
		for idx, pos := range g.Solution {
			fmt.Printf("  %3d: %s\n", idx+1, pos.String())
		}
	}
}
