package useraccount

import "time"

type User struct {
    ID              int
    Name            string
    AccountNo       string
    Email           string
    Phone           string
    Balance         float64
    AccountOpenDate time.Time
    AccountCloseDate *time.Time
    IsActive        bool
}

func CreateUser(id int, name, accountNo, email, phone string, balance float64) *User {
    return &User{
        ID:              id,
        Name:            name,
        AccountNo:       accountNo,
        Email:           email,
        Phone:           phone,
        Balance:         balance,
        AccountOpenDate: time.Now(),
        IsActive:        true,
    }
}

func (u *User) UpdateBalance(amount float64) {
    u.Balance += amount
}

func (u *User) CloseAccount() {
    now := time.Now()
    u.AccountCloseDate = &now
    u.IsActive = false
}

func (u *User) IsAccountActive() bool {
    return u.IsActive
}
