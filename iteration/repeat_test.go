package iteration

import (
	"testing"
)

func Repeat(character string, numberOfRepeat int) string {
	var repeated string
	for i := 0; i < numberOfRepeat; i++ {
		repeated += character
	}
	return repeated
}

func TestRepeat(t *testing.T) {
	t.Run("repeating a five times", func(t *testing.T) {
		input := Repeat("a", 5)
		expected := "aaaaa"

		if input != expected {
			t.Errorf("got %v, want %v", input, expected)
		}
	})

	t.Run("repeating mo six times", func(t *testing.T) {
		input := Repeat("mo", 6)
		expected := "momomomomomo"

		if input != expected {
			t.Errorf("got %v, want %v", input, expected)
		}
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
