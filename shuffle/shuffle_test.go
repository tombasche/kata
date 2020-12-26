package shuffle

import (
	"reflect"
	"testing"
)

func TestShuffle(t *testing.T) {

	tests := []struct {
		input    []int
		expected []int
		n        int
	}{
		{[]int{1, 1, 2, 2}, []int{1, 2, 1, 2}, 2},
		{[]int{2, 5, 1, 3, 4, 7}, []int{2, 3, 5, 4, 1, 7}, 3},
		{[]int{1, 2, 3, 4, 4, 3, 2, 1}, []int{1, 4, 2, 3, 3, 2, 4, 1}, 4},
	}

	for _, tt := range tests {
		r := Shuffle(tt.input, tt.n)

		if !reflect.DeepEqual(tt.expected, r) {
			t.Fatalf("lists dont match. want %v, got %v", tt.expected, r)
		}
	}

}
