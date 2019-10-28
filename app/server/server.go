package main

import (
	h "app/handler"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/graphql", h.GraphQL())

	if err := e.Start(":3000"); err != nil {
		log.Fatalf("ERROR: %s", err)
	}

}
