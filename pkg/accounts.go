package qtapi

import (
	"encoding/json"
	"log"
)

type Account struct {
	Type              string `json:"type"`
	Number            string `json:"number"`
	Status            string `json:"status"`
	IsPrimary         bool   `json:"isPrimary"`
	IsBilling         bool   `json:"isBilling"`
	ClientAccountType string `json:"clientAccountType"`
}

func (q Questrade) getAccounts() ([]Account, error) {
	accounts := []Account{}
	res, err := q.request("v1/accounts")
	if err != nil {
		log.Println(err)
		return accounts, err
	}

	if res == nil {
		return nil, nil
	}

	json.Unmarshal(res, &accounts)

	return accounts, nil
}
