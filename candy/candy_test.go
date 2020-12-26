package candy

import (
	"reflect"
	"testing"
)

func TestDistributeCandies(t *testing.T) {
	input := []int{2, 3, 5, 1, 3}
	extras := 3

	expected := []bool{true, true, true, false, true}
	r := DistributeCandies(input, extras)

	if !reflect.DeepEqual(r, expected) {
		t.Fatalf("arrays don't match. want %v, got %v", expected, r)
	}
}
