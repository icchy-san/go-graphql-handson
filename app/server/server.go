package main

import (
	"app/db"
	g "app/generated"
	r "app/resolver"
	"github.com/99designs/gqlgen/handler"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Printf("failed to load .env file")
	}

	db.InitDB()
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/graphql", func(c echo.Context) error {
		resolver := &r.Resolver{
			DBClient: db.GetDBConnection(),
		}

		h := handler.GraphQL(g.NewExecutableSchema(g.Config{Resolvers: resolver}))

		h.ServeHTTP(c.Response(), c.Request())
		return nil
	})

	if err := e.Start(":8080"); err != nil {
		log.Fatalf("%s", err)
	}
}
