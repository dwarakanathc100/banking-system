Feature: Deposit Money
  As a customer
  I want to deposit money into my account
  So that I can increase my balance

  Scenario: Deposit money into active account
    Given I have an active account with account number "ACC303" and balance 1000
    When I deposit 500 into my account
    Then my account balance should become 1500

  Scenario: Deposit money into inactive account
    Given I have an inactive account with account number "ACC304"
    When I try to deposit 200 into my account
    Then I should see an error "account is not active"
