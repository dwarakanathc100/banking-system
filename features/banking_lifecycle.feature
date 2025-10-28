Feature: Full Banking Lifecycle
  As a customer
  I want to perform end-to-end banking actions
  So that I can see how my balance changes throughout operations

  Scenario: Create account, deposit, withdraw, take loan, pay EMI
    Given I create a new account with account number "ACC999" and initial balance 1000
    When I deposit 500
    And I withdraw 200
    And I apply for a loan of 2000
    And I pay EMI of 800
    Then my final account balance should be 2500
    And my account should remain active
