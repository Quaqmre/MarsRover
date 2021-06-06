package compass

import (
	"errors"
	"log"
)

var directionBigConstant int = 8
var directionLowConstant int = 1

type Compass struct {
	Direction       int
	DirectionString string
}

func NewCompass(d rune) (*Compass, error) {
	compass := Compass{}
	switch d {
	case 'N':
		compass.Direction = 1 << 3
	case 'E':
		compass.Direction = 1 << 2
	case 'S':
		compass.Direction = 1 << 1
	case 'W':
		compass.Direction = 1
	default:
		log.Println("Wrong Direction: ", string(d))
		return nil, errors.New("wrong begin directive")
	}
	return &compass, nil
}

func (c *Compass) TurnCompass(t rune) error {
	switch t {
	case 'R':
		c.Direction = c.Direction >> 1
		if c.Direction == 0 {
			c.Direction = directionBigConstant
		}
		return nil
	case 'L':
		c.Direction = c.Direction << 1
		if c.Direction > directionBigConstant {
			c.Direction = directionLowConstant
		}
		return nil
	default:
		log.Println("Wrong turn Direction: ", string(t))
		return errors.New("Wrong turn Direction: " + string(t))
	}
}
