package twitter_client

import (
	"github.com/dghubble/go-twitter/twitter"
)

type TwitterTestClient struct {
	harness TwitterHarness
}

func (ttc TwitterTestClient) Search(s string, p *twitter.SearchTweetParams) (*twitter.Search, error) {
	return ttc.harness.search(s, p)
}

type TwitterHarness struct {
	search func(s string, p *twitter.SearchTweetParams) (*twitter.Search, error)
}
