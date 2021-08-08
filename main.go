package main

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
	"sis/data"
	"sis/routes"
)

func main() {

	e := echo.New()
	apiGroup := e.Group("/api")
	routes.CareerRoutes(apiGroup)
	routes.CourseRoutes(apiGroup)
	routes.TypeRoutes(apiGroup)

	d := data.New()
	if err := d.DB.Ping(); err != nil {
		e.Logger.Fatal(err)
	}

	e.Logger.Fatal(e.Start(":80"))
}
