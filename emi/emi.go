package emi

import (
    "errors"
    "banking-system/useraccount"
)

func PayEMI(u *useraccount.User, amount float64) error {
    if amount <= 0 {
        return errors.New("invalid EMI amount")
    }
    if u.Balance < amount {
        return errors.New("insufficient funds for EMI")
    }
    u.UpdateBalance(-amount)
    return nil
}
