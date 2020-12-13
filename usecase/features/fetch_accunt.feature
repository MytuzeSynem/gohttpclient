@fetch
Feature: fetch account

  Scenario: happy path: get account
    Given I creating account with accountId "4b81b3ba-a5cd-45aa-a935-e6d8610bdef2"
    When I send GET request with "4b81b3ba-a5cd-45aa-a935-e6d8610bdef2" accountId to get account details
    Then the response code should be 200

  Scenario: sad path: error when trying to get account that doesn't exists
    When I send GET request with "4bc3060a-cc0d-4c29-acc4-5834b3c0e09f" accountId to get account details
    Then the response code should be 404
    And  error message should be "record 4bc3060a-cc0d-4c29-acc4-5834b3c0e09f does not exist"

  Scenario: sad path: error when trying to get account details passing non UUID accountId
    When I send GET request with "1" accountId to get account details
    Then the response code should be 400
    And  error message should be "id is not a valid uuid"