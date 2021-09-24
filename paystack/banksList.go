package paystack

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// Bank represents a Paystack bank
type Bank struct {
	Name        string      `json:"name"`
	Slug        string      `json:"slug"`
	Code        string      `json:"code"`
	Longcode    string      `json:"longcode"`
	Gateway     interface{} `json:"gateway"`
	PayWithBank bool        `json:"pay_with_bank"`
	Active      bool        `json:"active"`
	IsDeleted   bool        `json:"is_deleted"`
	Country     string      `json:"country"`
	Currency    string      `json:"currency"`
	Type        string      `json:"type"`
	ID          int         `json:"id"`
	CreatedAt   time.Time   `json:"createdAt"`
	UpdatedAt   time.Time   `json:"updatedAt"`
}

type BankListReponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Banks   []Bank `json:"data"`
}

const (
	bankListURL = "https://api.paystack.co/bank"
)

var (
	errorCouldNotGetBankList = errors.New("could not get current bank lists")
)

//Get Account Details makes a call to paystacks Account resolution API
func GetBanksList() ([]Bank, error) {
	request, err := http.NewRequest(http.MethodGet, bankListURL, nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{
		Timeout: defaultApiTimeout,
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		log.Printf("PAYSTACK ERROR bank lists query returned statusCode [%d] which seems to be invalid \n", response.StatusCode)
		return nil, errorMalformedResponse
	}

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	data := new(BankListReponse)
	err = json.Unmarshal([]byte(body), &data)

	if err != nil {
		return nil, err
	}

	if !data.Status {
		return nil, errorCouldNotGetBankList
	}

	return data.Banks, nil
}
