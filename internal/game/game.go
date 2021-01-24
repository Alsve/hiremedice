package game

import (
	"errors"
	"fmt"
	"strings"

	"github.com/alsve/hiremedice/internal/player"
)

// New will create a new game
func New() *Game {
	g := &Game{ps: []player.Player{}}
	return g
}

var errFailedAddPlayerPoint = errors.New("game: failed add player point")

// Game represents game event
type Game struct {
	ps    []player.Player
	nTurn int
}

// AddPlayer will add player to the game
func (g *Game) AddPlayer(p *player.Player, ps ...*player.Player) {
	g.ps = append(g.ps, *p)
	for _, pp := range ps {
		g.ps = append(g.ps, *pp)
	}
}

func (g *Game) commitAddLaterDices() {
	for i, p := range g.ps {
		p.CommitAddLaterDices()
		g.ps[i] = p
	}
}

var playerRollDices = func(p *player.Player) {
	p.RollDices()
}

func (g *Game) transferOnesToNextPlayer(currIdx int) {
	ones := g.ps[currIdx].RemoveOneValuedDices()
	nextIdx := currIdx + 1
	if nextIdx >= len(g.ps) {
		nextIdx = 0
	}
	g.ps[nextIdx].AddDicesLater(ones...)
}

// PlayTurn rolls all players's dices
func (g *Game) PlayTurn() {
	for currIdx, p := range g.ps {
		playerRollDices(&p)
		g.ps[currIdx] = p
	}
	g.nTurn++
}

// Evaluate add points for player and transfer ones to their respective
// next player
func (g *Game) Evaluate() {
	for currIdx, p := range g.ps {
		p.AddPointFromSixValuedDices()
		g.ps[currIdx] = p
		g.transferOnesToNextPlayer(currIdx)
	}
	g.commitAddLaterDices()
}

// PlayTurnAndEvaluate represents game play turn with evaluation
func (g *Game) PlayTurnAndEvaluate() {
	for currIdx, p := range g.ps {
		playerRollDices(&p)
		p.AddPointFromSixValuedDices()
		g.ps[currIdx] = p
		g.transferOnesToNextPlayer(currIdx)
	}
	g.commitAddLaterDices()
	g.nTurn++
}

// TurnCount returns game turn count that has been taken
func (g *Game) TurnCount() int {
	return g.nTurn
}

// IsGameOver checks if the games cannot play another turn
func (g *Game) IsGameOver() bool {
	c := 0
	for _, p := range g.ps {
		if len(p.Dices) > 0 {
			c++
			if c > 1 {
				return true
			}
		}
	}
	return false
}

// RemainingPlayers will show player(s) that still has dice to play with
func (g *Game) RemainingPlayers() (indexes []int) {
	indexes = []int{}
	for i, p := range g.ps {
		if len(p.Dices) > 0 {
			indexes = append(indexes, i)
		}
	}
	return indexes
}

// WinningPlayersIndexes will returns player(s) who has the highest point
func (g *Game) WinningPlayersIndexes() (indexes []int) {
	maxPoint := 0
	indexes = []int{}
	for i := range g.ps {
		if p := g.ps[i].Point; p > maxPoint {
			maxPoint = p
		}
	}
	for i := range g.ps {
		if p := g.ps[i].Point; p == maxPoint {
			indexes = append(indexes, i)
		}
	}
	return indexes
}

func (g *Game) String() string {
	sb := strings.Builder{}
	for i, p := range g.ps {
		playerStr := fmt.Sprintf("\t%s\n", p.String())
		playerStrWithIdx := fmt.Sprintf(playerStr, i+1)
		sb.WriteString(playerStrWithIdx)
	}
	return sb.String()
}
