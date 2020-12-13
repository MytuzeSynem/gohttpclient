@list

Feature: list accounts

  Scenario: happy path: list 2 accounts on one page
    Given I have two accounts based on data from "test_account.json" with different accountsId
    When I send GET request with page number 0 and page size 2
    Then the response code should be 200
    And response should have 2 accounts on page
    And response shouldn't have next link

  Scenario: happy path: list 2 accounts on two pages
    Given I have two accounts based on data from "test_account.json" with different accountsId
    When I send GET request with page number 0 and page size 1
    Then the response code should be 200
    And response should have one account on page
    And response should have next link
    When i go to the next link
    Then the response code should be 200
    And response should have one account on page
    And response should have prev link