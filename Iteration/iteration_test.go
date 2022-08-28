package main

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("should repeat 5 times", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"

		assertCorrectMessage(t, repeated, expected)
	})

	t.Run("should repeat 8 times", func(t *testing.T) {
		repeated := Repeat("ba", 8)
		expected := "babababababababa"

		assertCorrectMessage(t, repeated, expected)
	})
}

func ExampleRepeat() {
	repeated := Repeat("a", 5)
	fmt.Println(repeated)
	// Outputs "aaaaa"
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func assertCorrectMessage(t *testing.T, repeated, expected string) {
	if repeated != expected {
		t.Errorf("got %s want %s", repeated, expected)
	}
}
