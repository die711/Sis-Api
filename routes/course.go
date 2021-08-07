package routes

import (
	"github.com/labstack/echo/v4"
	"sis/data"
	"sis/handlers"
	"sis/repository"
)

func CourseRoutes(group *echo.Group){

	courseHandlers := handlers.Course{Repository: repository.Course{Data: data.New()}}
	group.GET("/courses", courseHandlers.GetAll)
	group.GET("/courses/:id", courseHandlers.GetById)
	//group.POST("/careers", careerHandlers.Create)
	//group.PUT("/careers/:id", careerHandlers.Update)
	//group.DELETE("/careers/:id", careerHandlers.Delete)

}

