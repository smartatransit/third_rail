package twitter_client

import (
	"github.com/dghubble/go-twitter/twitter"
	"github.com/karlseguin/ccache"
	log "github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
	"time"
)

type TwitterAPIClient struct {
	client   *twitter.Client
	cache    *ccache.Cache
	cacheTTL time.Duration
}

func GetTwitterClient(clientId, clientSecret string, cacheTTL int) TwitterAPIClient {
	if len(clientId) < 1 || len(clientSecret) < 1 {
		log.Fatal("Twitter API credentials not found - Twitter alerts are unavailable.")
	}

	if cacheTTL < 1 {
		log.Warn("Twitter caching set to < 1ms - Twitter results are not being cached and will be rate-limited.")
	}

	var cache = ccache.New(ccache.Configure().MaxSize(1000).ItemsToPrune(100))

	config := &clientcredentials.Config{
		ClientID:     clientId,
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
