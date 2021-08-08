package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"sis/models"
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

func (co Course) Create(c echo.Context) error {
	course := models.Course{}

	if err := c.Bind(&course); err != nil {
		return c.JSON(http.StatusBadRequest, "")
	}
	err := co.Repository.Create(course)

	if err != nil {
		return c.JSON(http.StatusBadRequest, "")
	}

	return c.JSON(http.StatusOK, "Course Created")
}

func (co Course) Update(c echo.Context) error {
	course := models.Course{}

	id, _ := strconv.Atoi(c.Param("id"))

	if err := c.Bind(&course); err != nil {
		return c.JSON(http.StatusBadRequest, "")
	}
	err := co.Repository.Update(course, uint(id))

	if err != nil {
		return c.JSON(http.StatusBadRequest, "")
	}

	return c.JSON(http.StatusOK, "Course Updated")
}

func (co Course) Delete(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	err := co.Repository.Delete(uint(id))

	if err != nil {
		return c.JSON(http.StatusBadRequest, "")
	}

	return c.JSON(http.StatusOK, "Course Deleted")
}
