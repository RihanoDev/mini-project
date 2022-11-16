package controller

import (
	"net/http"
	"rest/config"
	"rest/model"
	"strconv"

	"github.com/labstack/echo"
)

func UserController(e echo.Context) error {

	id, _ := strconv.Atoi(e.Param("id"))

	age, _ := strconv.Atoi(e.Param("age"))

	search := e.QueryParam("search")

	sort := e.QueryParam("sort")

	user := model.User{Id: id, Age: age, Email: "rhr2109@gmail.com", Name: "Rizky"}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"user":   user,
		"search": search,
		"sort":   sort,
	})
}

func CreateUserController(c echo.Context) error {
	//binding data
	user := model.User{}
	c.Bind(&user)

	err := config.DB.Save(&user).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success create user",
		"user":     user,
	})
}

func GetUserController(c echo.Context) error {
	var users []model.User

	err := config.DB.Find(&users).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"data":    users,
	})
}
