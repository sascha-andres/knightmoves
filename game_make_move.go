package knightmoves

// makeMove handles a move with the knight
func (game *Game) makeMove(lastPosition Position, firstLength int, secondLength int, positions []Position, err error, depth uint) {
	if len(game.Solution) != 0 && len(positions)+1 >= len(game.Solution) {
		return
	}
	p, e := lastPosition.Move(firstLength, secondLength)
	if e == nil && !positionsContain(p, positions...) && err == nil {
		err = game.Move(depth+1, append(positions, *p)...)
	}
}
