package pointers

import "testing"

func TestWallet(t *testing.T) {
	assertTheBalance := func(t testing.TB, wallet Wallet, want BitCoin) {
		t.Helper()
		got := wallet.Balance()
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	}

	assertErr := func(t testing.TB, err error, want string) {
		t.Helper()
		if err == nil {
			t.Error("wanted an error but didn't get one")
		}
		if err.Error() != want {
			t.Errorf("got %q, want %q", err, want)
		}
	}

	t.Run("check for returning the balance", func(t *testing.T) {
		testWallet := Wallet{}
		want := BitCoin(0)

		assertTheBalance(t, testWallet, want)

	})
	t.Run("deposit 10 and returning the balance", func(t *testing.T) {
		testWallet := Wallet{}
		testWallet.Deposit(10)

		want := BitCoin(10)

		assertTheBalance(t, testWallet, want)
	})
	t.Run("withdraw 250 and returning the balance", func(t *testing.T) {
		testWallet := Wallet{
			WalletBalance: BitCoin(1000),
		}
		testWallet.Withdraw(250)
		want := BitCoin(750)

		assertTheBalance(t, testWallet, want)
	})
	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startingBalance := BitCoin(20)
		testWallet := Wallet{startingBalance}
		err := testWallet.Withdraw(BitCoin(100))

		assertErr(t, err, "cannot withdraw, insufficient funds")
		assertTheBalance(t, testWallet, startingBalance)
	})
}
