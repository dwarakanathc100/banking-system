package loan

import (
	"banking-system/useraccount"
	"errors"
)

func ApplyLoan(u *useraccount.User, amount float64) error {
	if amount <= 0 {
		return errors.New("invalid loan amount")
	}
	//check background
	//credit score
	//any existing loans
	u.UpdateBalance(amount)
	return nil
}
