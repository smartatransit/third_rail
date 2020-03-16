package twitter_client

import (
	"time"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/karlseguin/ccache"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

type TwitterAPIClient struct {
	client   *twitter.Client
	cache    *ccache.Cache
	cacheTTL time.Duration
}

func GetTwitterClient(clientID string, clientSecret string, cacheTTL int) TwitterAPIClient {

	var cache = ccache.New(ccache.Configure().MaxSize(1000).ItemsToPrune(100))

	config := &clientcredentials.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		TokenURL:     "https://api.twitter.com/oauth2/token",
	}

	httpClient := config.Client(oauth2.NoContext)

	client := twitter.NewClient(httpClient)

	return TwitterAPIClient{client: client, cache: cache, cacheTTL: time.Duration(cacheTTL)}
}

func (t TwitterAPIClient) Search(searchKey string, search *twitter.SearchTweetParams) (*twitter.Search, error) {
	tweets, err := t.cache.Fetch(searchKey, time.Second*t.cacheTTL, func() (interface{}, error) {
		result, _, err := t.client.Search.Tweets(search)
		return result, err
	})

	if err != nil {
		return nil, err
	}

	return tweets.Value().(*twitter.Search), nil
}
