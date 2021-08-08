package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"sis/models"
	"sis/repository"
	"strconv"
)

type Type struct {
	Repository repository.Type
}

func (t Type) GetAll(c echo.Context) error {
	careers, _ := t.Repository.GetAll()
	return c.JSON(http.StatusOK, careers)
}

func (t Type) GetById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	career, err := t.Repository.GetById(uint(id))

	if err != nil {
		return c.JSON(http.StatusBadRequest, "")
	}

	return c.JSON(http.StatusOK, career)
}

func (t Type) Create(c echo.Context) error {
	type_ := models.Type{}

	if err := c.Bind(&type_); err != nil {
		return c.JSON(http.StatusBadRequest, "")
	}
	err := t.Repository.Create(type_)

	if err != nil {
		return c.JSON(http.StatusBadRequest, "")
	}

	return c.JSON(http.StatusOK, "Type Created")
}

func (t Type) Update(c echo.Context) error {
	type_ := models.Type{}

	err := c.Bind(&type_)

	if err != nil {
		return c.JSON(http.StatusBadRequest, "")
	}

	id, err := strconv.Atoi(c.Param("id"))

	err = t.Repository.Update(type_, uint(id))

	if err != nil {
		return c.JSON(http.StatusBadRequest, "")
	}

	return c.JSON(http.StatusOK, "Type Updated")

}

func (t Type) Delete(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	err := t.Repository.Delete(uint(id))

	if err != nil {
		return c.JSON(http.StatusBadRequest, "")
	}

	return c.JSON(http.StatusOK, "Type Deleted")
}
