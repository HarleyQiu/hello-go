package pointers_errors

import (
	"testing"
)

func TestWallet(t *testing.T) {
	// 存款
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})
	// 提取
	t.Run("Withdraw with funds", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(10)
		assertBalance(t, wallet, Bitcoin(10))
		assertNoError(t, err)
	})
	// 提取不足的资金
	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(100))
		assertBalance(t, wallet, startingBalance)
		assertError(t, err, "无法提取，资金不足")
	})
}

func assertBalance(t *testing.T, wallet Wallet, want Bitcoin) {
	got := wallet.Balance()
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func assertNoError(t *testing.T, got error) {
	if got != nil {
		t.Fatal("got an error but didnt want one")
	}
}

func assertError(t *testing.T, got error, want string) {
	if got == nil {
		t.Fatal("没有收到错误，但想要一个。")
	}
	if got.Error() != want {
		t.Errorf("got '%s', want '%s'", got, want)
	}
}
