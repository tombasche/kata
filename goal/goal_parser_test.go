package goal

import "testing"

func TestGoalParser(t *testing.T) {
	input := "G()()()(al)"
	expected := "Goooal"

	r := GoalParser(input)
	if r != expected {
		t.Fatalf("didnt parse correctly. want %s, got %s", expected, r)
	}
}
