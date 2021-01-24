package player

import (
	"errors"
	"fmt"
	"strings"

	"github.com/alsve/hiremedice/internal/dice"
	"github.com/alsve/hiremedice/internal/logger"
)

// ErrIndexOutOfBound is returned by RemoveDice when the provided dice at
// provided i index is exceeded max index.
var ErrIndexOutOfBound = errors.New("player: index out of bound")

var errInsertDiceFailed = errors.New("player: insert dice failed")

// New make a new Player object
func New() *Player {
	p := Player{
		Dices:    []dice.Dice{},
		addLater: []dice.Dice{},
	}
	return &p
}

// Player is a player for the game
type Player struct {
	// Point is player's point in the game
	Point int
	// Dices is value of dice
	Dices []dice.Dice
	// addLater are dices that will be added after execute CommitDice()
	addLater []dice.Dice
}

var validateDice = func(d *dice.Dice) error {
	return d.Validate()
}

var rollDice = func(d *dice.Dice) int {
	return d.Roll()
}

// AddPointFromSixValuedDices will remove six-valued dice and
// add point to player
func (p *Player) AddPointFromSixValuedDices() {
	for i := 0; i < len(p.Dices); {
		d := p.Dices[i]
		if d.Number == 6 {
			copy(p.Dices[i:], p.Dices[i+1:])
			p.Dices = p.Dices[:len(p.Dices)-1]
			p.Point++
		} else {
			i++
		}
	}
}

// RemoveOneValuedDices will remove one valued dice and return
// the removed dices
func (p *Player) RemoveOneValuedDices() []*dice.Dice {
	ones := []*dice.Dice{}
	for i := 0; i < len(p.Dices); {
		d := p.Dices[i]
		if d.Number == 1 {
			one, _ := p.RemoveDice(i)
			ones = append(ones, one)
		} else {
			i++
		}
	}
	return ones
}

// RollDices will roll the dices the player have
func (p *Player) RollDices() {
	for currIdx, d := range p.Dices {
		rollDice(&d)
		p.Dices[currIdx] = d
	}
}

func (p *Player) validateDice(d *dice.Dice) error {
	if err := validateDice(d); err != nil {
		if err == dice.ErrInvalidDiceNumber {
			return errInsertDiceFailed
		}
		return err
	}
	return nil
}

// AddDice will add dice to player, will return error if dice invalid
func (p *Player) AddDice(d *dice.Dice) error {
	if err := p.validateDice(d); err != nil {
		return err
	}
	p.Dices = append(p.Dices, *d)
	return nil
}

// RemoveDice will remove index i dice from player Dices and return the dice
// object at i.
func (p *Player) RemoveDice(i int) (*dice.Dice, error) {
	if i >= len(p.Dices) {
		return nil, ErrIndexOutOfBound
	}
	d := p.Dices[i]
	copy(p.Dices[i:], p.Dices[i+1:])
	p.Dices = p.Dices[:len(p.Dices)-1]
	return &d, nil
}

// AddDicesLater will store dice to be added later
func (p *Player) AddDicesLater(ds ...*dice.Dice) error {
	if ds == nil {
		return errInsertDiceFailed
	}
	for i, d := range ds {
		if err := p.validateDice(d); err != nil {
			er := fmt.Errorf("error at index %d: %s", i, err.Error())
			logger.L.Error(er.Error())
			return er
		}
		p.addLater = append(p.addLater, *d)
	}
	return nil
}

// CommitAddLaterDices will commit added later dice
func (p *Player) CommitAddLaterDices() {
	p.Dices = append(p.Dices, p.addLater...)
	p.addLater = []dice.Dice{}
}

func (p *Player) String() string {
	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf("Pemain #%%d (%d): ", p.Point))
	for i, d := range p.Dices {
		sb.WriteString(fmt.Sprintf("%d", d.Number))
		if i < len(p.Dices)-1 {
			sb.WriteString(",")
		}
	}
	if len(p.Dices) == 0 {
		sb.WriteString("_ (Berhenti bermain karena tidak memiliki dadu)")
	}
	return sb.String()
}
