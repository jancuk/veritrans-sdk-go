package main

import (
	"fmt"
	vtsdk "github.com/jancuk/veritrans-sdk-go"
)

func main() {
	c, _ := vtsdk.NewClient("VT-CLIENT-ID", "VT-SERVER-KEY", vtsdk.APISandBox)

	c.GetTokenAuthorize()

	p := vtsdk.Payment{
		PaymentType: "VTWEB",
		TransactionDetails: vtsdk.TransactionDetails{
			GrossAmount: 100000,
			OrderId:     "my-unique-order-id-123454",
		},
	}

	paymentResponse, err := c.CreatePayment(p)

	if err != nil {
		panic(err)
	}

	fmt.Println(paymentResponse.RedirectURL)
	fmt.Println(paymentResponse.StatusCode)
	fmt.Println(paymentResponse.StatusMessage)

}
