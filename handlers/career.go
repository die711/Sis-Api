package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"sis/models"
	"sis/repository"
	"strconv"
)

type Career struct {
	Repository repository.Career
}

func (ca Career) GetAll(c echo.Context) error {
	careers, _ := ca.Repository.GetAll()
	return c.JSON(http.StatusOK, careers)
}

func (ca Career) GetById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	career, err := ca.Repository.GetById(uint(id))

	if err != nil {
		return c.JSON(http.StatusBadRequest, "")
	}

	return c.JSON(http.StatusOK, career)
}

func (ca Career) Create(c echo.Context) error {
	career := models.Career{}

	if err := c.Bind(&career); err != nil {
		return c.JSON(http.StatusBadRequest, "")
	}
	err := ca.Repository.Create(career)

	if err != nil {
		return c.JSON(http.StatusBadRequest, "")
	}

	return c.JSON(http.StatusOK, "Career Created")
}

func (ca Career) Update(c echo.Context) error {
	career := models.Career{}

	err := c.Bind(&career)

	if err != nil {
		return c.JSON(http.StatusBadRequest, "")
	}

	id, err := strconv.Atoi(c.Param("id"))

	err = ca.Repository.Update(career, uint(id))

	if err != nil {
		return c.JSON(http.StatusBadRequest, "")
	}

	return c.JSON(http.StatusOK, "Career Updated")

}

func (ca Career) Delete(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	err := ca.Repository.Delete(uint(id))

	if err != nil {
		return c.JSON(http.StatusBadRequest, "")
	}

	return c.JSON(http.StatusOK, "Career Deleted")
}
