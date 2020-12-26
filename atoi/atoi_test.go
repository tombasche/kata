package atoi

import "testing"

func TestAtoi(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"42", 42},
		{"-42", -42},
		{"    -23", -23},
		{"    ", 0},
		{"4193 with words", 4193},
		{"words and 987", 0},
		{"-91283472332", -2147483648},
		{"91283472332", 2147483647},
	}
	for _, tt := range tests {
		r := Atoi(tt.input)

		if r != tt.expected {
			t.Fatalf("numbers dont match. want %d, got %d", tt.expected, r)
		}
	}
}
