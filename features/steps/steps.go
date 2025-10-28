package features

import (
	"fmt"
	"strings"

	"banking-system/deposit"
	"banking-system/emi"
	"banking-system/loan"
	"banking-system/useraccount"
	"banking-system/withdraw"

	"github.com/cucumber/godog"
)

type world struct {
	user    *useraccount.User
	lastErr error
}

func (w *world) iAmANewCustomerNamed(name string) error {
	// store name in user struct later when opening account
	w.user = &useraccount.User{Name: name}
	return nil
}

func (w *world) iOpenANewAccountWithAccountNumberAndInitialBalance(acc string, balance int) error {
	// create user with given account and balance
	w.user = useraccount.CreateUser(1, w.user.Name, acc, w.user.Email, w.user.Phone, float64(balance))
	return nil
}

func (w *world) myAccountShouldBeActive() error {
	if w.user == nil || !w.user.IsAccountActive() {
		return fmt.Errorf("expected account to be active")
	}
	return nil
}

func (w *world) myAccountBalanceShouldBe(expected int) error {
	if w.user == nil {
		return fmt.Errorf("no user in context")
	}
	if int(w.user.Balance) != expected {
		return fmt.Errorf("expected balance %d, got %v", expected, w.user.Balance)
	}
	return nil
}

//func (w *world) iHaveAnActiveAccountWithAccountNumberAndBalance(acc string, balance int) error {
//    w.user = useraccount.CreateUser(2, "acctuser", acc, "acct@example.com", "000", float64(balance))
//    return nil
//}

func (w *world) iDepositAmountIntoMyAccount(amount int) error {
	w.lastErr = deposit.Deposit(w.user, float64(amount))
	return nil
}

func (w *world) iTryToDepositIntoMyAccount(amount int) error {
	w.lastErr = deposit.Deposit(w.user, float64(amount))
	return nil
}

func (w *world) iShouldSeeAnError(expected string) error {
	if w.lastErr == nil {
		return fmt.Errorf("expected error '%s' but got nil", expected)
	}
	if !strings.Contains(w.lastErr.Error(), expected) {
		return fmt.Errorf("expected error to contain '%s', got '%s'", expected, w.lastErr.Error())
	}
	return nil
}

func (w *world) myAccountBalanceShouldBecome(expected int) error {
	return w.myAccountBalanceShouldBe(expected)
}

func (w *world) iHaveAnInactiveAccountWithAccountNumber(acc string) error {
	w.user = useraccount.CreateUser(3, "inactive", acc, "inactive@example.com", "111", 0)
	w.user.CloseAccount()
	return nil
}

func (w *world) iHaveAnActiveAccountWithAccountNumberAndBalance(acc string, balance int) error {
	w.user = useraccount.CreateUser(4, "active", acc, "active@example.com", "222", float64(balance))
	return nil
}

func (w *world) iWithdrawAmount(amount int) error {
	w.lastErr = withdraw.Withdraw(w.user, float64(amount))
	return nil
}

func (w *world) iTryToWithdraw(amount int) error {
	w.lastErr = withdraw.Withdraw(w.user, float64(amount))
	return nil
}

func (w *world) iCreateANewAccountWithAccountNumberAndInitialBalance(acc string, balance int) error {
	w.user = useraccount.CreateUser(10, "flow", acc, "flow@example.com", "333", float64(balance))
	return nil
}

func (w *world) iApplyForALoanOf(amount int) error {
	w.lastErr = loan.ApplyLoan(w.user, float64(amount))
	return nil
}

func (w *world) iPayEMIOf(amount int) error {
	w.lastErr = emi.PayEMI(w.user, float64(amount))
	return nil
}

