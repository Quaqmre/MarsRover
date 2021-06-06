package compass

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCompass(t *testing.T) {
	_, err := NewCompass('T')

	if err == nil {
		t.Fatal("err cannot be null with wrong parameter")
	}
}

func TestTurnCompasss(t *testing.T) {
	comps, err := NewCompass('N')

	if err != nil {
		t.Fatal("unexpected error check error cases")
	}

	comps.TurnCompass('L')
	comps.TurnCompass('L')

	assert.Equal(t, comps.Direction, 2)

}
