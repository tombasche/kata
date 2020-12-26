package shuffle

import (
	"reflect"
	"testing"
)

func TestShuffle(t *testing.T) {
	input := []int{1, 1, 2, 2}
	expected := []int{1, 2, 1, 2}
	n := 2
	r := Shuffle(input, n)

	if !reflect.DeepEqual(expected, r) {
		t.Fatalf("lists dont match. want %v, got %v", expected, r)
	}
}
