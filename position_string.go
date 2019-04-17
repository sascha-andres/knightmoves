package knightmoves

import "fmt"

var literalExpression = []string{"A", "B", "C", "D", "E", "F", "G", "H"}

// String returns a textual representation of a position
func (pos *Position) String() string {
	return fmt.Sprintf("%s%d", literalExpression[pos.X], pos.Y+1)
}
