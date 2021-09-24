package paystack

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// Reference: https://paystack.com/docs/identity-verification/resolve-account-number/

const (
	defaultApiTimeout = time.Minute
	resolverapiURL    = "https://api.paystack.co/bank/resolve"
)

type AccountResolutionResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    struct {
		AccountNumber string `json:"account_number"`
		AccountName   string `json:"account_name"`
		BankID        int    `json:"bank_id"`
	} `json:"data"`
}

var (
	errorPaystackInvalidToken = errors.New("seems that you supplied an invalid request valid ate your access tokens")
	errorMalformedResponse    = errors.New("the returned api response cannont be processed")
	errorBankDetailsNotFound  = errors.New("bank details not found")
)

//Get Account Details makes a call to paystacks Account resolution API
func GetAccoutDetails(accountNumber string, bankCode string) (*AccountResolutionResponse, error) {
	paystackAPITOKEN := getPaystackTOKEN()
	url := fmt.Sprintf("%s?account_number=%s&bank_code=%s", resolverapiURL, accountNumber, bankCode)
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	bearerToken := fmt.Sprintf("Bearer %s", paystackAPITOKEN)
	request.Header.Add("Authorization", bearerToken)

	client := &http.Client{
		Timeout: defaultApiTimeout,
	}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	if response.StatusCode == 401 {
		log.Printf("PAYSTACK ERROR  the suppiled [API_TOKEN: '%s' ] is Invalid \n", paystackAPITOKEN)
		return nil, errorPaystackInvalidToken
	}

	if response.StatusCode == 400 {
		return nil, errorBankDetailsNotFound
	}

	if response.StatusCode != http.StatusOK {
		log.Printf("PAYSTACK ERROR  returned statusCode [%d] which seems to be invalid \n", response.StatusCode)
		return nil, errorMalformedResponse
	}

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	details := new(AccountResolutionResponse)
	err = json.Unmarshal([]byte(body), &details)

	if err != nil {
		return nil, err
	}

	return details, nil
}
