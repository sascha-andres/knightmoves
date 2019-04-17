package knightmoves

import "errors"

var ErrMoveInvalid = errors.New("move is not a valid move for a knight")
var ErrTargetPositionInvalid = errors.New("target position not valid")

// Move sets knight to new position
func (pos *Position) Move(x, y int) (*Position, error) {
	if !MoveValid(x, y) {
		return nil, ErrMoveInvalid
	}
	newPosition := &Position{X: pos.X + x, Y: pos.Y + y}
	if !newPosition.Valid() {
		return nil, ErrTargetPositionInvalid
	}
	return newPosition, nil
}
