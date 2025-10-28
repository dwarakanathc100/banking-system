Feature: Withdraw Money
  As a customer
  I want to withdraw money from my account
  So that I can use my funds when needed

  Scenario: Withdraw with sufficient balance
    Given I have an active account with account number "ACC401" and balance 1000
    When I withdraw 400
    Then my account balance should be 600

  Scenario: Withdraw with insufficient funds
    Given I have an active account with account number "ACC402" and balance 100
    When I try to withdraw 500
    Then I should see an error "insufficient funds"
