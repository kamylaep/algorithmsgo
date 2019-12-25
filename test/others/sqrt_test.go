package others

import (
	"kep/algorithms-go/others"
	"testing"
)

type Sqrt struct {
	X        int
	Expected float64
}

func TestSqrtWithSuccess(t *testing.T) {
	sqrts := []Sqrt{
		Sqrt{X: 2, Expected: 1.41},
		Sqrt{X: 3, Expected: 1.73},
		Sqrt{X: 10, Expected: 3.16},
		Sqrt{X: 99, Expected: 9.95},
	}

	for _, sq := range sqrts {
		sqrt := others.Sqrt(float64(sq.X))

		if sq.Expected != sqrt {
			t.Fatalf("Expected [%f] but got [%f]", sq.Expected, sqrt)
		}
	}
}
