package dice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDice_Roll_expected(t *testing.T) {
	dice := Dice{}
	expected := 6

	defer func(orig func(n int) int) {
		randomInt = orig
	}(randomInt)

	randomInt = func(n int) int {
		return 5
	}

	got := dice.Roll()
	assert.Equal(t, expected, got)
	assert.Equal(t, expected, dice.Number)
}

func TestDice_Validate_expected(t *testing.T) {
	dice := Dice{Number: 6}
	assert.NoError(t, dice.Validate())
}

func TestDice_Validate_error(t *testing.T) {
	dice := Dice{Number: 10}
	err := dice.Validate()
	assert.Error(t, err)
	assert.Equal(t, ErrInvalidDiceNumber, err)
}
