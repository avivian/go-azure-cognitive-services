package cognitiveservices

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

var bingSearchUrl = "https://api.cognitive.microsoft.com/bing/v5.0/search"

type BingSearchResult struct {
	Type            string                          `json:"_type"`
	Instrumentation BingSearchResultInstrumentation `json:"instrumentation"`
	WebPages        WebpageResult                   `json:"webPages"`
	Images          ImageResult                     `json:"images"`
	News            NewsResult                      `json:"news"`
	RelatedSearches RelatedSearchesResult           `json:"relatedSearches"`
	Videos          VideoResult                     `json:"videos"`
}

type BingSearchResultInstrumentation struct {
	PingUrlBase     string `json:"pingUrlBase"`
	PageLoadPingUrl string `json:"pageLoadPingUrl"`
}

type WebpageResult struct {
	WebSearchUrl          string               `json:"webSearchUrl"`
	TotalEstimatedMatches int                  `json:"totalEstimatedMatches"`
	Value                 []WebpageResultValue `json:"value"`
}

type WebpageResultValue struct {
	Id              string                       `json:"id"`
	Name            string                       `json:"name"`
	Url             string                       `json:"url"`
	UrlPingSuffix   string                       `json:"urlPingSuffix"`
	About           []WebpageResultValueAbout    `json:"about"`
	DisplayUrl      string                       `json:"displayUrl"`
	Snippet         string                       `json:"snippet"`
	DeepLinks       []WebpageResultValueDeeplink `json:"deepLinks"`
	DateLastCrawled string                       `json:"dateLastCrawled"`
}

type WebpageResultValueAbout struct {
	Name string `json:"name"`
}

type WebpageResultValueDeeplink struct {
	Name          string `json:"name"`
	Url           string `json:"url"`
	UrlPingSuffix string `json:"urlPingSuffix"`
}

type ImageResult struct {
	Id                           string             `json:"id"`
	ReadLink                     string             `json:"readLink"`
	WebSearchUrl                 string             `json:"webSearchUrl"`
	WebSearchUrlPingSuffix       string             `json:"webSearchUrlPingSuffix"`
	IsFamilyFriendly             bool               `json:"isFamilyFriendly"`
	DisplayShoppingSourcesBadges bool               `json:"displayShoppingSourcesBadges"`
	DisplayRecipeSourcesBadges   bool               `json:"displayRecipeSourcesBadges"`
	Value                        []ImageResultValue `json:"value"`
}

type ImageResultValue struct {
	Name                   string                    `json:"name"`
	WebSearchUrl           string                    `json:"webSearchUrl"`
	WebSearchUrlPingSuffix string                    `json:"webSearchUrlPingSuffix"`
	ThumbnailUrl           string                    `json:"thumbnailUrl"`
	DatePublished          string                    `json:"datePublished"`
	ContentUrl             string                    `json:"contentUrl"`
	HostPageUrl            string                    `json:"hostPageUrl"`
	HostPageUrlPingSuffix  string                    `json:"hostPageUrlPingSuffix"`
	ContentSize            string                    `json:"contentSize"`
	EncodingFormat         string                    `json:"encodingFormat"`
	HostPageDisplayUrl     string                    `json:"hostPageDisplayUrl"`
	Width                  int                       `json:"width"`
	Height                 int                       `json:"height"`
	Thumbnail              ImageResultValueThumbnail `json:"thumbnail"`
}

type ImageResultValueThumbnail struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

type NewsResult struct {
	Id       string            `json:"id"`
	ReadLink string            `json:"readLink"`
	Value    []NewsResultValue `json:"value"`
}

type NewsResultValue struct {
	Name          string                    `json:"name"`
	Url           string                    `json:"url"`
	UrlPingSuffix string                    `json:"urlPingSuffix"`
	Image         NewsResultValueImage      `json:"image"`
	Description   string                    `json:"description"`
	About         []NewsResultValueAbout    `json:"about"`
	Provider      []NewsResultValueProvider `json:"provider"`
	DatePublished string                    `json:"datepublished"`
	Category      string                    `json:"category"`
}

