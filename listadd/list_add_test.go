package listadd

import (
	"reflect"
	"testing"
)

func TestListAdd(t *testing.T) {

	tests := []struct {
		a        Node
		b        Node
		expected *Node
	}{
		{
			Node{Value: 2, Next: &Node{Value: 4, Next: &Node{Value: 3, Next: nil}}},
			Node{Value: 5, Next: &Node{Value: 6, Next: &Node{Value: 4, Next: nil}}},
			&Node{Value: 7, Next: &Node{Value: 0, Next: &Node{Value: 8, Next: nil}}},
		},
		{
			Node{Value: 1, Next: &Node{Value: 2, Next: &Node{Value: 3, Next: nil}}},
			Node{Value: 1, Next: &Node{Value: 2, Next: nil}},
			&Node{Value: 2, Next: &Node{Value: 4, Next: &Node{Value: 3, Next: nil}}},
		},
		{
			Node{Value: 9, Next: &Node{Value: 9, Next: nil}},
			Node{Value: 9, Next: &Node{Value: 9, Next: &Node{Value: 9, Next: nil}}},
			&Node{Value: 8, Next: &Node{Value: 9, Next: &Node{Value: 0, Next: &Node{Value: 1, Next: nil}}}},
		},
	}
	for _, tt := range tests {
		r := ListAdd(&tt.a, &tt.b)
		if !reflect.DeepEqual(tt.expected.Values(), r.Values()) {
			t.Fatalf("lists aren't the same. want=%v, got=%v", tt.expected.Values(), r.Values())
		}
	}

}
