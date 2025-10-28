package regression

import (
	"banking-system/deposit"
	"banking-system/emi"
	"banking-system/loan"
	"banking-system/useraccount"
	"banking-system/withdraw"
	"testing"
)

// Regression test to ensure the system still behaves as expected after changes
func TestFullBankingFlowRegression(t *testing.T) {
	user := useraccount.CreateUser(1, "Dwarkanath", "ACC900", "dwarka@bank.com", "9999999999", 1000)

	if err := deposit.Deposit(user, 200); err != nil {
		t.Errorf("Deposit failed: %v", err)
	}

	if err := withdraw.Withdraw(user, 100); err != nil {
		t.Errorf("Withdraw failed: %v", err)
	}

	if err := loan.ApplyLoan(user, 1000); err != nil {
		t.Errorf("Loan failed: %v", err)
	}

	if err := emi.PayEMI(user, 400); err != nil {
		t.Errorf("EMI payment failed: %v", err)
	}

	expected := 1700.0
	if user.Balance != expected {
		t.Errorf("Expected balance %.2f, got %.2f", expected, user.Balance)
	}
}
