package bingsearch

var bingSearchUrl = "https://api.cognitive.microsoft.com/bing/v5.0/search"

type WebService service

type BingSearchResult struct {
	Type            string                `json:"_type"`
	Instrumentation BingInstrumentation   `json:"instrumentation"`
	WebPages        WebpageResult         `json:"webPages"`
	Images          ImageResult           `json:"images"`
	News            NewsResult            `json:"news"`
	RelatedSearches RelatedSearchesResult `json:"relatedSearches"`
	Videos          VideoResult           `json:"videos"`
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
	Value                        []ImageSearchValue `json:"value"`
}

type NewsResult struct {
	Id       string            `json:"id"`
	ReadLink string            `json:"readLink"`
	Value    []NewsSearchValue `json:"value"`
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
	Value                  []VideoSearchValue `json:"value"`
	Scenario               string             `json:"scenario"`
}

func (s *WebService) Search(q string, count int, offset int, mkt string, safeSearch string) (*BingSearchResult, error) {
	var bingResult *BingSearchResult
	params := NewSearchQueryParams(q, count, offset, mkt, safeSearch)
	err := s.client.search(bingSearchUrl, params, &bingResult)
	if err != nil {
		return nil, err
	}
	return bingResult, nil
}
