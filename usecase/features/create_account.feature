@create
Feature: create account

  Scenario: happy path: create new account
    Given I want to create account with data from "test_account.json" file
    When I send POST request with that data create account
    Then the response code should be 201
    And response message should contain json with data matching request body"

  Scenario: sad path: error when trying to create account with id that already is in database
    Given I have account with data from "test_account.json" already in data base
    When I send POST request with the same data to create account
    Then the response code should be 409
    And  error message should be "Account cannot be created as it violates a duplicate constraint"

  Scenario: sad path: error when trying to create account without requestbody
    When I send POST request to create account without request body
    Then the response code should be 400
    And  error message should be "EOF"

  Scenario Outline: sad path: error when trying to create account with empty value of one of mandatory fields
    Given I want to create account with data from "test_account.json" file
    When I send POST request to create account setting "<mandatoryField>" mandatory field to hold empty value
    Then the response code should be 400
    And  error message should be "<errorMessage>"
    Examples:
      | mandatoryField         | errorMessage                                                                                                                                      |
      | id                     | validation failure list:\nvalidation failure list:\nid in body is required                                                                        |
      | organisationId         | validation failure list:\nvalidation failure list:\norganisation_id in body is required                                                           |
      | type                   | validation failure list:\nvalidation failure list:\ntype in body is required                                                                      |
      | country                | validation failure list:\nvalidation failure list:\nvalidation failure list:\ncountry in body should match '^[A-Z]{2}$'                           |
      | account_classification | validation failure list:\nvalidation failure list:\nvalidation failure list:\naccount_classification in body should be one of [Personal Business] |

  Scenario Outline: sad path: error when trying to create account when sending id or organisationId fields as a non UUID
    Given I want to create account with data from "test_account.json" file
    When I send POST request to create account setting "<fieldName>" field non UUID value of a
    Then the response code should be 400
#    And  error message should be "<errorMessage>"
#    TODO check error message, can`t escape " in response for now
    Examples:
      | fieldName      | errorMessage                                                                                            |
      | id             | validation failure list:\nvalidation failure list:\nid in body must be of type uuid: \"a\"              |
      | organisationId | validation failure list:\nvalidation failure list:\norganisation_id in body must be of type uuid: \"a\" |