func (w *world) myFinalAccountBalanceShouldBe(expected int) error {
	if int(w.user.Balance) != expected {
		return fmt.Errorf("expected final balance %d, got %v", expected, w.user.Balance)
	}
	return nil
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	w := &world{}
	ctx.Step(`^I am a new customer named "([^\"]*)"$`, w.iAmANewCustomerNamed)
	ctx.Step(`^I open a new account with account number "([^\"]*)" and initial balance (\d+)$`, w.iOpenANewAccountWithAccountNumberAndInitialBalance)
	ctx.Step(`^my account should be active$`, w.myAccountShouldBeActive)
	ctx.Step(`^my account balance should be (\d+)$`, w.myAccountBalanceShouldBe)

	ctx.Step(`^I have an active account with account number "([^\"]*)" and balance (\d+)$`, w.iHaveAnActiveAccountWithAccountNumberAndBalance)
	ctx.Step(`^I deposit (\d+) into my account$`, w.iDepositAmountIntoMyAccount)
	ctx.Step(`^I try to deposit (\d+) into my account$`, w.iTryToDepositIntoMyAccount)
	ctx.Step(`^I should see an error "([^\"]*)"$`, w.iShouldSeeAnError)
	ctx.Step(`^my account balance should become (\d+)$`, w.myAccountBalanceShouldBecome)
	ctx.Step(`^I have an inactive account with account number "([^\"]*)"$`, w.iHaveAnInactiveAccountWithAccountNumber)

	ctx.Step(`^I withdraw (\d+)$`, w.iWithdrawAmount)
	ctx.Step(`^I try to withdraw (\d+)$`, w.iTryToWithdraw)

	ctx.Step(`^I create a new account with account number "([^\"]*)" and initial balance (\d+)$`, w.iCreateANewAccountWithAccountNumberAndInitialBalance)
	ctx.Step(`^I apply for a loan of (\d+)$`, w.iApplyForALoanOf)
	ctx.Step(`^I pay EMI of (\d+)$`, w.iPayEMIOf)
	ctx.Step(`^my final account balance should be (\d+)$`, w.myFinalAccountBalanceShouldBe)
	ctx.Step(`^I close my account$`, iCloseMyAccount)
	ctx.Step(`^I deposit (\d+)$`, iDeposit)
	ctx.Step(`^I have an active account with account number "([^"]*)"$`, w.iHaveAnActiveAccountWithAccountNumber)
	ctx.Step(`^my account should be inactive$`, myAccountShouldBeInactive)
	ctx.Step(`^my account should remain active$`, myAccountShouldRemainActive)
}

func iAmANewCustomerNamed(arg1 string) error {
	return godog.ErrPending
}

func iApplyForALoanOf(arg1 int) error {
	return godog.ErrPending
}

func iCloseMyAccount() error {
	return godog.ErrPending
}

func iCreateANewAccountWithAccountNumberAndInitialBalance(arg1 string, arg2 int) error {
	return godog.ErrPending
}

func iDeposit(arg1 int) error {
	return godog.ErrPending
}

func iDepositIntoMyAccount(arg1 int) error {
	return godog.ErrPending
}

func (w *world) iHaveAnActiveAccountWithAccountNumber(arg1 string) error {
	//w.user.CloseAccount()
	return godog.ErrPending
}

func iHaveAnActiveAccountWithAccountNumberAndBalance(arg1 string, arg2 int) error {
	return godog.ErrPending
}

func iHaveAnInactiveAccountWithAccountNumber(arg1 string) error {
	return godog.ErrPending
}

func iOpenANewAccountWithAccountNumberAndInitialBalance(arg1 string, arg2 int) error {
	return godog.ErrPending
}

func iPayEMIOf(arg1 int) error {
	return godog.ErrPending
}

func iShouldSeeAnError(arg1 string) error {
	return godog.ErrPending
}

func iTryToDepositIntoMyAccount(arg1 int) error {
	return godog.ErrPending
}

func iTryToWithdraw(arg1 int) error {
	return godog.ErrPending
}

func iWithdraw(arg1 int) error {
	return godog.ErrPending
}

func myAccountBalanceShouldBe(arg1 int) error {
	return godog.ErrPending
}

func myAccountBalanceShouldBecome(arg1 int) error {
	return godog.ErrPending
}

func myAccountShouldBeActive() error {
	return godog.ErrPending
}

func myAccountShouldBeInactive() error {
	return godog.ErrPending
}

func myAccountShouldRemainActive() error {
	return godog.ErrPending
}

func myFinalAccountBalanceShouldBe(arg1 int) error {
	return godog.ErrPending
}

func iAmANewCusgodogRunFeaturesdepositfeaturetomerNamed(arg1 string) error {
	return godog.ErrPending
}
