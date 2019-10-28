package app

import (
	"context"

	"github.com/ito-lvgs/go-graphql-handson/generated"
	"github.com/ito-lvgs/go-graphql-handson/model"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{}

func (r *Resolver) Query() generated.QueryResolver {
	return &queryResolver{r}
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Tweet(ctx context.Context) (*model.Tweet, error) {
	panic("not implemented")
}
