package vtsdk

import (
	"fmt"
	"testing"
)

var testClientID = "AZgwu4yt5Ba0gyTu1dGBH3txHCJbMuFNvrmQxBaQbfDncDiCs6W_rwJD8Ir-0pZrN-_eq7n9zVd8Y-5f"
var testSecret = "EBzA1wRl5t73OMugOieDj_tI3vihfJmGl47ukQT-cpctooIzDu0K7IPESNC0cKodlLSOXzwI8qXSM0rd"

func TestNewClient(t *testing.T) {
	_, err := NewClient("", "", "")
	if err == nil {
		t.Errorf("All arguments are required in NewClient()")
	} else {
		fmt.Println(err.Error())
	}

	_ , err = NewClient(testClientID, testSecret, APISandBox)
	if err != nil {
		t.Errorf("NewClient() must not return error for valid creds: " + err.Error())
	}

}
