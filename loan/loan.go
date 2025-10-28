package loan

import (
    "errors"
    "banking-system/useraccount"
)

func ApplyLoan(u *useraccount.User, amount float64) error {
    if amount <= 0 {
        return errors.New("invalid loan amount")
    }
    u.UpdateBalance(amount)
    return nil
}
