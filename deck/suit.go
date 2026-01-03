package deck

type Suit int

// Implement the fmt.Stringer interface to represent itself as a string
func (s Suit) String() string {
	switch s {
	case Spades:
		return "SPADES"
	case Harts:
		return "HARTS"
	case Diamonds:
		return "DIAMONDS"
	case Clubs:
		return "CLUBS"
	default:
		panic("Invalid Card suit")
	}
}

const (
	Spades Suit = iota
	Harts
	Diamonds
	Clubs
)
