package arrays

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		want := 15

		assertCorrectMessage(t, got, want)
	})

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		got := Sum(numbers)
		want := 6

		assertCorrectMessage(t, got, want)
	})
}

func TestSumAll(t *testing.T) {
	t.Run("collection of 2 arrays", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		numbers2 := []int{1, 2, 3, 4, 5}
		got := SumAll(numbers, numbers2)
		want := []int{15, 15}

		assertSumArray(t, got, want)
	})
}

func TestSumAllTails(t *testing.T) {
	t.Run("sum of tails", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3, 4}, []int{0, 9})
		want := []int{9, 9}

		assertSumArray(t, got, want)
	})

	t.Run("sum of empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{1, 2, 3, 4})
		want := []int{0, 9}

		assertSumArray(t, got, want)
	})
}

func assertCorrectMessage(t *testing.T, got, want int) {
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func assertSumArray(t *testing.T, got, want []int) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
