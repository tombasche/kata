package divide

import "testing"

func TestDivide(t *testing.T) {

	tests := []struct {
		dividend int
		divisor  int
		expected int
	}{
		{10, 3, 3},
		{7, -3, -2},
		{0, 1, 0},
		{1, 1, 1},
	}
	for _, tt := range tests {
		r := Divide(tt.dividend, tt.divisor)

		if r != tt.expected {
			t.Fatalf("wrong answer. want %d, got %d", tt.expected, r)
		}
	}

}
