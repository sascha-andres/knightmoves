package knightmoves

import "math"

// MoveValid tests if the move is valid for a knight
func MoveValid(x, y int) bool {
	return (math.Abs(float64(x)) == 2 || math.Abs(float64(y)) == 2) && (math.Abs(float64(x)) == 1 || math.Abs(float64(y)) == 1)
}
