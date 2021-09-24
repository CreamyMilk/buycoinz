package paystack

import "testing"

func TestGetBankList(t *testing.T) {
	banks, err := GetBanksList()

	if err != nil || !(len(banks) > 0) {
		t.Errorf("Expected Bank list, got %d, returned error %v", len(banks), err)
	}
}
