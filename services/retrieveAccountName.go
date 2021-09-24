package services

import (
	"errors"

	"github.com/CreamyMilk/buycoinz/paystack"
	"github.com/CreamyMilk/buycoinz/storage"
)

const (
	foundPreferedName   = true
	foundFromThirdParty = true
	accountNameNotFound = false
)

var (
	errorCouldNotFindAccountName = errors.New("could not find your account name")
	errorInvalidAccountNumber    = errors.New("provided account number is invalid")
	errorInvalidBankCode         = errors.New("provided bank_code is invalid")
)

type PreferedAccountNameRequest struct {
	BankCode      string
	AccountNumber string
}
type PreferedAccountNameResponse struct {
	AccountName string
	Message     string
	Status      bool
}

func (req *PreferedAccountNameRequest) GetPreferedAccountName() (*PreferedAccountNameResponse, error) {
	if req.AccountNumber == "" {
		return nil, errorInvalidAccountNumber
	}

	if req.BankCode == "" {
		return nil, errorInvalidBankCode
	}
	resp := new(PreferedAccountNameResponse)
	prefedName, found := storage.FakeStore[req.AccountNumber]
	if found {
		resp.Message = "Found users Preferd Account Name"
		resp.AccountName = prefedName
		resp.Status = foundPreferedName
		return resp, nil
	}
	details, err := paystack.GetAccoutDetails(req.AccountNumber, req.BankCode)
	if err != nil {
		resp.Message = "Please provide valid account credentials"
		resp.AccountName = ""
		resp.Status = accountNameNotFound
		return resp, errorCouldNotFindAccountName
	}

	resp.Message = "Successfully got default accountName"
	resp.AccountName = details.Data.AccountName
	resp.Status = foundFromThirdParty
	return resp, nil
}
