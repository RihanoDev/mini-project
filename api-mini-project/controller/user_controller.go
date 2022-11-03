package controller

import (
	"api-mini-project/config"
	"api-mini-project/model"
	"net/http"

	"github.com/labstack/echo"
)

func RegisterUser(r echo.Context) error {
	user := model.User{}
	r.Bind(&user)

	err := config.DB.Save(&user).Error
	if err != nil {
		return r.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return r.JSON(http.StatusCreated, map[string]interface{}{
		"message": "Success Register User",
		"data":    user,
	})
}

// func UpdateUser(u echo.Context) error {

// }

// func DeleteUser(d echo.Context) error {

// }

// func GetUser(g echo.Context) error {

// }

// func GetUsers(g echo.Context) error {

// }

// func Login(l echo.Context) error {

// }
