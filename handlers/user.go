package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"sis/models"
	"sis/repository"
	"strconv"
)

type User struct {
	Repository repository.User
}

func (u User) GetAll(c echo.Context) error {
	careers, _ := u.Repository.GetAll()
	return c.JSON(http.StatusOK, careers)
}

func (u User) GetById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	career, err := u.Repository.GetById(uint(id))

	if err != nil {
		return c.JSON(http.StatusBadRequest, "")
	}

	return c.JSON(http.StatusOK, career)
}

func (u User) Create(c echo.Context) error {
	user := models.User{}

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, "")
	}
	err := u.Repository.Create(user)

	if err != nil {
		return c.JSON(http.StatusBadRequest, "")
	}

	return c.JSON(http.StatusOK, "User Created")
}

func (u User) Update(c echo.Context) error {
	user := models.User{}

	id, _ := strconv.Atoi(c.Param("id"))

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, "")
	}
	err := u.Repository.Update(user, uint(id))

	if err != nil {
		return c.JSON(http.StatusBadRequest, "")
	}

	return c.JSON(http.StatusOK, "User Updated")
}

func (u User) Delete(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	err := u.Repository.Delete(uint(id))

	if err != nil {
		return c.JSON(http.StatusBadRequest, "")
	}

	return c.JSON(http.StatusOK, "User Deleted")
}
