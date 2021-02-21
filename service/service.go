package service

import (
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/ryanda/product-api-test/config"
)

const LIMIT_DATA = "5"
const SERVICE_ID = "Golang-Rest-Ryanda"

func FetchProductData() string {

	endpoint := config.BackendUrl + "/wc/v3/products"

	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return "Error:" + err.Error()
	}

	req.Header.Set("X-Service", SERVICE_ID)
	req.Header.Set("Content-Type", "application/json")

	q := url.Values{}
	q.Add("per_page", LIMIT_DATA)
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
			return "Error occurred due to a timeout."
		}

		// log error and exit
		return "Error:" + err.Error()
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	return string(body)
}
