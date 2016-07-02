package bingsearch

var bingImageSearchURL = "https://api.cognitive.microsoft.com/bing/v5.0/images/search"

type BingImageSearchResult struct {
	Type                         string                            `json:"_type"`
	Instrumentation              BingInstrumentation               `json:"instrumentation"`
	ReadLink                     string                            `json:"readLink"`
	WebSearchUrl                 string                            `json:"webSearchUrl"`
	WebSearchUrlPingSuffix       string                            `json:"webSearchUrlPingSuffix"`
	TotalEstimatedMatches        int                               `json:"totalEstimatedMatches"`
	Value                        []ImageSearchValue                `json:"value"`
	QueryExpansions              []QueryExpansions                 `json:"queryExpansions"`
	NextOffsetAddCount           int                               `json:"nextOffsetAddCount"`
	PivotSuggestions             []ImageSearchPivotSuggestionsList `json:"pivotSuggestions"`
	DisplayShoppingSourcesBadges bool                              `json:"displayShoppingSourcesBadges"`
	DisplayRecipeSourcesBadges   bool                              `json:"displayRecipeSourcesBadges"`
}

type ImageSearchValue struct {
	Name                   string                                 `json:"name"`
	WebSearchUrl           string                                 `json:"webSearchUrl"`
	WebSearchUrlPingSuffix string                                 `json:"webSearchUrlPingSuffix"`
	ThumbnailUrl           string                                 `json:"thumbnailUrl"`
	DatePublished          string                                 `json:"datePublished"`
	ContentUrl             string                                 `json:"contentUrl"`
	HostPageUrl            string                                 `json:"hostPageUrl"`
	HostPageUrlPingSuffix  string                                 `json:"hostPageUrlPingSuffix"`
	ContentSize            string                                 `json:"contentSize"`
	EncodingFormat         string                                 `json:"encodingFormat"`
	HostPageDisplayUrl     string                                 `json:"hostPageDisplayUrl"`
	Width                  int                                    `json:"width"`
	Height                 int                                    `json:"height"`
	Thumbnail              ImageSearchValueThumbnail              `json:"thumbnail"`
	ImageInsightsToken     string                                 `json:"imageInsightsToken"`
	InsightsSourcesSummary ImageSearchValueInsightsSourcesSummary `json:"insightsSourcesSummary"`
	ImageId                string                                 `json:"imageId"`
	AccentColor            string                                 `json:"accentColor"`
}

type QueryExpansions struct {
	DisplayText            string               `json:"displayText"`
	WebSearchUrl           string               `json:"webSearchUrl"`
	WebSearchUrlPingSuffix string               `json:"webSearchUrlPingSuffix"`
	SearchLink             string               `json:"searchLink"`
	Thumbnail              ImageSearchThumbnail `json:"thumbnail"`
}

type ImageSearchValueThumbnail struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

type ImageSearchValueInsightsSourcesSummary struct {
	ShoppingSourcesCount int `json:"shoppingSourcesCount"`
	RecipeSourcesCount   int `json:"recipeSourcesCount"`
}

type ImageSearchPivotSuggestionsList struct {
	Pivot       string                        `json:"pivot"`
	Suggestions []ImageSearchPivotSuggestions `json:"suggestions"`
}

type ImageSearchPivotSuggestions struct {
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

func (bsc *BingSearchClient) ImageSearch(q string, count int, offset int, mkt string, safeSearch string) (*BingImageSearchResult, error) {
	var result *BingImageSearchResult
	params := NewSearchQueryParams(q, count, offset, mkt, safeSearch)
	err := bsc.searchRequest(bingImageSearchURL, params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
