package usecase

import (
	"errors"
)

func (a *apiFeature) iSendDeleteRequestWithAccountIdToDeleteAccount(accountId string) error {
	a.resp, a.err = httpClient.Delete("http://localhost:8080/v1/organisation/accounts/" +accountId + "?version=0", nil, nil)
	if a.err != nil {
		return errors.New("can't connect to tested api")
	}
	return nil
}

func (a * apiFeature) iSendDeleteRequestWithAccountIdWithoutVersionQueryParameter(accountId string) error {
	a.resp, a.err = httpClient.Delete("http://localhost:8080/v1/organisation/accounts/" +accountId, nil, nil)
	if a.err != nil {
		return errors.New("can't connect to tested api")
	}
	defer deleteTestAccount(accountId)
	return nil
}


