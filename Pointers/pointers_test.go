package pointers

import "testing"

func TestWallet(t *testing.T) {
	t.Run("wallet deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))

		got := wallet.balance
		want := Bitcoin(10)
		assertBalance(t, got, want)
	})

	t.Run("wallet withdrawal", func(t *testing.T) {
		wallet := Wallet{balance: 10}

		wallet.Withdraw(Bitcoin(5))

		got := wallet.balance
		want := Bitcoin(5)
		assertBalance(t, got, want)
	})

	t.Run("wallet withdrawal with insufficient funds", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(10)}

		err := wallet.Withdraw(Bitcoin(15))

		if err != nil {
			t.Errorf("expected an error but didn't get one")
		}
	})
}

func assertBalance(t *testing.T, got, want Bitcoin) {
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
