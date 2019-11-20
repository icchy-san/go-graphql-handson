package resolver

import (
	m "app/model"
	"context"
)

func (r *queryResolver) Tweet(ctx context.Context, input m.TweetInput) (*m.Tweet, error) {
	panic("Implement here")
}

func (r *tweetResolver) User(ctx context.Context, obj *m.Tweet) (*m.User, error) {
	panic("Implement here")
}

func (r *tweetResolver) CreatedAt(ctx context.Context, obj *m.Tweet) (string, error) {
	panic("Implement here")
}

func (r *mutationResolver) CreateTweet(ctx context.Context, input m.CreateTweetInput) (*m.Tweet, error) {
	// TODO: 作成したユーザを取得できるようにする
	user := m.User {
		ID: 1,
		Identifier: "dklajf",
		Name: "icchy",
	}

	var tweet m.Tweet
	if err := tweet.Create(r.DBClient, input, &user); err != nil {
		return &m.Tweet{}, err
	}

	return &tweet, nil
}
