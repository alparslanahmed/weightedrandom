package weightedRandom

import (
	"testing"
)

// ## Example Data

type Elem struct {
	ID     int
	Weight int64
}

func (e Elem) GetWeight() int64 { return e.Weight }

var DummyArray = []Elem{
	{ID: 100, Weight: 1},
	{ID: 200, Weight: 2},
	{ID: 300, Weight: 3},
	{ID: 400, Weight: 4},
	{ID: 500, Weight: 5},
	{ID: 600, Weight: 6},
	{ID: 700, Weight: 7},
	{ID: 800, Weight: 8},
	{ID: 900, Weight: 9},
	{ID: 1000, Weight: 10},
}

// ## EXAMPLES
// TODO: Write some examples

// ## TESTS

func TestWeightedRandom(t *testing.T) {
	_, err := WeightedRandom(DummyArray)
	if err != nil {
		t.Fatal(err)
	}
}

// ## BENCHMARKS
func BenchmarkWeightedRandom(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := WeightedRandom(DummyArray)
		if err != nil {
			b.Failed()
		}
	}
}

func BenchmarkWeightedRandomParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, err := WeightedRandom(DummyArray)
			if err != nil {
				b.Failed()
			}
		}
	})
}
