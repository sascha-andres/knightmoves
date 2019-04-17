package knightmoves

// Valid returns true if a position is valid on the chess board
func (pos *Position) Valid() bool {
	if pos.X < 0 || pos.Y < 0 {
		return false
	}
	if pos.X > 7 || pos.Y > 7 {
		return false
	}
	return true
}
