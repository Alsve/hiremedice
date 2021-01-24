package player

import (
	"fmt"
	"testing"

	"github.com/alsve/hiremedice/internal/dice"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	p := New()
	assert.NotNil(t, p.Dices)
	assert.NotNil(t, p.addLater)
}

func TestPlayer_AddDice_expected(t *testing.T) {
	expected := []dice.Dice{{Number: 5}}
	input := &dice.Dice{Number: 5}

	defer func(orig func(*dice.Dice) error) {
		validateDice = orig
	}(validateDice)

	validateDice = func(d *dice.Dice) error {
		return nil
	}

	// exec the function
	p := Player{Dices: []dice.Dice{}}
	err := p.AddDice(input)
	assert.NoError(t, err)
	assert.Equal(t, p.Dices, expected)
}

func TestPlayer_AddDice_invalidDiceNumber(t *testing.T) {
	expected := []dice.Dice{}
	input := &dice.Dice{Number: 7}

	defer func(orig func(*dice.Dice) error) {
		validateDice = orig
	}(validateDice)

	validateDice = func(d *dice.Dice) error {
		return dice.ErrInvalidDiceNumber
	}

	// exec the function
	p := Player{Dices: []dice.Dice{}}
	err := p.AddDice(input)
	assert.Equal(t, err, errInsertDiceFailed)
	assert.Equal(t, p.Dices, expected)
}

func TestPlayer_AddDice_otherError(t *testing.T) {
	expected := []dice.Dice{}
	input := &dice.Dice{Number: 7}

	defer func(orig func(*dice.Dice) error) {
		validateDice = orig
	}(validateDice)

	validateDice = func(d *dice.Dice) error {
		return assert.AnError
	}

	// exec the function
	p := Player{Dices: []dice.Dice{}}
	err := p.AddDice(input)
	assert.Equal(t, assert.AnError, err)
	assert.Equal(t, p.Dices, expected)
}

func TestPlayer_RemoveDice_expected(t *testing.T) {
	inputIdx := 0
	expectedP := []dice.Dice{{Number: 2}}
	expectedReturn := &dice.Dice{Number: 6}

	// exec the function
	p := Player{Dices: []dice.Dice{{Number: 6}, {Number: 2}}}
	val, err := p.RemoveDice(inputIdx)
	assert.NoError(t, err)
	assert.Equal(t, expectedP, p.Dices)
	assert.Equal(t, expectedReturn, val)
}

func TestPlayer_RemoveDice_indexError(t *testing.T) {
	inputIdx := 2
	expected := []dice.Dice{{Number: 6}, {Number: 2}}

	// exec the function
	p := Player{Dices: []dice.Dice{{Number: 6}, {Number: 2}}}
	val, err := p.RemoveDice(inputIdx)
	assert.Equal(t, ErrIndexOutOfBound, err)
	assert.Nil(t, val)
	assert.Equal(t, expected, p.Dices)
}

func TestPlayer_AddDiceLater_expected(t *testing.T) {
	input := &dice.Dice{Number: 5}
	expected := []dice.Dice{{Number: 5}}

	defer func(orig func(*dice.Dice) error) {
		validateDice = orig
	}(validateDice)

	validateDice = func(d *dice.Dice) error {
		return nil
	}

	// exec the function
	p := &Player{Dices: []dice.Dice{}}
	err := p.AddDicesLater(input)
	assert.NoError(t, err)
	assert.Equal(t, expected, p.addLater)
}

func TestPlayer_AddDiceLater_nilParam(t *testing.T) {
	var input []*dice.Dice
	expected := []dice.Dice{}

	p := Player{addLater: []dice.Dice{}}
	err := p.AddDicesLater(input...)
	assert.Equal(t, expected, p.addLater)
	assert.Equal(t, errInsertDiceFailed, err)
}

func TestPlayer_AddDicesLater_invalidDiceNumber(t *testing.T) {
	input := &dice.Dice{Number: 10}
	expected := []dice.Dice{}
	expectedErr := fmt.Errorf("error at index %d: %s", 0, errInsertDiceFailed)

	defer func(orig func(*dice.Dice) error) {
		validateDice = orig
	}(validateDice)

	validateDice = func(d *dice.Dice) error {
		return dice.ErrInvalidDiceNumber
	}

	// exec the function
	p := Player{addLater: []dice.Dice{}}
	err := p.AddDicesLater(input)
	assert.Equal(t, expectedErr, err)
	assert.Equal(t, expected, p.addLater)
}

func TestPlayer_AddDicesLater_otherError(t *testing.T) {
	input := &dice.Dice{Number: 10}
	expected := []dice.Dice{}
	expectedErr := fmt.Errorf("error at index %d: %s", 0, assert.AnError.Error())

	defer func(orig func(*dice.Dice) error) {
		validateDice = orig
	}(validateDice)

	validateDice = func(d *dice.Dice) error {
		return assert.AnError
	}

	// exec the function
	p := &Player{addLater: []dice.Dice{}}
	err := p.AddDicesLater(input)
	assert.Equal(t, expectedErr, err)
	assert.Equal(t, expected, p.addLater)
}

func TestPlayer_CommitDice_expected(t *testing.T) {
	expectedAddLater := []dice.Dice{}
	expectedDices := []dice.Dice{{Number: 3}, {Number: 5}}

	// exec the function
	p := Player{
		Dices:    []dice.Dice{{Number: 3}},
		addLater: []dice.Dice{{Number: 5}},
	}
	p.CommitAddLaterDices()
	assert.Equal(t, expectedAddLater, p.addLater)
	assert.Equal(t, expectedDices, p.Dices)
}

func TestPlayer_AddPointFromSixValuedDices_expected(t *testing.T) {
	expectedDices := []dice.Dice{{Number: 1}}

	// exec the function
	p := &Player{Dices: []dice.Dice{{Number: 6}, {Number: 1}, {Number: 6}}}
	p.AddPointFromSixValuedDices()
	assert.Equal(t, 2, p.Point)
	assert.Equal(t, expectedDices, p.Dices)
}

func TestPlayer_RemoveOneValuedDices_expected(t *testing.T) {
	expectedDices := []dice.Dice{{Number: 6}, {Number: 5}, {Number: 3}}
	expectedRet := []*dice.Dice{{Number: 1}, {Number: 1}}

	// exec the function
	p := &Player{Dices: []dice.Dice{
		{Number: 6}, {Number: 1}, {Number: 5}, {Number: 1}, {Number: 3},
	}}
	got := p.RemoveOneValuedDices()
	assert.Equal(t, expectedDices, p.Dices)
	assert.Equal(t, expectedRet, got)
}

func TestPlayer_String_4Dices(t *testing.T) {
	expected := "Pemain #%d (3): 3,6,1,3"

	// exec the function
	p := &Player{
		Point: 3,
		Dices: []dice.Dice{{Number: 3}, {Number: 6}, {Number: 1}, {Number: 3}},
	}
	assert.Equal(t, expected, p.String())
}

func TestPlayer_String_noDice(t *testing.T) {
	expected := "Pemain #%d (4): _ (Berhenti bermain karena tidak memiliki dadu)"

	p := &Player{
		Point: 4,
		Dices: []dice.Dice{},
	}
	assert.Equal(t, expected, p.String())
}
