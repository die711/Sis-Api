package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"sis/models"
	"sis/repository"
	"strconv"
)

type Matter struct {
	Repository repository.Matter
}

func (m Matter) GetAll(c echo.Context) error {
	careers, _ := m.Repository.GetAll()
	return c.JSON(http.StatusOK, careers)
}

func (m Matter) GetById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	matter, err := m.Repository.GetById(uint(id))

	if err != nil {
		return c.JSON(http.StatusBadRequest, "")
	}

	return c.JSON(http.StatusOK, matter)
}

//
func (m Matter) Create(c echo.Context) error {
	matter := models.Matter{}

	if err := c.Bind(&matter); err != nil {
		return c.JSON(http.StatusBadRequest, "")
	}
	err := m.Repository.Create(matter)

	if err != nil {
		return c.JSON(http.StatusBadRequest, "")
	}

	return c.JSON(http.StatusOK, "Matter Created")
}

func (m Matter) Update(c echo.Context) error {
	matter := models.Matter{}

	err := c.Bind(&matter)

	if err != nil {
		return c.JSON(http.StatusBadRequest, "")
	}

	id, err := strconv.Atoi(c.Param("id"))

	err = m.Repository.Update(matter, uint(id))

	if err != nil {
		return c.JSON(http.StatusBadRequest, "")
	}

	return c.JSON(http.StatusOK, "Matter Updated")

}

func (m Matter) Delete(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	err := m.Repository.Delete(uint(id))

	if err != nil {
		return c.JSON(http.StatusBadRequest, "")
	}

	return c.JSON(http.StatusOK, "Matter Deleted")
}
