package arrays

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	// numbers := [5]int{1, 2, 3, 4, 5}

	// got := Sum(numbers)
	// want := 15

	// // %v prints the default format, which works well for arrays
	// if got != want {
	// 	t.Errorf("got %d want %d, was given %v", got, want, numbers)
	// }

	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d, given %v", got, want, numbers)
		}
	})

	// Run `go test -cover` to determine test coverage
	// As shown... the following subtest is not needed
	// t.Run("collection of any size", func(t *testing.T) {
	// 	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// 	got := Sum(numbers)
	// 	want := 55

	// 	if got != want {
	// 		t.Errorf("got %d want %d, given %v", got, want, numbers)
	// 	}
	// })
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2, 3}, []int{4, 5, 6})
	want := []int{6, 15}

	// Go does not let you use equality operators with slices
	// Be careful, reflect.DeepEqual is not "type safe"
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSumAllTails(t *testing.T) {
	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("make the sums of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		checkSums(t, got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}

		checkSums(t, got, want)
	})
}
