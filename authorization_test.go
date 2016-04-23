package vtsdk

import (
	"testing"
)

func TestGetTokenAuthorize(t *testing.T) {
	c, _ := NewClient(testClientID, testSecret, APISandBox)
	token := c.GetTokenAuthorize()
	if token.Token == "" {
		t.Errorf("Token is not returned by GetTokenAuthorize")
	}
}
