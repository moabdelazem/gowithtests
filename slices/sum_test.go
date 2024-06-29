package slices

import "testing"

func Sum(numbers []int) int {
	var sum int
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func TestSum(t *testing.T) {

	t.Run("sum slice of 5 elements", func(t *testing.T) {
		testNumbers := []int{1, 2, 3, 4, 5}
		input := Sum(testNumbers)
		output := 15

		if input != output {
			t.Errorf("got %v want %v", input, output)
		}
	})
}
