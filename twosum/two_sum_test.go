package twosum

import (
	"math/rand"
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		input    []int
		target   int
		expected []int
	}{
		{
			[]int{2, 7, 11, 15},
			9,
			[]int{0, 1},
		},
		{
			[]int{3, 2, 4},
			6,
			[]int{1, 2},
		},
		{
			[]int{3, 3},
			6,
			[]int{0, 1},
		},
	}
	for _, tt := range tests {
		r := TwoSum(tt.input, tt.target)

		if len(r) != len(tt.expected) {
			t.Fatalf("resultant arrays have different lengths. want=%v, got=%v", tt.expected, r)
		}

		if len(r) == 0 {
			t.Fatalf("Result is empty")
		}

		if !reflect.DeepEqual(r, tt.expected) {
			t.Fatalf("result of TwoSum wrong. want=%v, got=%v", tt.expected, r)
		}
	}
}

func BenchmarkTwoSum(b *testing.B) {
	input := rand.Perm(10000000)
	_ = TwoSum(input, 479941)
}
