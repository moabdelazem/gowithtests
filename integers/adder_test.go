package integers

import "testing"

// Add takes two int and return the sum of them
func Add(x, y int) int {
	return x + y
}

func TestAdder(t *testing.T) {
	t.Run("test adding 2 + 2", func(t *testing.T) {
		input := Add(2, 2)
		expected := 4

		if input != expected {
			t.Errorf("got %v want %v", input, expected)
		}
	})
}
