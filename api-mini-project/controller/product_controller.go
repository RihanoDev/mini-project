package controller

import (
	"api-mini-project/model"
	"api-mini-project/service"
	"net/http"

	"github.com/labstack/echo"
)

var ProductService service.ProductService = service.New()

func GetAll(c echo.Context) error {
	var products []model.Product = ProductService.GetAll()

	return c.JSON(http.StatusOK, products)
}

func GetByID(c echo.Context) error {
	var id string = c.Param("id")

	product := ProductService.GetByID(id)

	if product.ID == 0 {
		return c.JSON(http.StatusNotFound, map[string]string{
			"message": "Product not Found",
		})
	}

	return c.JSON(http.StatusOK, product)
}

func Create(c echo.Context) error {
	var input *model.ProductInput = new(model.ProductInput)

	if err := c.Bind(input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Invalid Request",
		})
	}

	product := ProductService.Create(*input)

	return c.JSON(http.StatusCreated, product)
}

func Update(c echo.Context) error {
	var input *model.ProductInput = new(model.ProductInput)

	if err := c.Bind(input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Invalid Request",
		})
	}

	var productId string = c.Param("id")

	product := ProductService.Update(productId, *input)

	return c.JSON(http.StatusOK, product)
}

func Delete(c echo.Context) error {
	var productId string = c.Param("id")

	isSuccess := ProductService.Delete(productId)

	if !isSuccess {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "failed to delete a data",
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Data deleted",
	})
}

func Restore(c echo.Context) error {
	var productId string = c.Param("id")

	product := ProductService.Restore(productId)

	return c.JSON(http.StatusOK, map[string]any{
		"message": "Data Restored",
		"data":    product,
	})
}

func ForceDelete(c echo.Context) error {
	var productId string = c.Param("id")

	isSuccess := ProductService.ForceDelete(productId)

	if !isSuccess {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "failed to delete a data",
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Data force deleted",
	})
}
