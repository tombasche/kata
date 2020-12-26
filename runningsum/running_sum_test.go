package runningsum

import (
	"reflect"
	"testing"
)

func TestRunningSum(t *testing.T) {

	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{1, 2, 3, 4}, []int{1, 3, 6, 10}},
		{[]int{1, 1, 1, 1, 1}, []int{1, 2, 3, 4, 5}},
		{[]int{3, 1, 2, 10, 1}, []int{3, 4, 6, 16, 17}},
	}

	for _, tt := range tests {
		r := RunningSum(tt.input)

		if !reflect.DeepEqual(r, tt.expected) {
			t.Fatalf("results dont match. want %v got %v", tt.expected, r)
		}
	}
}
