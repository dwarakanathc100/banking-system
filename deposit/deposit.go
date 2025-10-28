package deposit

import (
    "errors"
    "banking-system/useraccount"
)

func Deposit(u *useraccount.User, amount float64) error {
    if !u.IsAccountActive() {
        return errors.New("account is not active")
    }
    if amount <= 0 {
        return errors.New("deposit amount must be positive")
    }
    u.UpdateBalance(amount)
    return nil
}
