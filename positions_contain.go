package knightmoves

// positionsContain checks if a position is part of a list of positions
func positionsContain(newPosition *Position, positions ...Position) bool {
	for _, p := range positions {
		if p.Equals(newPosition) {
			return true
		}
	}
	return false
}
