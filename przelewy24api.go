package przelewy24api

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

type Przelewy24Api struct {
	UserToken string
	ApiToken  string
	Crc       string
	Live      bool
}

//getApiURL returns base API URL
func (api *Przelewy24Api) getApiURL() string {
	switch api.Live {
	case true:
		return "https://secure.przelewy24.pl/api/v1"
	case false:
		return "https://sandbox.przelewy24.pl/api/v1"
	}
	return ""
}

//sendRequest sends request to api with specified method, path and payload
func (api *Przelewy24Api) sendRequest(method, path string, payload interface{}) ([]byte, error) {
	client := &http.Client{}

	req, err := http.NewRequest(method, api.getApiURL()+path, bytes.NewReader(HTMLUnescapedJSON(payload)))
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(api.UserToken, api.ApiToken)
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	response, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return response, nil
}

//SendTransactionRegister sends request that registers transaction in Przelewy24 api. Returns unparsed json response and error if communication with API failed.
func (api *Przelewy24Api) SendTransactionRegister(tr *TransactionRegister) ([]byte, error) {
	tr.sign(api.Crc)
	return api.sendRequest(http.MethodPost, "/transaction/register", tr)
}

//SendTransactionVerify sends request that tries to verify transaction in Przelewy24 api. Returns unparsed json response and error if communication with API failed.
func (api *Przelewy24Api) SendTransactionVerify(tv *TransactionVerify) ([]byte, error) {
	tv.sign(api.Crc)
	return api.sendRequest(http.MethodPut, "/transaction/verify", tv)
}
