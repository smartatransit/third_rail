package clients

import (
	"github.com/dghubble/go-twitter/twitter"
	"github.com/karlseguin/ccache"
	"github.com/smartatransit/gomarta"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
	"google.golang.org/appengine/log"
	"os"
	"strconv"
	"time"
)

type MartaClient interface {
	GetTrains() (gomarta.TrainAPIResponse, error)
}

type TwitterClient interface {
	Search(string, *twitter.SearchTweetParams) (*twitter.Search, error)
}

type TwitterAPIClient struct {
	client *twitter.Client
	cache  *ccache.Cache
	cacheTTL int
}

type MartaAPIClient struct {
	client *gomarta.Client
	cache  *ccache.Cache
	cacheTTL int
}

func GetMartaClient() MartaAPIClient {
	var cache = ccache.New(ccache.Configure().MaxSize(1000).ItemsToPrune(100))
	var marta = gomarta.NewDefaultClient(os.Getenv("MARTA_API_KEY"))
	cacheTTL, err := strconv.Atoi(os.Getenv("MARTA_CACHE_TTL"))

	if err != nil {
		cacheTTL = 15
	}

	return MartaAPIClient{client: marta, cache: cache, cacheTTL:cacheTTL}
}

func (m MartaAPIClient) GetTrains() (gomarta.TrainAPIResponse, error) {

	trains, err := m.cache.Fetch("trains", time.Second*m.cacheTTL, func() (interface{}, error) {
		return m.client.GetTrains()
	})

	if err != nil {
		return nil, err
	}

	return trains.Value().(gomarta.TrainAPIResponse), nil
}

func GetTwitterClient() TwitterAPIClient {
	var cache = ccache.New(ccache.Configure().MaxSize(1000).ItemsToPrune(100))

	config := &clientcredentials.Config{
		ClientID:     os.Getenv("TWITTER_CLIENT_ID"),
		ClientSecret: os.Getenv("TWITTER_CLIENT_SECRET"),
		TokenURL:     "https://api.twitter.com/oauth2/token",
	}

	httpClient := config.Client(oauth2.NoContext)

	client := twitter.NewClient(httpClient)
	cacheTTL, err := strconv.Atoi(os.Getenv("MARTA_CACHE_TTL"))

	if err != nil {
		cacheTTL = 15
	}


	return TwitterAPIClient{client: client, cache: cache, cacheTTL:cacheTTL}
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
