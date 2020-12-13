package usecase

import (
	"errors"
)

func iCreatingAccountWithAccountId(accountId string) error {
	createAccountForTest(accountId)
	return nil
}

func (a *apiFeature) iSendGETRequestWithAccountIdToGetAccountDetails(accountId string) error {

	a.resp, a.err = httpClient.Get("http://localhost:8080/v1/organisation/accounts/"+accountId, nil, nil)
	if a.err != nil {
		return errors.New("can't connect to tested api")
	}
	defer deleteTestAccount(accountId)
	return nil
}
