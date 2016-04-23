package vtsdk

import (
	"fmt"
	"net/http"
)

const (
	APISandBox = "https://api.sandbox.veritrans.co.id/v2"
	APILive    = "https://api.veritrans.co.id/v2"
)

type (

	// Payment http://docs.veritrans.co.id/en/vtweb/integration_php.html
	Client struct {
		client       *http.Client
		ClientID     string
		ClientSecret string
		APIBase      string
		LogFile      string
		Token        *TokenBase64
	}

	TokenBase64 struct {
		Token string `json:"token"`
	}

	TokenSecure struct {
		StatusCode    string `json:"status_code"`
		StatusMessage string `json:"status_message"`
		TokenID       string `json:"token_id"`
	}

	TokenRequest struct {
		CardNumber     string `json:"card_number"`
		CardCVV        string `json:"card_cvv"`
		CardExpMonth   string `json:"card_exp_month"`
		CardExpYear    string `json:"card_exp_year"`
		GrossAmount    string `json:"gross_amount"`
		Secure         string `json:"secure"`
		SecureCallback string `json:"secure_callback"`
		Callback       string `json:"callback"`
		ClientID       string `json:"ClientID"`
	}

	Payment struct {
		PaymentType        string             `json:"payment_type"`
		CreditCard         CreditCard         `json:"credit_card"`
		TransactionDetails TransactionDetails `json:"transaction_details"`
	}

	ItemDetails struct {
		ID       string `json:"id"`
		Price    string `json:"price"`
		Quantity string `json:"quantity"`
		Name     string `json:"name"`
	}

	CreditCard struct {
		TokenID string `json:"token_id"`
	}

	TransactionDetails struct {
		GrossAmount int    `json:"gross_amount"`
		OrderId     string `json:"order_id"`
	}

	BillingAddress struct {
		FirstName   string `json:"first_name"`
		LastName    string `json:"last_name"`
		Address     string `json:"address"`
		City        string `json:"city"`
		PostalCode  string `json:"postal_code"`
		Phone       string `json:"phone"`
		CountryCode string `json:"country_code"`
	}

	ShippingAddress struct {
		FirstName   string `json:"first_name"`
		LastName    string `json:"last_name"`
		Address     string `json:"address"`
		City        string `json:"city"`
		PostalCode  string `json:"postal_code"`
		Phone       string `json:"phone"`
		CountryCode string `json:"country_code"`
	}

	ErrorResponse struct {
		Response        *http.Response `json:"-"`
		Name            string         `json:"name"`
		DebugID         string         `json:"debug_id"`
		Message         string         `json:"message"`
		InformationLink string         `json:"information_link"`
		Details         []ErrorDetail  `json:"details"`
	}

	ErrorDetail struct {
		Field string `json:"field"`
		Issue string `json:"issue"`
	}
)

func (r *ErrorResponse) Error() string {
	return fmt.Sprintf("%v %v: %d %v", r.Response.Request.Method, r.Response.Request.URL, r.Response.StatusCode, r.Message)
}
