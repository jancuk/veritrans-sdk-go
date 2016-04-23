package vtsdk

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func (c *Client) GetTokenAuthorize() *TokenBase64 {
	str := base64.StdEncoding.EncodeToString([]byte(c.ClientSecret + ":"))
	t := &TokenBase64{
		Token: str,
	}
	c.Token = t
	return t
}

func (c *Client) GetAccessTokenID(t TokenRequest) (*TokenSecure, error) {
	params := url.Values{}

	params.Add("card_number", url.QueryEscape(t.CardNumber))
	params.Add("card_cvv", t.CardCVV)
	params.Add("card_exp_month", t.CardExpMonth)
	params.Add("card_exp_year", t.CardExpYear)
	params.Add("gross_amount", t.GrossAmount)
	// not support 3ds :(
	params.Add("secure", t.Secure)
	params.Add("client_key", c.ClientID)

	response, err := http.Get(c.APIBase + "/token?" + params.Encode())

	tokenSecure := &TokenSecure{}

	if err != nil {
		fmt.Printf("%s", err)
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Printf("%s", err)
		}
		json.Unmarshal(contents, &tokenSecure)
	}

	return tokenSecure, nil
}
