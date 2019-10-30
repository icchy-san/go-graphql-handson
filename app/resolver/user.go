package resolver

import (
	m "app/model"
	"context"
)

func (r *userResolver) Identifier(ctx context.Context, obj *m.User) (string, error) {
	panic("not implemented")
}

func (r *mutationResolver) CreateUser(ctx context.Context, input m.CreateUserInput) (*m.User, error) {
	var user m.User
	if err := user.Create(r.DBClient, input); err != nil {
		return &m.User{}, err
	}

	return &m.User{}, nil
}

func (r *queryResolver) User(ctx context.Context, input m.UserInput) (*m.User, error) {
	var user m.User
	if err := user.FindByIdentifier(r.DBClient, input.Identifier); err != nil {
		return &m.User{}, err
	}

	return &user, nil
}
