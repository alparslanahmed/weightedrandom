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

	var rng, _ = rand.Int(rand.Reader, total)
	var randomNumber = rng.Int64()
	var countWeight int64 = 0
	var target P

	for _, element := range elements {
		var weight = element.GetWeight()

		if weight < 0 {
			continue
		}

		countWeight += weight
		if randomNumber < countWeight {
			target = element
			break
		}
	}

	return target, nil
}