type NewsResultValueImage struct {
	NewsResultValueImageThumbnail struct {
		ContentUrl string `json:"contentUrl"`
		Width      int    `json:"width"`
		Height     int    `json:"height"`
	} `json:"thumbnail"`
}

type NewsResultValueAbout struct {
	ReadLink string `json:"readLink"`
	Name     string `json:"name"`
}

type NewsResultValueProvider struct {
	Type string `json:"_type"`
	Name string `json:"name"`
}

type RelatedSearchesResult struct {
	Id    string                       `json:"id"`
	Value []RelatedSearchesResultValue `json:"value"`
}

type RelatedSearchesResultValue struct {
	Text                   string `json:"text"`
	DisplayText            string `json:"displayText"`
	WebSearchUrl           string `json:"webSearchUrl"`
	WebSearchUrlPingSuffix string `json:"webSearchUrlPingSuffix"`
}

type VideoResult struct {
	Id                     string             `json:"id"`
	ReadLink               string             `json:"readLink"`
	WebSearchUrl           string             `json:"webSearchUrl"`
	WebSearchUrlPingSuffix string             `json:"webSearchUrlPingSuffix"`
	IsFamilyFriendly       bool               `json:"isFamilyFriendly"`
	Value                  []VideoResultValue `json:"value"`
	Scenario               string             `json:"scenario"`
}

type VideoResultValue struct {
	Name                   string                      `json:"name"`
	Description            string                      `json:"description"`
	WebSearchUrl           string                      `json:"webSearchUrl"`
	WebSearchUrlPingSuffix string                      `json:"webSearchUrlPingSuffix"`
	ThumbnailUrl           string                      `json:"thumbnailUrl"`
	DatePublished          string                      `json:"datePublished"`
	Publisher              []VideoResultValuePublisher `json:"publisher"`
	ContentUrl             string                      `json:"contentUrl"`
	HostPageUrl            string                      `json:"hostPageUrl"`
	HostPageUrlPingSuffix  string                      `json:"hostPageUrlPingSuffix"`
	EncodingFormat         string                      `json:"encodingFormat"`
	HostPageDisplayUrl     string                      `json:"hostPageDisplayUrl"`
	Width                  int                         `json:"width"`
	Height                 int                         `json:"height"`
	Duration               string                      `json:"duration"`
	MotionThumbnailUrl     string                      `json:"motionThumbnailUrl"`
	EmbedHtml              string                      `json:"embedHtml"`
	AllowHttpsEmbed        bool                        `json:"allowHttpsEmbed"`
	ViewCount              int                         `json:"viewCount"`
	Thumbnail              VideoResultValueThumbnail   `json:"thumbnail"`
}

type VideoResultValuePublisher struct {
	Name string `json:"name"`
}

type VideoResultValueThumbnail struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

type WebsearchError struct {
	ErrorResponse struct {
		StatusCode int    `json:"statusCode"`
		Message    string `json:"message"`
	} `json:"error"`
}

func (we *WebsearchError) Error() string {
	return we.ErrorResponse.Message
}

func getResponseError(res *http.Response) error {
	var webSearchError *WebsearchError
	err := json.NewDecoder(res.Body).Decode(webSearchError)
	if err != nil {
		return err
	}
	return webSearchError
}

func (cs *CognitiveServicesClient) WebSearch(query string, count int, offset int, mkt string, safesearch string) (*BingSearchResult, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", bingSearchUrl, nil)
	if err != nil {
		log.Fatal(err)
	}
	params := req.URL.Query()
	params.Add("q", query)
	params.Add("count", strconv.Itoa(count))
	params.Add("offset", strconv.Itoa(offset))
	params.Add("mkt", mkt)
	params.Add("safesearch", safesearch)
	req.URL.RawQuery = params.Encode()

	req.Header.Add("Ocp-Apim-Subscription-Key", cs.SubscriptionKey)

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, getResponseError(res)
	}

	var bingResult BingSearchResult
	err = json.NewDecoder(res.Body).Decode(&bingResult)
	if err != nil {
		return nil, err
	}

	return &bingResult, nil
}
