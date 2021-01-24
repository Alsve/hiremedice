package main

import (
	"fmt"

	"github.com/alsve/hiremedice/internal/dice"
	"github.com/alsve/hiremedice/internal/game"
	"github.com/alsve/hiremedice/internal/player"
)

func main() {
	var playerCount, diceCount int
	g := game.New()
	fmt.Scanf("%d", &playerCount)
	fmt.Scanf("%d", &diceCount)

	fmt.Printf("Pemain = %d, Dadu = %d\n", playerCount, diceCount)
	for i := 0; i < playerCount; i++ {
		p := player.New()
		for j := 0; j < diceCount; j++ {
			p.AddDice(&dice.Dice{Number: 1})
		}
		g.AddPlayer(p)
	}

	// Game loop start
	for g.IsGameOver() {
		fmt.Println("==================")
		g.PlayTurn()
		fmt.Printf("Giliran %d lempar dadu:\n%s", g.TurnCount(), g.String())
		g.Evaluate()
		fmt.Printf("Setelah evaluasi:\n%s", g.String())
	}

	// Game result
	fmt.Println("==================")
	fmt.Printf(
		"Game berakhir karena hanya pemain #%d yang memiliki dadu.\n",
		g.RemainingPlayers()[0]+1,
	)
	fmt.Printf(
		"Game dimenangkan oleh pemain #%d karena memiliki poin lebih banyak dari pemain lainnya.\n",
		g.WinningPlayersIndexes()[0]+1,
	)
}
