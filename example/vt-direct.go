package main

import (
	"fmt"
	vtsdk "github.com/jancuk/veritrans-sdk-go"
)

func main() {
	c, _ := vtsdk.NewClient("VT-CLIENT-ID", "VT-SERVER-KEY", vtsdk.APISandBox)

	c.GetTokenAuthorize()
	createToken := vtsdk.TokenRequest{
		CardNumber:   "4811 1111 1111 1114",
		CardCVV:      "123",
		CardExpMonth: "12",
		CardExpYear:  "2016",
		GrossAmount:  "3000",
		Secure:       "false", // not support 3ds
		ClientID:     c.ClientID,
	}

	geTokenID, _ := c.GetAccessTokenID(createToken)

	p := vtsdk.Payment{
		PaymentType: "credit_card",
		CreditCard: vtsdk.CreditCard{
			TokenID: geTokenID.TokenID,
		},
		TransactionDetails: vtsdk.TransactionDetails{
			GrossAmount: 100000,
			OrderId:     "my-unique-order-id-99999",
		},
	}

	paymentResponse, _ := c.CreatePayment(p)

	fmt.Println(paymentResponse.StatusCode)
	fmt.Println(paymentResponse.StatusMessage)

}
