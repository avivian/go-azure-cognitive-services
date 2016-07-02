package bingsearch

var bingNewsSearchUrl = "https://api.cognitive.microsoft.com/bing/v5.0/news/search"

type BingNewsSearchResult struct {
	Type                  string              `json:"_type"`
	Instrumentation       BingInstrumentation `json:"instrumentation"`
	ReadLink              string              `json:"readLink"`
	TotalEstimatedMatches int                 `json:"totalEstimatedMatches"`
	Value                 []NewsSearchValue   `json:"value"`
}

type NewsSearchValue struct {
	Name          string         `json:"name"`
	Url           string         `json:"url"`
	UrlPingSuffix string         `json:"urlPingSuffix"`
	Image         NewsImage      `json:"image"`
	Description   string         `json:"description"`
	About         []NewsAbout    `json:"about"`
	Provider      []NewsProvider `json:"provider"`
	DatePublished string         `json:"datePublished"`
	Category      string         `json:"category"`
}

type NewsImage struct {
	Thumbnail NewsImageThumbnail `json:"thumbnail"`
}

type NewsImageThumbnail struct {
	ContentUrl string `json:"contentUrl"`
	Width      int    `json:"width"`
	Height     int    `json:"height"`
}

type NewsAbout struct {
	ReadLink string `json:"readLink"`
	Name     string `json:"name"`
}

type NewsProvider struct {
	Type string `json:"_type"`
	Name string `json:"name"`
}

func (bsc *BingSearchClient) NewsSearch(q string, count int, offset int, mkt string, safeSearch string) (*BingNewsSearchResult, error) {
	var result *BingNewsSearchResult
	params := NewSearchQueryParams(q, count, offset, mkt, safeSearch)
	err := bsc.searchRequest(bingNewsSearchUrl, params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
