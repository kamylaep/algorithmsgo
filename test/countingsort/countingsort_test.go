package bubblesort

import (
	"reflect"
	"testing"

	"github.com/kamylaep/algorithmsgo/countingsort"
)

type testData struct {
	unordered []int
	ordered   []int
}

func TestBubbleSortWithSuccess(t *testing.T) {
	data := []testData{
		testData{unordered: []int{8, 5, 6, 1, 2, 7, 9, 3, 0, 4}, ordered: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}},
		testData{unordered: []int{10, 5, 2, 3}, ordered: []int{2, 3, 5, 10}},
		testData{unordered: []int{}, ordered: []int{}},
		testData{unordered: []int{5}, ordered: []int{5}},
		testData{unordered: []int{2, 1}, ordered: []int{1, 2}},
		testData{unordered: []int{1, 2, 3}, ordered: []int{1, 2, 3}},
		testData{unordered: nil, ordered: nil},
	}

	for _, d := range data {
		countingsort.Sort(d.unordered)
		isOrdered := reflect.DeepEqual(d.ordered, d.unordered)
		if !isOrdered {
			t.Fatalf("Expected true but got false [%v, %v]", d.ordered, d.unordered)
		}
	}
}
