package wealth

import "testing"

func TestWealthiest(t *testing.T) {
	input := [][]int{{1, 2, 3}, {3, 2, 1}}
	expected := 6
	r := Wealthiest(input)
	if r != expected {
		t.Fatalf("wrong total wealth. want %d got %d", expected, r)
	}
}
