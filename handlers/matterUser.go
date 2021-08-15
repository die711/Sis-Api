package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"sis/repository"
)

type MatterUser struct {
	Repository repository.MatterUser
}

func (mu MatterUser) GetAll(c echo.Context) error {
	careers, _ := mu.Repository.GetAll()
	return c.JSON(http.StatusOK, careers)
}

//func (mu MatterUser) GetById(c echo.Context) error {
//	id, _ := strconv.Atoi(c.Param("id"))
//
//	career, err := mu.Repository.GetById(uint(id))
//
//	if err != nil {
//		return c.JSON(http.StatusBadRequest, "")
//	}
//
//	return c.JSON(http.StatusOK, career)
//}
//
//func (mu MatterUser) Create(c echo.Context) error {
//	career := models.Career{}
//
//	if err := c.Bind(&career); err != nil {
//		return c.JSON(http.StatusBadRequest, "")
//	}
//	err := mu.Repository.Create(career)
//
//	if err != nil {
//		return c.JSON(http.StatusBadRequest, "")
//	}
//
//	return c.JSON(http.StatusOK, "Career Created")
//}
//
//func (mu MatterUser) Update(c echo.Context) error {
//	career := models.Career{}
//
//	err := c.Bind(&career)
//
//	if err != nil {
//		return c.JSON(http.StatusBadRequest, "")
//	}
//
//	id, err := strconv.Atoi(c.Param("id"))
//
//	err = mu.Repository.Update(career, uint(id))
//
//	if err != nil {
//		return c.JSON(http.StatusBadRequest, "")
//	}
//
//	return c.JSON(http.StatusOK, "Career Updated")
//
//}
//
//func (mu MatterUser) Delete(c echo.Context) error {
//	id, _ := strconv.Atoi(c.Param("id"))
//
//	err := mu.Repository.Delete(uint(id))
//
//	if err != nil {
//		return c.JSON(http.StatusBadRequest, "")
//	}
//
//	return c.JSON(http.StatusOK, "Career Deleted")
//}
