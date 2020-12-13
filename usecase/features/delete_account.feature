@del
Feature: delete account

  Scenario: happy path: delete existing account
    Given I creating account with accountId "23343911-447d-4d91-8762-ecfcf007942c"
    When I send Delete request with "23343911-447d-4d91-8762-ecfcf007942c" accountId to delete account
    Then the response code should be 204

  Scenario: sad path: deletion error when version is not specified
    Given I creating account with accountId "77e9e8f6-abe9-4d98-8acd-2d909b5dc5b9"
    When I send Delete request with "77e9e8f6-abe9-4d98-8acd-2d909b5dc5b9" accountId without version query parameter
    Then the response code should be 400
    And error message should be "invalid version number"