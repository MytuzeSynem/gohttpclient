package usecase

import (
	"errors"
	"fmt"
	"github.com/cucumber/godog"
	"http-client/gohttp"
	"http-client/usecase/util"
)

type apiFeature struct {
	resp        *gohttp.Response
	err         error
	requestBody *util.AccountDetails
}

type errorResponse struct {
	ErrorMessage string `json:"error_message"`
}

func (a *apiFeature) theResponseCodeShouldBe(statusCode int) error {

	if a.resp.StatusCode() != statusCode {
		return fmt.Errorf("should be a %d status code", statusCode)
	}
	return nil
}

func (a *apiFeature) errorMessageShouldBe(message string) error {
	var response errorResponse
	err := a.resp.UnmarshalJson(&response)
	if err != nil {
		return errors.New("can't parse response body to json")
	}

	if message != response.ErrorMessage {
		return fmt.Errorf("response should be \"%s\", not \"%s\"", message, response.ErrorMessage)
	}
	return nil
}

func createAccountForTest(accountId string) int {
	account := util.NewAccount(accountId)
	response, err := httpClient.Post("http://localhost:8080/v1/organisation/accounts/", account, nil, nil)
	if err != nil {
		fmt.Println(err)
	}
	return response.StatusCode()
}

func deleteTestAccount(accountId string) int {

	response, err := httpClient.Delete("http://localhost:8080/v1/organisation/accounts/"+accountId+"?version=0", nil, nil)
	if err != nil {
		fmt.Println(err)
	}
	return response.StatusCode()

}

func InitializeScenario(ctx *godog.ScenarioContext) {
	api := &apiFeature{}
	ctx.Step(`^I creating account with accountId "([^"]*)"$`, iCreatingAccountWithAccountId)
	ctx.Step(`^I send GET request with "([^"]*)" accountId to get account details`, api.iSendGETRequestWithAccountIdToGetAccountDetails)
	ctx.Step(`^the response code should be (\d+)$`, api.theResponseCodeShouldBe)
	ctx.Step(`^error message should be "([^"]*)"$`, api.errorMessageShouldBe)
	ctx.Step(`^I send Delete request with "([^"]*)" accountId to delete account$`, api.iSendDeleteRequestWithAccountIdToDeleteAccount)
	ctx.Step(`^I send Delete request with "([^"]*)" accountId without version query parameter$`, api.iSendDeleteRequestWithAccountIdWithoutVersionQueryParameter)
	ctx.Step(`^I want to create account with data from "([^"]*)" file$`, api.iWantToCreateAccountWithDataFromFile)
	ctx.Step(`^I send POST request with that data create account$`, api.iSendPOSTRequestWithThatDataCreateAccount)
	ctx.Step(`^response message should contain json with data matching request body"$`, api.responseMessageShouldContainJsonWithDataMatchingRequestBody)
	ctx.Step(`^I have account with data from "([^"]*)" already in data base$`, api.iHaveAccountWithDataFromAlreadyInDataBase)
	ctx.Step(`^I send POST request with the same data to create account$`, api.iSendPOSTRequestWithTheSameDataToCreateAccount)
	ctx.Step(`^I send POST request to create account without request body$`, api.iSendPOSTRequestToCreateAccountWithoutRequestBody)
	ctx.Step(`^I send POST request to create account setting "([^"]*)" mandatory field to hold empty value$`, api.iSendPOSTRequestToCreateAccountSettingMandatoryFieldToHoldEmptyValue)
	ctx.Step(`^I send POST request to create account setting "([^"]*)" field non UUID value of a$`, api.iSendPOSTRequestToCreateAccountSettingFieldNonUUIDValueOf)
	ctx.Step(`^I have two accounts based on data from "([^"]*)" with different accountsId$`, api.iHaveTwoAccountsBasedOnDataFromWithDifferentAccountsId)
	ctx.Step(`^I send GET request with page number (\d+) and page size (\d+)$`, api.iSendGETRequestWithPageNumberAndPageSize)
	ctx.Step(`^response should have (\d+) accounts on page$`, api.responseShouldHaveAccountsOnPage)
	ctx.Step(`^response shouldn\'t have next link$`, api.responseShouldntHaveNextLink)
	ctx.Step(`^i go to the next link$`, api.iGoToTheNextLink)
	ctx.Step(`^response should have one account on page$`, api.responseShouldHaveOneAccountOnPage)
	ctx.Step(`^response should have next link$`, api.responseShouldHaveNextLink)
	ctx.Step(`^response should have prev link$`, api.responseShouldHavePrevLink)


}
