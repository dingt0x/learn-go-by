package main

import (
    "testing"
)

func TestWallet(t *testing.T) {

    assertBalance := func(t testing.TB, wallet Wallet, want Bitcoin) {
        t.Helper()
        got := wallet.Balance()

        if got != want {
            t.Errorf("got %s want %s", got, want)
        }
    }
    assertError := func(t testing.TB, got error, want error) {
        t.Helper()
        if got == nil {
            t.Fatal("wanted an error but didn't get one")
        }
        if got != want {
            t.Errorf("got %q , want %q", got, want)
        }
    }

    assertNoErr := func(t testing.TB, got error) {
        t.Helper()

        if got != nil {
            t.Fatal("got an error but didn't want one")
        }
    }

    t.Run("deposit", func(t *testing.T) {
        wallet := Wallet{}
        wallet.Deposit(10)
        want := Bitcoin(10)

        // fmt.Printf("address of balance in test is %p \n", &wallet.balance)
        assertBalance(t, wallet, want)

    })

    t.Run("withdraw", func(t *testing.T) {
        wallet := Wallet{}
        wallet.balance = 20
        want := Bitcoin(10)

        err := wallet.Withdraw(10)
        assertNoErr(t, err)

        assertBalance(t, wallet, want)

    })

    t.Run("withdraw insufficient", func(t *testing.T) {

        stringBalance := Bitcoin(20)
        wallet := Wallet{balance: stringBalance}
        err := wallet.Withdraw(Bitcoin(100))

        assertError(t, err, ErrInsufficientFunds)
        assertBalance(t, wallet, stringBalance)

    })

}
