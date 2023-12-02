package direction

import (
	"errors"
	"regexp"
	"strconv"
)

type Type string

const (
	FORWARD = "forward"
	DOWN    = "down"
	UP      = "up"
)

var commandRegExpr = regexp.MustCompile("\\w+")

type Command struct {
	Direction Type
	Distance  int
}

func NewCommand(commandString string) (*Command, error) {
	matches := commandRegExpr.FindAllString(commandString, -1)

	if len(matches) != 2 {
		return nil, errors.New("command string should contain 2 values")
	}

	dist, err := strconv.Atoi(matches[1])
	if err != nil {
		return nil, err
	}

	var dir Type

	switch matches[0] {
	case FORWARD:
		dir = FORWARD
		break
	case DOWN:
		dir = DOWN
		break
	case UP:
		dir = UP
		break
	default:
		return nil, errors.New("direction type does not exist")
	}

	command := &Command{
		Direction: dir,
		Distance:  dist,
	}

	return command, nil
}
