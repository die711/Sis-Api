package routes

import (
	"github.com/labstack/echo/v4"
	"sis/data"
	"sis/handlers"
	"sis/repository"
)

func MatterUserRoutes(group *echo.Group) {
	matterUserHandlers := handlers.MatterUser{Repository: repository.MatterUser{Data: data.New()}}
	group.GET("/matterUsers", matterUserHandlers.GetAll)
	group.GET("/matterUsers/:userId/:matterId", matterUserHandlers.GetById)
	//group.POST("/matterUsers", matterUserHandlers.Create)
	//group.PUT("/matterUsers/:id", matterUserHandlers.Update)
	//group.DELETE("/matterUsers/:id", matterUserHandlers.Delete)
}
