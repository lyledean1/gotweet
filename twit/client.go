package twit

import (
	"github.com/dghubble/go-twitter/twitter"
)

type TwitterClient struct {
	client *twitter.Client
}

func NewTwitterClient(client *twitter.Client) *TwitterClient {
	return &TwitterClient{client: client}
}

func (tc *TwitterClient) UpdateStatus(tweet string, updateParams *twitter.StatusUpdateParams) error {
	_, _, err := tc.client.Statuses.Update(tweet, updateParams)
	return err
}

func (tc *TwitterClient) Search(searchParams *twitter.SearchTweetParams) (*twitter.Search, error) {
	search, _, err := tc.client.Search.Tweets(searchParams)
	if err != nil {
		return nil, err
	}
	return search, err
}