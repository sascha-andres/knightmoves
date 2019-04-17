package knightmoves

// Equals returns true if positions are equal on the board
func (pos *Position) Equals(other *Position) bool {
	return pos.X == other.X && pos.Y == other.Y
}
