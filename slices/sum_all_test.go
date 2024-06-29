package slices

import (
	"reflect"
	"testing"
)

func SumAll(numbersToSum ...[]int) []int {
	var totalSums []int
	for _, numbers := range numbersToSum {
		totalSums = append(totalSums, Sum(numbers))
	}
	return totalSums
}

func SumAllTails(numbersToSum ...[]int) []int {
	var totalSums []int
	for _, numbers := range numbersToSum {

		if len(numbers) == 0 {
			totalSums = append(totalSums, 0)
		} else {
			tailNumbers := numbers[1:]
			totalSums = append(totalSums, Sum(tailNumbers))
		}

	}
	return totalSums
}

func TestSumAll(t *testing.T) {
	t.Run("sum collection of two slices", func(t *testing.T) {
		input := SumAll([]int{1, 2}, []int{0, 9})
		output := []int{3, 9}

		if !reflect.DeepEqual(input, output) {
			t.Errorf("got %q want %q", input, output)
		}
	})
}

func TestSumTails(t *testing.T) {
	checkSums := func(t testing.TB, input, output []int) {
		t.Helper()
		if !reflect.DeepEqual(input, output) {
			t.Errorf("got %v want %v", input, output)
		}
	}

	t.Run("make the sums of some slices", func(t *testing.T) {
		input := SumAllTails([]int{1, 2}, []int{0, 9})
		output := []int{2, 9}

		checkSums(t, input, output)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		input := SumAllTails([]int{}, []int{3, 4, 5})
		output := []int{0, 9}

		checkSums(t, input, output)
	})
}
