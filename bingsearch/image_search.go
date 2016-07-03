package bingsearch

var bingImageSearchURL = "https://api.cognitive.microsoft.com/bing/v5.0/images/search"

type ImageService service

type BingImageSearchResult struct {
	Type                         string                 `json:"_type"`
	Instrumentation              BingInstrumentation    `json:"instrumentation"`
	ReadLink                     string                 `json:"readLink"`
	WebSearchUrl                 string                 `json:"webSearchUrl"`
	WebSearchUrlPingSuffix       string                 `json:"webSearchUrlPingSuffix"`
	TotalEstimatedMatches        int                    `json:"totalEstimatedMatches"`
	Value                        []ImageSearchValue     `json:"value"`
	QueryExpansions              []QueryExpansions      `json:"queryExpansions"`
	NextOffsetAddCount           int                    `json:"nextOffsetAddCount"`
	PivotSuggestions             []PivotSuggestionsList `json:"pivotSuggestions"`
	DisplayShoppingSourcesBadges bool                   `json:"displayShoppingSourcesBadges"`
	DisplayRecipeSourcesBadges   bool                   `json:"displayRecipeSourcesBadges"`
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
	Thumbnail              ImageThumbnail                         `json:"thumbnail"`
	ImageInsightsToken     string                                 `json:"imageInsightsToken"`
	InsightsSourcesSummary ImageSearchValueInsightsSourcesSummary `json:"insightsSourcesSummary"`
	ImageId                string                                 `json:"imageId"`
	AccentColor            string                                 `json:"accentColor"`
}

type ImageSearchValueInsightsSourcesSummary struct {
	ShoppingSourcesCount int `json:"shoppingSourcesCount"`
	RecipeSourcesCount   int `json:"recipeSourcesCount"`
}

func (s *ImageService) Search(q string, count int, offset int, mkt string, safeSearch string) (*BingImageSearchResult, error) {
	var result *BingImageSearchResult
	params := NewSearchQueryParams(q, count, offset, mkt, safeSearch)
	err := s.client.search(bingImageSearchURL, params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
