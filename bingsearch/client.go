package bingsearch

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type BingSearchClient struct {
	SubscriptionKey string
	client          *http.Client
}

type SearchQueryParams struct {
	Query      string
	Count      string
	Offset     string
	Mkt        string
	SafeSearch string
}

func NewSearchQueryParams(q string, count int, offset int, mkt string, safeSearch string) *SearchQueryParams {
	return &SearchQueryParams{
		Query:      q,
		Count:      strconv.Itoa(count),
		Offset:     strconv.Itoa(offset),
		Mkt:        mkt,
		SafeSearch: safeSearch,
	}
}

func NewClient(subscriptionKey string) *BingSearchClient {
	return &BingSearchClient{
		SubscriptionKey: subscriptionKey,
		client:          &http.Client{},
	}
}

func getSearchRawQuery(req *http.Request, params *SearchQueryParams) string {
	query := req.URL.Query()
	query.Add("q", params.Query)
	query.Add("count", params.Count)
	query.Add("offset", params.Offset)
	query.Add("mkt", params.Mkt)
	query.Add("safeSearch", params.SafeSearch)
	return query.Encode()
}

func (bsc *BingSearchClient) searchRequest(url string, params *SearchQueryParams, result interface{}) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	req.URL.RawQuery = getSearchRawQuery(req, params)
	req.Header.Add("Ocp-Apim-Subscription-Key", bsc.SubscriptionKey)

	res, err := bsc.client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return getResponseError(res)
	}

	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return err
	}
	return nil
}
