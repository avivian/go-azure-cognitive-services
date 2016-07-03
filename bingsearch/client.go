package bingsearch

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

var (
	SubscriptionKeyHeader = "Ocp-Apim-Subscription-Key"
)

type service struct {
	client *Client
}

type Client struct {
	client http.Client

	common service

	subscriptionKey string

	Web    *WebService
	Images *ImageService
	Videos *VideoService
	News   *NewsService
}

func NewClient(subscriptionKey string) *Client {

	c := &Client{
		client:          *http.DefaultClient,
		subscriptionKey: subscriptionKey,
	}
	c.common.client = c

	c.Web = (*WebService)(&c.common)
	c.Images = (*ImageService)(&c.common)
	c.Videos = (*VideoService)(&c.common)
	c.News = (*NewsService)(&c.common)

	return c
}

func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {

	var buf io.ReadWriter
	if body != nil {
		buf := new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, urlStr, buf)
	if err != nil {
		return nil, err
	}

	req.Header.Add(SubscriptionKeyHeader, c.subscriptionKey)
	req.Header.Add("Accept", "application/json")

	return req, nil
}

func (c *Client) Do(req *http.Request, body interface{}) error {
	res, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&body)
	if err != nil {
		return err
	}

	fmt.Printf("%d\n", res.StatusCode)

	// TODO: parse response
	return nil
}

func (c *Client) search(urlStr string, params *SearchQueryParams, result interface{}) error {
	req, err := c.NewRequest("GET", urlStr, nil)
	if err != nil {
		return err
	}
	req.URL.RawQuery = getSearchRawQuery(req, params)
	err = c.Do(req, &result)
	return err

}
