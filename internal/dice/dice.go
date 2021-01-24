package dice

import (
	"errors"
	"math/rand"
	"time"

	"github.com/alsve/hiremedice/internal/logger"
)

func init() {
	rand.Seed(time.Now().Unix())
}

// ErrInvalidDiceNumber will make
var ErrInvalidDiceNumber = errors.New("dice: invalid number")

// Dice represents dice
type Dice struct {
	Number int
}

var randomInt = rand.Intn

// Roll get random integer and save it to struct
func (d *Dice) Roll() int {
	num := randomInt(6)
	d.Number = num + 1
	return num + 1
}

// Validate will validate the dice number
func (d *Dice) Validate() error {
	if d.Number <= 0 || d.Number > 6 {
		logger.L.Error(ErrInvalidDiceNumber.Error())
		return ErrInvalidDiceNumber
	}
	return nil
}
