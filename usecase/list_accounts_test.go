package usecase

import (
	"errors"
	"fmt"
	"http-client/usecase/util"
	"strconv"
)

type AccountDetailsList struct {
	Data  []util.Data `json:"data"`
	Links Links       `json:"links"`
}

type Links struct {
	First string `json:"first"`
	Last  string `json:"last"`
	Self  string `json:"self"`
	Next  string `json:"next"`
	Prev  string `json:"prev"`
}

func (a *apiFeature) iHaveTwoAccountsBasedOnDataFromWithDifferentAccountsId(file string) error {
	a.requestBody = getTestAccount(file)
	ids := []string{"37b653f0-c157-4c43-8049-905b931bcc0d", "77e1f8b5-6473-42fd-b0cd-f37035b6904f"}

	for _, value := range ids {
		a.requestBody.Data.ID = value
		a.resp, a.err = httpClient.Post("http://localhost:8080/v1/organisation/accounts/", a.requestBody, nil, nil)
		if a.err != nil {
			return errors.New("can't connect to tested api")
		}
	}
	return nil
}

func (a *apiFeature) iSendGETRequestWithPageNumberAndPageSize(pageNumber, pageSize int) error {

	queryParams := map[string]string{
		"page[number]": strconv.Itoa(pageNumber),
		"page[size]":   strconv.Itoa(pageSize),
	}
	a.resp, a.err = httpClient.Get("http://localhost:8080/v1/organisation/accounts/", nil, queryParams)
	if a.err != nil {
		return errors.New("can't connect to tested api")
	}
	return nil
}

func (a *apiFeature) responseShouldHaveAccountsOnPage(accounts int) error {
	var responseBody AccountDetailsList
	err := a.resp.UnmarshalJson(&responseBody)
	if err != nil {
		return errors.New("can't parse response body to json")
	}
	if len(responseBody.Data) != accounts {
		return fmt.Errorf("response body contains %d accounts, but should have %d", len(responseBody.Data), accounts)
	}
	return nil
}

func (a *apiFeature) responseShouldntHaveNextLink() error {
	var responseBody AccountDetailsList
	err := a.resp.UnmarshalJson(&responseBody)
	if err != nil {
		return errors.New("can't parse response body to json")
	}
	if responseBody.Links.Next != "" {
		return errors.New("response body shouldn't have next link")
	}
	return nil
}

func (a *apiFeature) responseShouldHaveOneAccountOnPage() error {
	var responseBody AccountDetailsList
	err := a.resp.UnmarshalJson(&responseBody)
	if err != nil {
		return errors.New("can't parse response body to json")
	}
	if len(responseBody.Data) != 1 {
		return fmt.Errorf("response body contains %d accounts, but should have 1", len(responseBody.Data))
	}
	return nil
}

func (a *apiFeature) responseShouldHaveNextLink() error {
	var responseBody AccountDetailsList
	err := a.resp.UnmarshalJson(&responseBody)
	if err != nil {
		return errors.New("can't parse response body to json")
	}
	if responseBody.Links.Next == "" {
		return errors.New("response body should have next link")
	}
	return nil
}

func (a *apiFeature) iGoToTheNextLink() error {
	var responseBody AccountDetailsList
	err := a.resp.UnmarshalJson(&responseBody)
	if err != nil {
		return errors.New("can't parse response body to json")
	}

	a.resp, a.err = httpClient.Get("http://localhost:8080"+responseBody.Links.Next, nil, nil)
	if a.err != nil {
		return errors.New("can't connect to tested api")
	}
	return nil
}

func (a *apiFeature) responseShouldHavePrevLink() error {
	var responseBody AccountDetailsList
	err := a.resp.UnmarshalJson(&responseBody)
	if err != nil {
		return errors.New("can't parse response body to json")
	}
	if responseBody.Links.Prev == "" {
		return errors.New("response body should have prev link")
	}
	return nil
}
