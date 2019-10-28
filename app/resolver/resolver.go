package resolver

//go:generate go run github.com/99designs/gqlgen

import (
	g "app/generated"
	"github.com/jinzhu/gorm"
)

type (
	Resolver struct {
		DBClient *gorm.DB
	}
	queryResolver struct{ *Resolver }
	userResolver  struct{ *Resolver }
	tweetResolver struct{ *Resolver }
)

func (r *Resolver) Query() g.QueryResolver {
	return &queryResolver{r}
}
func (r *Resolver) User() g.UserResolver {
	return &userResolver{r}
}
