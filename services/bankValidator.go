package services

import (
	"log"

	"github.com/CreamyMilk/buycoinz/paystack"
	"github.com/CreamyMilk/buycoinz/storage"
	"github.com/texttheater/golang-levenshtein/levenshtein"
)

const (
	levenshteinThreshold = 2
	verifiedSuccessfully = true
	failedToVerify       = false
)

type BankVerificationRequest struct {
	UserAccountNumber string
	UserBankCode      string
	UserAccountName   string
}

type BankVerificationReponse struct {
	Status  bool
	Message string
}

func (req *BankVerificationRequest) Validate() (*BankVerificationReponse, error) {
	resp := new(BankVerificationReponse)
	details, err := paystack.GetAccoutDetails(req.UserAccountNumber, req.UserBankCode)

	if err != nil || details == nil {
		resp.Status = failedToVerify
		resp.Message = "Please supply valid Account Details"
		return resp, err
	}

	if req.UserAccountName == details.Data.AccountName && req.UserAccountNumber == details.Data.AccountNumber {
		resp.Status = verifiedSuccessfully
		resp.Message = "User Account details are correct and is now verified"

		//Stores Account Name
		storage.FakeStore[details.Data.AccountNumber] = req.UserAccountName
		return resp, err
	}

	if req.isAFuzzyMatch(details.Data.AccountName) {
		resp.Status = verifiedSuccessfully
		resp.Message = "Fuzzily verified bank Details but further verification is required"

		//Store Users prefs
		storage.FakeStore[details.Data.AccountNumber] = req.UserAccountName

		return resp, err
	}

	//We assume that if the score is below the threshold it is considerd as an invalid verfication attempt
	resp.Status = failedToVerify
	resp.Message = "Please supply a valid Account Details"

	return resp, nil
}

// isAFuzzyMatch returns whether an expected account is fuzzily matched
func (req *BankVerificationRequest) isAFuzzyMatch(expectedAccountName string) bool {
	score := levenshtein.DistanceForStrings([]rune(req.UserAccountName), []rune(expectedAccountName), levenshtein.DefaultOptions)
	result := score <= levenshteinThreshold
	log.Printf("Computed Leven... Distance for %s with %s yielding a score of %d ", req.UserAccountName, expectedAccountName, score)
	return result
}
