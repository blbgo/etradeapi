package etradeapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func (r *etrade) MakeGetRequest(url string, form url.Values, result interface{}) error {
	fullURL := fmt.Sprint(r.baseServerURL, url, ".json")

	response, err := r.Client.Get(nil, r.AccessCred, fullURL, form)
	if err != nil {
		return err
	}

	return r.processResponse(response, result)
}

func (r *etrade) MakePostRequest(
	url string,
	form url.Values,
	body interface{},
	result interface{},
) error {
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return err
	}

	fullURL := fmt.Sprint(r.baseServerURL, url, ".json")

	req, err := http.NewRequest(http.MethodPost, fullURL, bytes.NewReader(bodyBytes))
	if err != nil {
		return err
	}

	err = r.Client.SetAuthorizationHeader(req.Header, r.AccessCred, http.MethodPost, req.URL, form)
	if err != nil {
		return err
	}
	req.URL.RawQuery = form.Encode()
	req.Header.Set("Content-Type", "application/json")

	response, err := r.RoundTripper.RoundTrip(req)
	if err != nil {
		return err
	}

	return r.processResponse(response, result)
}

func (r *etrade) processResponse(response *http.Response, result interface{}) error {
	defer response.Body.Close()

	responseBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	err = r.DumperFactory.Dump("APIResponse", responseBytes)
	if err != nil {
		return err
	}

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("Status code not OK: %v", response.StatusCode)
	}

	err = json.Unmarshal(responseBytes, result)
	if err != nil {
		return err
	}

	return nil
}
