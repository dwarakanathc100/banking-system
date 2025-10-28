package withdraw

import (
    "errors"
    "banking-system/useraccount"
)

func Withdraw(u *useraccount.User, amount float64) error {
    if !u.IsAccountActive() {
        return errors.New("account is inactive")
    }
    if amount <= 0 {
        return errors.New("withdrawal amount must be positive")
    }
    if u.Balance < amount {
        return errors.New("insufficient funds")
    }
    u.UpdateBalance(-amount)
    return nil
}
