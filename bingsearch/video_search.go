package bingsearch

var bingVideoSearchUrl = "https://api.cognitive.microsoft.com/bing/v5.0/videos/search"

type BingVideoSearchResult struct {
	Type                   string                 `json:"_type"`
	Instrumentation        BingInstrumentation    `json:"instrumentation"`
	ReadLink               string                 `json:"readLink"`
	WebSearchUrl           string                 `json:"webSearchUrl"`
	WebSearchUrlPingSuffix string                 `json:"webSearchUrlPingSuffix"`
	TotalEstimatedMatches  int                    `json:"totalEstimatedMatches"`
	Value                  []VideoSearchValue     `json:"value"`
	QueryExpansions        []QueryExpansions      `json:"queryExpansions"`
	NextOffsetAddCount     int                    `json:"nextOffsetAddCount"`
	PivotSuggestions       []PivotSuggestionsList `json:"pivotSuggestions"`
}

type VideoSearchValue struct {
	Name                   string           `json:"name"`
	Description            string           `json:"description"`
	WebSearchUrl           string           `json:"webSearchUrl"`
	WebSearchUrlPingSuffix string           `json:"webSearchUrlPingSuffix"`
	ThumbnailUrl           string           `json:"thumbnailUrl"`
	DatePublished          string           `json:"datePublished"`
	Publisher              []VideoPublisher `json:"publisher"`
	Creator                VideoCreator     `json:"creator"`
	ContentUrl             string           `json:"contentUrl"`
	HostPageUrl            string           `json:"hostPageUrl"`
	HostPageUrlPingSuffix  string           `json:"hostPageUrlPingSuffix"`
	EncodingFormat         string           `json:"encodingFormat"`
	HostPageDisplayUrl     string           `json:"hostPageDisplayUrl"`
	Width                  int              `json:"width"`
	Height                 int              `json:"height"`
	Duration               string           `json:"duration"`
	MotionThumbnailUrl     string           `json:"motionThumbnailUrl"`
	EmbedHtml              string           `json:"embedHtml"`
	AllowHttpsEmbed        bool             `json:"allowHttpsEmbed"`
	ViewCount              int64            `json:"viewCount"`
	Thumbnail              ImageThumbnail   `json:"thumbnail"`
	VideoId                string           `json:"videoId"`
}

type VideoPublisher struct {
	Name string `json:"name"`
}

type VideoCreator struct {
	Name string `json:"name"`
}

func (bsc *BingSearchClient) VideoSearch(q string, count int, offset int, mkt string, safeSearch string) (*BingVideoSearchResult, error) {
	var result *BingVideoSearchResult
	params := NewSearchQueryParams(q, count, offset, mkt, safeSearch)
	err := bsc.searchRequest(bingVideoSearchUrl, params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
