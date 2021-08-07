package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"sis/repository"
	"strconv"
)

type Course struct {
	Repository repository.Course
}

func (co Course) GetAll(c echo.Context) error {
	careers, _ := co.Repository.GetAll()
	return c.JSON(http.StatusOK, careers)
}

func (co Course) GetById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	career, err := co.Repository.GetById(uint(id))

	if err != nil {
		return c.JSON(http.StatusBadRequest, "")
	}

	return c.JSON(http.StatusOK, career)
}


