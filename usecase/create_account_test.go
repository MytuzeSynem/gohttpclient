package usecase

import (
	"encoding/json"
	"errors"
	"fmt"
	"http-client/usecase/util"
	"io/ioutil"
)

func (a *apiFeature) iWantToCreateAccountWithDataFromFile(file string) error {
	a.requestBody = getTestAccount(file)
	return nil
}

func (a *apiFeature) iSendPOSTRequestWithThatDataCreateAccount() error {

	a.resp, a.err = httpClient.Post("http://localhost:8080/v1/organisation/accounts/", a.requestBody, nil, nil)
	if a.err != nil {
		return errors.New("can't connect to tested api")
	}
	defer deleteTestAccount(a.requestBody.Data.ID)
	return nil
}

func (a *apiFeature) responseMessageShouldContainJsonWithDataMatchingRequestBody() error {
	var responseBody util.AccountDetails
	err := a.resp.UnmarshalJson(&responseBody)
	if err != nil {
		return errors.New("can't parse response body to json")
	}

	if *a.requestBody != responseBody {
		return errors.New("response body doesn't match request body")
	}

	return nil
}

func (a *apiFeature) iHaveAccountWithDataFromAlreadyInDataBase(file string) error {
	a.requestBody = getTestAccount(file)
	a.resp, a.err = httpClient.Post("http://localhost:8080/v1/organisation/accounts/", a.requestBody, nil, nil)
	if a.err != nil {
		return errors.New("can't connect to tested api")
	}
	return nil
}

func (a *apiFeature) iSendPOSTRequestWithTheSameDataToCreateAccount() error {
	a.resp, a.err = httpClient.Post("http://localhost:8080/v1/organisation/accounts/", a.requestBody, nil, nil)
	if a.err != nil {
		return errors.New("can't connect to tested api")
	}
	defer deleteTestAccount(a.requestBody.Data.ID)
	return nil
}

func (a *apiFeature) iSendPOSTRequestToCreateAccountWithoutRequestBody() error {
	a.resp, a.err = httpClient.Post("http://localhost:8080/v1/organisation/accounts/", nil, nil, nil)
	if a.err != nil {
		return errors.New("can't connect to tested api")
	}
	return nil
}

func (a *apiFeature) iSendPOSTRequestToCreateAccountSettingMandatoryFieldToHoldEmptyValue(mandatoryField string) error {
	switch mandatoryField {
	case "id":
		a.requestBody.Data.ID = ""
	case "organisationId":
		a.requestBody.Data.OrganisationID = ""
	case "type":
		a.requestBody.Data.Type = ""
	case "country":
		a.requestBody.Data.Attributes.Country = ""
	case "account_classification":
		a.requestBody.Data.Attributes.AccountClassification = ""

	}
	a.resp, a.err = httpClient.Post("http://localhost:8080/v1/organisation/accounts/", a.requestBody, nil, nil)
	if a.err != nil {
		return errors.New("can't connect to tested api")
	}
	return nil
}

func (a *apiFeature) iSendPOSTRequestToCreateAccountSettingFieldNonUUIDValueOf(fieldName string) error {
	switch fieldName {
	case "id":
		a.requestBody.Data.ID = "1"
	case "organisationId":
		a.requestBody.Data.OrganisationID = "1"
	}
	a.resp, a.err = httpClient.Post("http://localhost:8080/v1/organisation/accounts/", a.requestBody, nil, nil)
	if a.err != nil {
		return errors.New("can't connect to tested api")
	}
	return nil
}


func getTestAccount(file string) *util.AccountDetails {

	var data []byte
	data, _ = ioutil.ReadFile("test-files/" + file)

	var str util.AccountDetails
	err := json.Unmarshal(data, &str)

	if err != nil {
		fmt.Println(err)
	}

	return &str
}
