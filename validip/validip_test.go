package validip

import "testing"

func TestValidIP(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"192.168.1.1", Ipv4},
		{"2001:0db8:85a3:0:0:8A2E:0370:7334", Ipv6},
		{"2001:0db8:85a3:0:0:8a2e:0370:7334", Ipv6},
		{"200121:0db8:85a3:0:0:8a2e:0370:73344", Neither},
		{"bananas", Neither},
		{"256.256.256.256", Neither},
		{"192.168.1.00", Neither},
		{"192.168@1.1", Neither},
		{"2001:0db8:85a3::8A2E:037j:7334", Neither},
		{"1e1.4.5.6", Neither},
	}
	for _, tt := range tests {
		r := ValidIP(tt.input)
		if r != tt.expected {
			t.Fatalf("didnt return correct result for %s. want %v got %v", tt.input, tt.expected, r)
		}
	}
}
