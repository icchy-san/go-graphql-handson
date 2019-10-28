package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/graphql", func(c echo.Context) error {
		return nil
	})

	if err := e.Start(":8080"); err != nil {
		log.Fatalf("ERROR: %s", err)
	}
}
