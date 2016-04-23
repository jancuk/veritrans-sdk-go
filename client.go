package vtsdk

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"os"
)

func NewClient(ClientID string, ClientSecret string, APIBase string) (*Client, error) {

	if ClientID == "" || ClientSecret == "" || APIBase == "" {
		return &Client{}, errors.New("ClientID, ClientSecret, and APIBase are required")
	}

	return &Client{
		&http.Client{},
		ClientID,
		ClientSecret,
		APIBase,
		"",
		nil,
	}, nil

}

func (c *Client) Send(req *http.Request, v interface{}) error {
	// Set default headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	resp, err := c.client.Do(req)

	c.log(req, resp)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		errResp := &ErrorResponse{Response: resp}
		data, err := ioutil.ReadAll(resp.Body)

		if err == nil && len(data) > 0 {
			json.Unmarshal(data, errResp)
		}

		return errResp
	}

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			io.Copy(w, resp.Body)
		} else {
			err = json.NewDecoder(resp.Body).Decode(v)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (c *Client) SendWithAuthorization(req *http.Request, v interface{}) error {
	req.Header.Set("Authorization", c.Token.Token)

	return c.Send(req, v)
}

func (c *Client) NewRequest(method, url string, payment interface{}) (*http.Request, error) {
	var buf io.Reader
	if payment != nil {
		var b []byte
		b, err := json.Marshal(&payment)
		if err != nil {
			return nil, err
		}
		buf = bytes.NewBuffer(b)
	}

	return http.NewRequest(method, url, buf)
}

func (c *Client) log(req *http.Request, resp *http.Response) {
	if c.LogFile != "" {
		os.OpenFile(c.LogFile, os.O_CREATE, 0755)

		logFile, err := os.OpenFile(c.LogFile, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0755)
		if err == nil {
			reqDump, _ := httputil.DumpRequestOut(req, true)
			respDump, _ := httputil.DumpResponse(resp, true)

			logFile.WriteString("Request: " + string(reqDump) + "\nResponse: " + string(respDump) + "\n\n")
		}
	}
}
