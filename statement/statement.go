package statement

import (
    "fmt"
    "time"
    "banking-system/useraccount"
)

func GenerateStatement(u *useraccount.User) string {
    return fmt.Sprintf("Account: %s | Name: %s | Balance: %.2f | Active: %t | Date: %s",
        u.AccountNo, u.Name, u.Balance, u.IsActive, time.Now().Format("2006-01-02 15:04:05"))
}
