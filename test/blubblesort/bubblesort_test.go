package bubblesort

import (
	"reflect"
	"testing"

	bubblesort "github.com/kamylaep/algorithmsgo/blubblesort"
)

type testData struct {
	unordered []int
	ordered   []int
}

func TestBubbleSortWithSuccess(t *testing.T) {
	data := []testData{
		testData{unordered: []int{5, 9, 10, 4, 3, 45, 1, 8, 20, 11}, ordered: []int{1, 3, 4, 5, 8, 9, 10, 11, 20, 45}},
		testData{unordered: []int{8, 15, 0, -1, 5, -6, 13, -5, 44, 3}, ordered: []int{-6, -5, -1, 0, 3, 5, 8, 13, 15, 44}},
		testData{unordered: []int{8, 5, 6, 1, 2, 7, 9, 3, 0, 4}, ordered: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}},
		testData{unordered: []int{10, 5, 2, 3}, ordered: []int{2, 3, 5, 10}},
		testData{unordered: []int{}, ordered: []int{}},
		testData{unordered: []int{5}, ordered: []int{5}},
		testData{unordered: []int{2, 1}, ordered: []int{1, 2}},
		testData{unordered: []int{1, 2, 3}, ordered: []int{1, 2, 3}},
		testData{unordered: nil, ordered: nil},
	}

	for _, d := range data {
		bubblesort.Sort(d.unordered)
		isOrdered := reflect.DeepEqual(d.ordered, d.unordered)
		if !isOrdered {
			t.Fatalf("Expected true but got false [%v, %v]", d.ordered, d.unordered)
		}
	}
}
