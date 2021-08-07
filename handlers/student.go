package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"sis/data"
	"sis/repository"
)

//var users = models.AllUsers{
//	{
//		NoControl: 15490905,
//		Name:      "Diego",
//		LastName:  "Resendiz",
//		Career:    "Ing Sistema Computacionales",
//		Semester:  4,
//	},
//}
var studentRepository = repository.StudentRepository{Data: data.New()}

func GetStudents(c echo.Context) error {
	students, _ := studentRepository.GetAll()
	return c.JSON(http.StatusOK, students)
}

//func GetUser(c echo.Context) error {
//	noControl, _ := strconv.Atoi(c.Param("noControl"))
//
//	for _, user := range users {
//		if user.NoControl == noControl {
//			return c.JSON(http.StatusOK, user)
//		}
//	}
//
//	return c.JSON(http.StatusBadRequest, "")
//}
//
//func CreateUser(c echo.Context) error {
//	user := models.User{}
//
//	if err := c.Bind(&user); err != nil {
//		return err
//	}
//	users = append(users, user)
//
//	return c.JSON(http.StatusOK, "User created")
//}
//
//func UpdateUser(c echo.Context) error {
//	user := models.User{}
//
//	if err := c.Bind(&user); err != nil {
//		return err
//	}
//	noControl, _ := strconv.Atoi(c.Param("noControl"))
//
//	for i, u := range users {
//		if u.NoControl == noControl {
//			users[i]= user
//			return c.JSON(http.StatusOK, "User updated")
//		}
//	}
//
//	return c.JSON(http.StatusBadRequest, "")
//}
//
//func DeleteUser(c echo.Context) error {
//	noControl, _ := strconv.Atoi(c.Param("noControl"))
//
//	for i, user := range users {
//		if user.NoControl == noControl {
//			users = append(users[:i], users[i+1:]...)
//			return c.JSON(http.StatusOK, "User deleted")
//		}
//	}
//
//	return c.JSON(http.StatusBadRequest, "")
//}
