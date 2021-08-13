package routes

import (
	"github.com/labstack/echo/v4"
	"sis/data"
	"sis/handlers"
	"sis/repository"
)

func UserRoutes(group *echo.Group) {

	UserHandlers := handlers.User{Repository: repository.User{Data: data.New()}}
	group.GET("/users", UserHandlers.GetAll)
	group.GET("/users/:id", UserHandlers.GetById)
	group.POST("/users", UserHandlers.Create)
	group.PUT("/users/:id", UserHandlers.Update)
	group.DELETE("/users/:id", UserHandlers.Delete)

}
