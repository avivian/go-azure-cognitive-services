package cognitiveservices

type CognitiveServicesClient struct {
	SubscriptionKey string
}

func NewClient(subscriptionKey string) CognitiveServicesClient {
	return CognitiveServicesClient{
		SubscriptionKey: subscriptionKey,
	}
}
