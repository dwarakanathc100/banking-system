package withdraw

import (
    "testing"
    "banking-system/useraccount"
)

func TestWithdraw(t *testing.T) {
    tests := []struct {
        name    string
        balance float64
        amount  float64
        wantErr bool
    }{
        {"SufficientFunds", 500, 200, false},
        {"InsufficientFunds", 100, 200, true},
        {"NegativeAmount", 100, -10, true},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            u := useraccount.CreateUser(1, "John", "ACC202", "john@example.com", "123", tt.balance)
            err := Withdraw(u, tt.amount)
            if (err != nil) != tt.wantErr {
                t.Errorf("expected error: %v, got: %v", tt.wantErr, err)
            }
        })
    }
}
