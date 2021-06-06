package command

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCurrentCommand(t *testing.T) {
	command := "LLMM"
	c := NewCommand(command)

	err := c.GetNextCommand()

	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 'L', c.GetCurrentCommand())
}

func TestGetCurrentCommandIndex(t *testing.T) {
	command := "LLMM"
	c := NewCommand(command)

	err := c.GetNextCommand()
	err = c.GetNextCommand()

	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 2, c.currentCommandIndex)
}
