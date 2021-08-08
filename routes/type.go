package routes

import (
	"github.com/labstack/echo/v4"
	"sis/data"
	"sis/handlers"
	"sis/repository"
)

func TypeRoutes(group *echo.Group) {

	TypeHandlers := handlers.Type{Repository: repository.Type{Data: data.New()}}
	group.GET("/types", TypeHandlers.GetAll)
	group.GET("/types/:id", TypeHandlers.GetById)
	group.POST("/types", TypeHandlers.Create)
	group.PUT("/types/:id", TypeHandlers.Update)
	group.DELETE("/types/:id", TypeHandlers.Delete)

}
