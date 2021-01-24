package game

import (
	"testing"

	"github.com/alsve/hiremedice/internal/dice"
	"github.com/alsve/hiremedice/internal/player"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	g := New()
	assert.NotNil(t, g.ps)
}

func TestGame_AddPlayer_addOne(t *testing.T) {
	input := &player.Player{Dices: []dice.Dice{{Number: 3}}}
	expected := []player.Player{{Dices: []dice.Dice{{Number: 3}}}}

	// exec the function
	g := &Game{ps: []player.Player{}}
	g.AddPlayer(input)
	assert.Equal(t, expected, g.ps)
}

func TestGame_AddPlayer_addThree(t *testing.T) {
	input := []*player.Player{
		{Dices: []dice.Dice{{Number: 2}}},
		{Dices: []dice.Dice{{Number: 5}}},
		{Dices: []dice.Dice{{Number: 4}}},
	}
	expected := []player.Player{
		{Dices: []dice.Dice{{Number: 2}}},
		{Dices: []dice.Dice{{Number: 5}}},
		{Dices: []dice.Dice{{Number: 4}}},
	}

	// exec the function
	g := &Game{ps: []player.Player{}}
	g.AddPlayer(input[0], input[1:]...)
	assert.Equal(t, expected, g.ps)
}

func TestGame_IsGameOver_tableTest(t *testing.T) {
	scenarios := []struct {
		desc     string
		input    []player.Player
		expected bool
	}{
		{
			desc: "the game is over",
			input: []player.Player{
				{Dices: []dice.Dice{{}, {}}},
				{Dices: []dice.Dice{{}}},
			},
			expected: true,
		},
		{
			desc:     "the game still on",
			input:    []player.Player{{Dices: []dice.Dice{{}, {}, {}}}},
			expected: false,
		},
	}

	for _, scenario := range scenarios {
		g := &Game{ps: scenario.input}
		got := g.IsGameOver()
		assert.Equal(t, scenario.expected, got, scenario.desc)
	}
}

func TestGame_WinningPlayersIndexes_expected(t *testing.T) {
	expected := []int{0, 2}

	g := &Game{
		ps: []player.Player{{Point: 3}, {Point: 2}, {Point: 3}},
	}
	got := g.WinningPlayersIndexes()
	assert.Equal(t, expected, got)
}

func TestGame_PlayTurn_2Player2Dice(t *testing.T) {
	p1exp := player.New()
	p1exp.Dices = []dice.Dice{{Number: 3}, {Number: 2}}
	p2exp := player.New()
	p2exp.Dices = []dice.Dice{{Number: 6}, {Number: 1}}
	expected := &Game{nTurn: 1, ps: []player.Player{*p1exp, *p2exp}}

	defer func(orig func(*player.Player)) {
		playerRollDices = orig
	}(playerRollDices)

	// replace playerRollDices to function generator to make test simpler
	generateIdx := 0
	generatorSet := [][]dice.Dice{
		{{Number: 3}, {Number: 2}},
		{{Number: 6}, {Number: 1}},
	}
	playerRollDices = func(p *player.Player) {
		p.Dices = generatorSet[generateIdx]
		generateIdx++
	}

	// exec the function
	p1 := player.New()
	p1.Dices = []dice.Dice{{}, {}}
	p2 := player.New()
	p2.Dices = []dice.Dice{{}, {}}
	g := &Game{ps: []player.Player{*p1, *p2}}
	g.PlayTurn()
	assert.Equal(t, expected, g)
}

func TestGame_Evaluate_2Player2Dice(t *testing.T) {
	p1exp := player.New()
	p1exp.Dices = []dice.Dice{{Number: 3}, {Number: 2}, {Number: 1}}
	p2exp := player.New()
	p2exp.Point = 1
	p2exp.Dices = []dice.Dice{}
	expected := &Game{ps: []player.Player{*p1exp, *p2exp}}

	defer func(orig func(*player.Player)) {
		playerRollDices = orig
	}(playerRollDices)

	// replace playerRollDices to function generator to make test simpler
	generateIdx := 0
	generatorSet := [][]dice.Dice{
		{{Number: 3}, {Number: 2}},
		{{Number: 6}, {Number: 1}},
	}
	playerRollDices = func(p *player.Player) {
		p.Dices = generatorSet[generateIdx]
		generateIdx++
	}

	// exec the function
	p1 := player.New()
	p1.Dices = []dice.Dice{{Number: 3}, {Number: 2}}
	p2 := player.New()
	p2.Dices = []dice.Dice{{Number: 6}, {Number: 1}}
	g := &Game{ps: []player.Player{*p1, *p2}}
	g.Evaluate()
	assert.Equal(t, expected, g)
}

func TestGame_PlayTurnAndEvaluate_2Player2Dice(t *testing.T) {
	p1exp := player.New()
	p1exp.Dices = []dice.Dice{{Number: 3}, {Number: 2}, {Number: 1}}
	p2exp := player.New()
	p2exp.Point = 1
	expected := &Game{nTurn: 1, ps: []player.Player{*p1exp, *p2exp}}

	defer func(orig func(*player.Player)) {
		playerRollDices = orig
	}(playerRollDices)

	// replace playerRollDices to function generator to make test simpler
	generateIdx := 0
	generatorSet := [][]dice.Dice{
		{{Number: 3}, {Number: 2}},
		{{Number: 6}, {Number: 1}},
	}
	playerRollDices = func(p *player.Player) {
		p.Dices = generatorSet[generateIdx]
		generateIdx++
	}

	// exec the function
	p1 := player.New()
	p1.Dices = []dice.Dice{{}, {}}
	p2 := player.New()
	p2.Dices = []dice.Dice{{}, {}}
	g := &Game{ps: []player.Player{*p1, *p2}}
	g.PlayTurnAndEvaluate()
	assert.Equal(t, expected, g)
}

func TestGame_String_expected(t *testing.T) {
	expected := "\tPemain #1 (2): 2\n\tPemain #2 (3): 5\n\tPemain #3 (4): 5\n"

	g := &Game{
		ps: []player.Player{
			{Point: 2, Dices: []dice.Dice{{Number: 2}}},
			{Point: 3, Dices: []dice.Dice{{Number: 5}}},
			{Point: 4, Dices: []dice.Dice{{Number: 5}}},
		},
	}
	assert.Equal(t, expected, g.String())
}
