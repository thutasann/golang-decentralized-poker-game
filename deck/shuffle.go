package deck

import "math/rand"

// Shuffle Fn
func Shuffle(d Deck) Deck {
	for i := 0; i < len(d); i++ {
		r := rand.Intn(i + 1)

		if r != i {
			d[i], d[r] = d[r], d[i]
		}
	}

	return d
}
