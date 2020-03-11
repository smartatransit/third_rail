package twitter_client

import (
	"time"

	"os"
	"strconv"
	"time"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/karlseguin/ccache"
	"github.com/smartatransit/gomarta"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

type TwitterAPIClient struct {
	client   *twitter.Client
	cache    *ccache.Cache
	cacheTTL int
}

type MartaAPIClient struct {
	client   *gomarta.Client
	cache    *ccache.Cache
	cacheTTL int
}

func GetMartaClient(apiKey string, cacheTTL int) MartaAPIClient {
	var cache = ccache.New(ccache.Configure().MaxSize(1000).ItemsToPrune(100))
	var marta = gomarta.NewDefaultClient(apiKey)

	return MartaAPIClient{client: marta, cache: cache, cacheTTL: cacheTTL}
}

func (m MartaAPIClient) GetTrains() (gomarta.TrainAPIResponse, error) {

	trains, err := m.cache.Fetch("trains", time.Second*time.Duration(m.cacheTTL), func() (interface{}, error) {
		return m.client.GetTrains()
	})

	if err != nil {
		return nil, err
	}

	return trains.Value().(gomarta.TrainAPIResponse), nil
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

	return TwitterAPIClient{client: client, cache: cache, cacheTTL: cacheTTL}
}

func (t TwitterAPIClient) Search(searchKey string, search *twitter.SearchTweetParams) (*twitter.Search, error) {
	tweets, err := t.cache.Fetch(searchKey, time.Second*time.Duration(t.cacheTTL), func() (interface{}, error) {
		result, _, err := t.client.Search.Tweets(search)
		return result, err
	})

	if err != nil {
		return nil, err
	}

	return tweets.Value().(*twitter.Search), nil
}
