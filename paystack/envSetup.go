package paystack

import (
	"os"
)

func getPaystackTOKEN() string {
	var validKey string
	if os.Getenv("ENV") == "production" {
		validKey = os.Getenv("PAYSTACK_PROD_KEY")
	} else {
		validKey = os.Getenv("PAYSTACK_TEST_KEY")
	}
	return validKey
}
