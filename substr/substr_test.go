package substr

import "testing"

func TestSubstr(t *testing.T) {
	tests := []struct {
		input  string
		sub    string
		length int
	}{
		{"abcabcbb", "abc", 3},
		{"bbbbb", "b", 1},
		{"pwwkew", "wke", 3},
	}
	for _, tt := range tests {
		rSub, rLen := Substr(tt.input)

		if rSub != tt.sub {
			t.Fatalf("substrings dont match. want=%q, got=%q", tt.sub, rSub)
		}
		if tt.length != rLen {
			t.Fatalf("substring has incorrect length. want %d, got %d", tt.length, rLen)
		}
	}

}

//bbbbb
