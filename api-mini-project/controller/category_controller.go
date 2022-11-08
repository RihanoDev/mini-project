package controller

import (
	"api-mini-project/model"
	"api-mini-project/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

var CategoryService service.CategoryService = service.NewCategoryService()

func GetAllCategories(c echo.Context) error {
	var categories []model.Category = CategoryService.GetAllCategories()

	return c.JSON(http.StatusOK, categories)
}

func GetCategoriesByID(c echo.Context) error {
	var id string = c.Param("id")

	category := CategoryService.GetCategoriesByID(id)

	if category.ID == 0 {
		return c.JSON(http.StatusNotFound, map[string]string{
			"messsage": "Category Not Found",
		})
	}

	return c.JSON(http.StatusOK, category)
}

func CreateCategories(c echo.Context) error {
	var input *model.CategoryInput = new(model.CategoryInput)

	if err := c.Bind(input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Invalid Request",
		})
	}

	err := input.Validate()

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Validation Failed",
		})
	}

	sameCategories := CategoryService.CheckDataCategories(*input).Name

	if input.Name == sameCategories {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Category Already Available",
		})
	}

	category := CategoryService.CreateCategories(*input)

	return c.JSON(http.StatusCreated, category)
}

func UpdateCategories(c echo.Context) error {
	var input *model.CategoryInput = new(model.CategoryInput)

	if err := c.Bind(input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Invalid Request",
		})
	}

	err := input.Validate()

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Validation Failed",
		})
	}

	sameCategories := CategoryService.CheckDataCategories(*input).Name

	if input.Name == sameCategories {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Category Already Available",
		})
	}

	var categoryId string = c.Param("id")

	category := CategoryService.UpdateCategories(categoryId, *input)

	return c.JSON(http.StatusOK, category)
}

func ForceDeleteCategories(c echo.Context) error {
	var categoryId string = c.Param("id")

	isSuccess := CategoryService.ForceDeleteCategories(categoryId)

	if !isSuccess {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Failed to Delete Category",
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Category Deleted",
	})
}
