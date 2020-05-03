package twitter_client

import (
	"github.com/dghubble/go-twitter/twitter"
)

func CreateTwitterTestClient(h TwitterHarness) TwitterTestClient {
	return TwitterTestClient{h}
}

type TwitterTestClient struct {
	harness TwitterHarness
}

func (ttc TwitterTestClient) Search(s string, p *twitter.SearchTweetParams) (*twitter.Search, error) {
	return ttc.harness.Search(s, p)
}

type TwitterHarness struct {
	Search func(s string, p *twitter.SearchTweetParams) (*twitter.Search, error)
}
