package util

type AccountDetails struct {
	Data Data `json:"data"`
}
type Data struct {
	ID             string     `json:"id"`
	OrganisationID string     `json:"organisation_id"`
	Type           string     `json:"type"`
	Attributes     Attributes `json:"attributes"`
}
type Attributes struct {
	Country                 string   `json:"country"`
	BaseCurrency            string   `json:"base_currency"`
	BankID                  string   `json:"bank_id"`
	BankIDCode              string   `json:"bank_id_code"`
	CustomerID              string   `json:"customer_id"`
	Bic                     string   `json:"bic"`
	AccountClassification   string   `json:"account_classification"`
	JointAccount            bool     `json:"joint_account"`
	AccountMatchingOptOut   bool     `json:"account_matching_opt_out"`
	SecondaryIdentification string   `json:"secondary_identification"`
}

func NewAccount(accountId string) *AccountDetails {
	a := AccountDetails{
		Data: Data{
			ID:             accountId,
			OrganisationID: "4bc3060a-cc0d-4c29-acc4-5834b3c0e09f",
			Type:           "accounts",
			Attributes: Attributes{
				Country:      "GB",
				BaseCurrency: "GBP",
				BankID:       "400302",
				BankIDCode:   "GBDSC",
				CustomerID:   "234",
				Bic:          "NWBKGB42",
				AccountClassification:   "Personal",
				JointAccount:            false,
				AccountMatchingOptOut:   false,
				SecondaryIdentification: "A1B2C3D4",
			},
		},
	}
	return &a
}
