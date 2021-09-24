package services

import (
	"log"
	"os"
	"testing"

	"github.com/CreamyMilk/buycoinz/storage"
	"github.com/joho/godotenv"
)

//TestGetPreferedAccountName tests account name retrival using the
//live paystack api (hence you need to provide some api_keys to run them)
//so as to ensure it still conforms with expexcted responses
func TestGetPreferedAccountName(t *testing.T) {
	storage.InitalizeDB()
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env.tests file", err)
	}

	validTestBankCode := os.Getenv("TEST_BANKCODE")
	validTestAccountNo := os.Getenv("TEST_ACCOUNTNO")
	validTestAccountName := os.Getenv("TEST_VALID_ACCOUNT_NAME")
	validTestPreferdAccountName := os.Getenv("TEST_PREFERED_ACCOUNT_NAME")
	//Test Table
	tt := []struct {
		testName            string
		bankCode            string
		persistName         bool
		accountNumber       string
		preferedAccountName string
		expectedName        string
		expectedError       error
	}{
		{"Empty Bank code", "", false, "123456789", "", "", errorInvalidBankCode},
		{"Empty Account Number", "123", false, "", "", "", errorInvalidAccountNumber},
		{"Stored Name", "100", true, "10293929292", "Mr John Doe", "Mr John Doe", nil},
		{"FakeDetail", "000", false, "65656565656", "", "sholdnotbereached", errorCouldNotFindAccountName},
		{"Correct Details with no Prefered Account Name", validTestBankCode, false, validTestAccountNo, validTestAccountNo, validTestAccountName, nil},
		{"Correct Details with  a  Prefered Account Name", validTestBankCode, true, validTestAccountNo, validTestPreferdAccountName, validTestPreferdAccountName, nil},
	}

	for _, tc := range tt {
		t.Run(tc.testName, func(t *testing.T) {
			nameReq := PreferedAccountNameRequest{
				AccountNumber: tc.accountNumber,
				BankCode:      tc.bankCode,
			}
			if tc.persistName {
				storage.FakeStore[tc.accountNumber] = tc.preferedAccountName
			}

			res, err := nameReq.GetPreferedAccountName()
			if err != tc.expectedError {
				t.Errorf("Expected the returned error to be (%s) but got [%s] for [%+v]", tc.expectedError.Error(), err.Error(), nameReq)
				return
			}

			if res != nil && res.AccountName != tc.expectedName && res.Status == true {
				t.Errorf("Expected Account name (%s) to be [%s] request ", res.AccountName, tc.expectedName)
				return
			}
		})
	}
}
