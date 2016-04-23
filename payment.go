package vtsdk

type ResponsePayment struct {
	RedirectURL   string `json:"redirect_url"`
	StatusCode    string `json:"status_code"`
	StatusMessage string `json:"status_message"`
}

func (c *Client) CreatePayment(p Payment) (*ResponsePayment, error) {

	req, err := c.NewRequest("POST", c.APIBase+"/charge", p)

	if err != nil {
		return &ResponsePayment{}, err
	}

	response := &ResponsePayment{}

	err = c.SendWithAuthorization(req, response)

	if err != nil {
		return response, err
	}

	return response, nil
}
