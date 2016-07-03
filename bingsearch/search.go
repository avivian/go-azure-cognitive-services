package bingsearch

import (
	"net/http"
	"strconv"
)

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

func getSearchRawQuery(req *http.Request, params *SearchQueryParams) string {
	query := req.URL.Query()
	query.Add("q", params.Query)
	query.Add("count", params.Count)
	query.Add("offset", params.Offset)
	query.Add("mkt", params.Mkt)
	query.Add("safeSearch", params.SafeSearch)
	return query.Encode()
}
