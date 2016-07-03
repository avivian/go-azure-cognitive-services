package bingsearch

type BingInstrumentation struct {
	PingUrlBase     string `json:"pingUrlBase"`
	PageLoadPingUrl string `json:"pageLoadPingUrl"`
}

type PivotSuggestionsList struct {
	Pivot       string             `json:"pivot"`
	Suggestions []PivotSuggestions `json:"suggestions"`
}

type PivotSuggestions struct {
	Text                   string               `json:"text"`
	DisplayText            string               `json:"displayText"`
	WebSearchUrl           string               `json:"webSearchUrl"`
	WebSearchUrlPingSuffix string               `json:"webSearchUrlPingSuffix"`
	SearchLink             string               `json:"searchLink"`
	Thumbnail              ImageSearchThumbnail `json:"thumbnail"`
}

type ImageSearchThumbnail struct {
	ThumbnailUrl string `json:"thumbnailUrl"`
}

type QueryExpansions struct {
	DisplayText            string               `json:"displayText"`
	WebSearchUrl           string               `json:"webSearchUrl"`
	WebSearchUrlPingSuffix string               `json:"webSearchUrlPingSuffix"`
	SearchLink             string               `json:"searchLink"`
	Thumbnail              ImageSearchThumbnail `json:"thumbnail"`
}

type ImageThumbnail struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}
