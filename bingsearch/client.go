package bingsearch

type BingSearchClient struct {
	SubscriptionKey string
}

func NewClient(subscriptionKey string) BingSearchClient {
	return BingSearchClient{
		SubscriptionKey: subscriptionKey,
	}
}
