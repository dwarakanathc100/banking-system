package useraccount

import (
	"testing"
)

func TestUserFunctions(t *testing.T) {
	t.Run("CreateUser_ShouldInitializeCorrectly", func(t *testing.T) {
		u := CreateUser(1, "Dwarkanath", "ACC123", "dwark@example.com", "9876543210", 1000)
		if u.Name != "Dwarkanath" {
			t.Errorf("Expected name Dwarkanath, got %s", u.Name)
		}
		if !u.IsActive {
			t.Error("Expected account to be active")
		}
	})

	t.Run("UpdateBalance_ShouldIncreaseBalance", func(t *testing.T) {
		u := CreateUser(2, "Gaurak", "ACC456", "gaurak@example.com", "9998887777", 100)
		u.UpdateBalance(50)
		if u.Balance != 150 {
			t.Errorf("Expected 150, got %f", u.Balance)
		}
	})

	t.Run("CloseAccount_ShouldDeactivateAccount", func(t *testing.T) {
		u := CreateUser(3, "TestUser", "ACC789", "test@example.com", "1234567890", 500)
		u.CloseAccount()
		if u.IsActive {
			t.Error("Expected account to be inactive")
		}
		if u.AccountCloseDate == nil {
			t.Error("Expected closing date to be set")
		}
	})
}

func BenchmarkCreateUser(b *testing.B) {
	for i := 0; i < b.N; i++ { //10000
		CreateUser(1, "User", "ACC001", "user@example.com", "12345", 1000)
	}
}
