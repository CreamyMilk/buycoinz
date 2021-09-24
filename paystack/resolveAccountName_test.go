package paystack

import (
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

// Test of paystack intergration [setup the .env.tests]
func TestResolveAccountNumber(t *testing.T) {
	//Log all the required Configaraition for the tests
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env.tests file", err)
	}

	validTestAccountNo := os.Getenv("TEST_ACCOUNTNO")
	validTestBankCode := os.Getenv("TEST_BANKCODE")
	validTestAccountName := os.Getenv("TEST_VALID_ACCOUNT_NAME")

	if validTestAccountNo == "" || validTestBankCode == "" || validTestAccountName == "" {
		t.Error("You should provide [TEST_ACCOUNT ENVIRONMENT VARIABLES DETAILS] in order to run this test")
		return
	}

	resp, err := GetAccoutDetails(validTestAccountNo, validTestBankCode)

	if err != nil {
		t.Errorf("Didn't expected an error but got this error[%s]'", err.Error())
		return
	}

	if resp.Status {
		if resp.Data.AccountName == "" {
			t.Error("Expected response to contain an 'account_number'")
			return
		}
		if resp.Data.BankID == 0 {
			t.Error("Expected response to have a 'BankID'")
			return
		}
		if resp.Data.AccountName != validTestAccountName {
			t.Errorf("Expected the retured Account Number to be [%s] but we got [%s]", validTestAccountName, resp.Data.AccountName)
			return
		}
	} else {
		t.Errorf("Expected the provided request wold have exectured successfuly but we got \n %+v", resp)
		return
	}
}
