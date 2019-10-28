package go_graphql_handson

import (
	"app/generated"
	"app/model"
	"context"
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
