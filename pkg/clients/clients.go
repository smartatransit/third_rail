package clients

import (
	"github.com/dghubble/go-twitter/twitter"
	"github.com/smartatransit/gomarta"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
	"os"
)

type MartaClient interface{}
type TwitterClient interface{}


func GetMartaClient() *gomarta.Client {
	return gomarta.NewDefaultClient(os.Getenv("MARTA_API_KEY"))
}

func GetTwitterClient() *twitter.Client {
	config := &clientcredentials.Config{
		ClientID:     os.Getenv("TWITTER_CLIENT_ID"),
		ClientSecret: os.Getenv("TWITTER_CLIENT_SECRET"),
		TokenURL:     "https://api.twitter.com/oauth2/token",
	}

	httpClient := config.Client(oauth2.NoContext)

	client := twitter.NewClient(httpClient)

	return client
}
