package routes

import (
	"github.com/labstack/echo/v4"
	"sis/data"
	"sis/handlers"
	"sis/repository"
)

func MatterRoutes(group *echo.Group) {

	matterHandlers := handlers.Matter{Repository: repository.Matter{Data: data.New()}}
	group.GET("/matters", matterHandlers.GetAll)
	group.GET("/matters/:id", matterHandlers.GetById)
	group.POST("/matters", matterHandlers.Create)
	group.PUT("/matters/:id", matterHandlers.Update)
	group.DELETE("/matters/:id", matterHandlers.Delete)

}
