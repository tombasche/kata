package atoi

import "testing"

func TestAtoi(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"42", 42},
	}
	for _, tt := range tests {
		r := Atoi(tt.input)

		if r != tt.expected {
			t.Fatalf("numbers dont match. want %d, got %d", tt.expected, r)
		}
	}
}
