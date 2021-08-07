package routes

import (
	"github.com/labstack/echo/v4"
	"sis/data"
	"sis/handlers"
	"sis/repository"
)

func CareerRoutes(group *echo.Group){

	careerHandlers := handlers.Career{Repository: repository.Career{Data: data.New()}}
	group.GET("/careers", careerHandlers.GetAll)
	group.GET("/careers/:id", careerHandlers.GetById)
	group.POST("/careers", careerHandlers.Create)
	group.PUT("/careers/:id", careerHandlers.Update)
	group.DELETE("/careers/:id", careerHandlers.Delete)

}
