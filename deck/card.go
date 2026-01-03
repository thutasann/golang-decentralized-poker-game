package deck

import (
	"fmt"
	"strconv"
)

type Card struct {
	suit  Suit
	value int
}

// Initialize a new card
func NewCard(s Suit, v int) Card {
	if v > 13 {
		panic("the value of the card cannot be higher than 13")
	}

	return Card{
		suit:  s,
		value: v,
	}
}

func (c Card) String() string {
	value := strconv.Itoa(c.value)
	if c.value == 1 {
		value = "ACE"
	}
	return fmt.Sprintf("%s of %s %s", value, c.suit, suitToUnicode(c.suit))
}
