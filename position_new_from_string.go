package knightmoves

import (
	"errors"
	"regexp"
	"strconv"
)

var ErrInvalidFormat = errors.New("invalid format. should [A-H][1-8]")

const parseExpression = "([A-H])([1-8])"

func NewFromString(value string) (*Position, error) {
	r := regexp.MustCompile(parseExpression)
	if !r.MatchString(value) {
		return nil, ErrInvalidFormat
	}

	res := r.FindAllStringSubmatch(value, 2)
	split := res[0][1:]

	result := &Position{}
	val, _ := strconv.ParseInt(split[1], 10, 32)
	result.Y = int(val) - 1
	switch split[0] {
	case "A":
		result.X = 0
		break
	case "B":
		result.X = 1
		break
	case "C":
		result.X = 2
		break
	case "D":
		result.X = 3
		break
	case "E":
		result.X = 4
		break
	case "F":
		result.X = 5
		break
	case "G":
		result.X = 6
		break
	case "H":
		result.X = 7
		break
	}

	return result, nil
}
