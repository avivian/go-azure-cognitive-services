# Azure Cognitive Services

Go Client Library for accessing [Microsoft's Cognitive Services API](https://www.microsoft.com/cognitive-services).

## Services Available
* Bing Web Search API

## Usage
```go
import "github.com/avivian/go-azure-cognitive-services/cognitiveservices"
```

You will need to sign up to the Cognitive Services API via Azure and obtain a subscription key before you can use the client. 

Once you have your subscription key, construct a new Cognitive Services client and then use the client to access the various services on the Cognitive Services API.

```
client := cognitiveservices.NewClient("your-subscription-key")
results, err := client.WebSearch("cats", 10, 0, "GB-en", "Moderate")
```

## License
MIT
