package deposit

import (
	"banking-system/useraccount"
	"testing"
)

func TestDeposit(t *testing.T) {
	t.Run("ValidDeposit", func(t *testing.T) {
		u := useraccount.CreateUser(1, "John", "ACC100", "john@example.com", "111222333", 100)
		err := Deposit(u, 50)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if u.Balance != 150 {
			t.Errorf("Expected 150, got %f", u.Balance)
		}
	})

	t.Run("NegativeDeposit", func(t *testing.T) {
		u := useraccount.CreateUser(2, "Sam", "ACC101", "sam@example.com", "444555666", 100)
		err := Deposit(u, -100)
		if err == nil {
			t.Error("expected error for negative deposit")
		}
	})
}
