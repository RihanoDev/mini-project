package controller

import (
	"net/http"
	"rest/config"
	"rest/model"

	"github.com/labstack/echo"
)

func AddProductController(a echo.Context) error {
	product := model.Product{}
	a.Bind(&product)

	err := config.DB.Save(&product).Error
	if err != nil {
		return a.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return a.JSON(http.StatusInternalServerError, map[string]interface{}{
		"message": "Success add product",
		"data":    product,
	})
}
