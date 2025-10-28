package integration

import (
	"banking-system/deposit"
	"banking-system/statement"
	"banking-system/useraccount"
	"banking-system/withdraw"
	"testing"
	//"banking-system/notification"
	"banking-system/emi"
	"banking-system/loan"
)

type MockNotifier struct {
	messages []string
}

func (m *MockNotifier) Send(msg string) string {
	m.messages = append(m.messages, msg)
	return "Mock notification sent"
}

func TestBankingIntegration(t *testing.T) {
	user := useraccount.CreateUser(1, "Dwarkanath", "ACC999", "dwarka@bank.com", "1234567890", 1000)

	if err := deposit.Deposit(user, 500); err != nil {
		t.Fatalf("deposit failed: %v", err)
	}

	if err := withdraw.Withdraw(user, 300); err != nil {
		t.Fatalf("withdraw failed: %v", err)
	}

	if err := loan.ApplyLoan(user, 1000); err != nil {
		t.Fatalf("loan failed: %v", err)
	}

	if err := emi.PayEMI(user, 400); err != nil {
		t.Fatalf("emi payment failed: %v", err)
	}

	mock := &MockNotifier{}
	msg := statement.GenerateStatement(user)
	mock.Send(msg)

	if len(mock.messages) == 0 {
		t.Error("Expected notification to be sent")
	}
}
