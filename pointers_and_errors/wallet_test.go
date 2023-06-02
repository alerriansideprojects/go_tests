package pointers_and_errors

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {
	assertError := func(t testing.TB, got error, want error) {
		t.Helper()
		if got == nil {
			t.Fatal("did not receive an error but expected one")
		}

		if got != want {
			t.Errorf("\nGot: %q\nWant: %q", got, want)
		}
	}

	assertWalletBalance := func(t testing.TB, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()
		if got != want {
			t.Errorf("\nGot: %s\nWant: %s", got, want)
		}
	}

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))

		assertWalletBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		wallet.Withdraw(Bitcoin(10))

		assertWalletBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw insufficient funds", func(t *testing.T){
		startBalance := Bitcoin(20)
		wallet := Wallet{startBalance}

		err := wallet.Withdraw(Bitcoin(100))

		assertWalletBalance(t, wallet, startBalance)
		assertError(t, err, ErrInsufficientFunds)
	})
}

func ExampleWallet() {
	wallet := Wallet{}
	wallet.Deposit(Bitcoin(10))
	balance := wallet.Balance()
	fmt.Println(balance)
	// Output: 10 BTC
}

func BenchmarkWalletDeposit(b *testing.B){
	wallet := Wallet{}
	for i := 0; i < b.N; i++ {
		wallet.Deposit(Bitcoin(10))
	}
}

func BenchmarkWalletWithdraw(b *testing.B){
	wallet := Wallet{}
	for i := 0; i < b.N; i++ {
		wallet.Withdraw(Bitcoin(10))
	}
}