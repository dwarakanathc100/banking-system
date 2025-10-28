Feature: User Account Management
  As a bank customer
  I want to create and manage my account
  So that I can perform banking operations

  Scenario: Create a new user account
    Given I am a new customer named "Dwarkanath"
    When I open a new account with account number "ACC101" and initial balance 1000
    Then my account should be active
    And my account balance should be 1000

  Scenario: Close an existing account
    Given I have an active account with account number "ACC202"
    When I close my account
    Then my account should be inactive
