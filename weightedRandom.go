package weightedRandom

import (
	"crypto/rand"
	"math/big"
)

// WeightedRandom returns a random element from a slice of P with CSPRNG.
// P must have GetWeight() int64 method
func WeightedRandom[P interface{ GetWeight() int64 }](elements []P) (P, error) {
	var total = big.NewInt(0)

	for _, element := range elements {
		total.Add(total, big.NewInt(element.GetWeight()))
	}

	var shuffler []P

	for _, element := range elements {
		for j := int64(0); j < element.GetWeight(); j++ {
			shuffler = append(shuffler, element)
		}
	}

	var rnd, _ = rand.Int(rand.Reader, total)
	return shuffler[rnd.Int64()], nil
}
