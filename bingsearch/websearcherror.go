package bingsearch

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type WebsearchError struct {
	ErrorResponse struct {
		StatusCode int    `json:"statusCode"`
		Message    string `json:"message"`
	} `json:"error"`
}

func (we *WebsearchError) Error() string {
	return fmt.Sprintf("WebsearchError: Status: %d Message: %s", we.ErrorResponse.StatusCode, we.ErrorResponse.Message)
}

func getResponseError(res *http.Response) error {
	var webSearchError *WebsearchError
	err := json.NewDecoder(res.Body).Decode(&webSearchError)
	if err != nil {
		return err
	}
	return webSearchError
}
