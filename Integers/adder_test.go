package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	t.Run("TestAdder", func(t *testing.T) {
		sum := Add(2, 2)
		expected := 4
		assertCorrectMessage(t, sum, expected)
	})
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}

func assertCorrectMessage(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
