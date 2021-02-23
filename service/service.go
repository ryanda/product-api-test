package service

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/ryanda/product-api-test/config"
	"github.com/ryanda/product-api-test/obj"
)

const LIMIT_DATA = "5"
const ORDER_DESC = "desc"
const ORDER_BY = "date"
const STATUS_PUBLISH = "publish"
const SERVICE_ID = "Golang-Rest-Ryanda"

func FetchProductData() ([]byte, error) {

	endpoint := config.BackendUrl + "/wc/v3/products"

	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("X-Service", SERVICE_ID)
	req.Header.Set("Content-Type", "application/json")

	q := url.Values{}
	q.Add("per_page", LIMIT_DATA)
	q.Add("order", ORDER_DESC)
	q.Add("orderby", ORDER_BY)
	q.Add("status", STATUS_PUBLISH)
	q.Add("consumer_key", config.ConsumerKey)
	q.Add("consumer_secret", config.ConsumerSecret)
	req.URL.RawQuery = q.Encode()

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		// get `url.Error` struct pointer from `err` interface
		urlErr := err.(*url.Error)

		// check if error occurred due to timeout
		if urlErr.Timeout() {
			return nil, errors.New("Error occurred due to a timeout.")
		}

		// return generic err
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		errResponse, _ := obj.UnmarshalErrorResponse(body)
		return nil, fmt.Errorf("%s: %s", errResponse.Code, errResponse.Message)
	}

	return body, nil
}
