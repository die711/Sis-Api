package routes

import (
	"github.com/labstack/echo/v4"
	"sis/data"
	"sis/handlers"
	"sis/repository"
)

func CourseRoutes(group *echo.Group) {

	courseHandlers := handlers.Course{Repository: repository.Course{Data: data.New()}}
	group.GET("/courses", courseHandlers.GetAll)
	group.GET("/courses/:id", courseHandlers.GetById)
	group.POST("/courses", courseHandlers.Create)
	group.PUT("/courses/:id", courseHandlers.Update)
	group.DELETE("/courses/:id", courseHandlers.Delete)

}
