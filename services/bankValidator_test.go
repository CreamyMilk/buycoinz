package services

import (
	"log"
	"testing"

	"github.com/joho/godotenv"
)

//TestValidateBankDetails test
func TestValidateAccountNameMatches(t *testing.T) {

	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env.tests file", err)
	}

	//Test Table
	tt := []struct {
		testName                string
		actualAccountName       string
		userSuppliedAccountName string
		shouldMatch             bool
	}{
		{"Simple Single letter deletion", "Buycoins", "Buycoi", true},
		{"Simple Single letter Addittion", "Buycoins", "Buycoinss", true},
		{"Simple Single letter substituion", "Buycoins", "buycoins", true},
		{"Word with spaces", "Paystack Limited", "PaystackLimited", true},
		{"Extra Padding added Name", "Buycoins", " Buycoins ", true},
		{"Three letters missing", "Buycoins", "Buyco", false},
		{"Exact Opposite Name", "Black", "White", false},
	}

	for _, tc := range tt {
		t.Run(tc.testName, func(t *testing.T) {
			verificationReq := BankVerificationRequest{
				UserAccountName: tc.actualAccountName,
			}
			if verificationReq.isAFuzzyMatch(tc.userSuppliedAccountName) != tc.shouldMatch {
				t.Errorf("Expected the  fuzzy matched between (%s) and (%s) ", verificationReq.UserAccountName, tc.actualAccountName)
				return
			}
		})
	}
}
