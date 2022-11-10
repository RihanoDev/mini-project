package categories

import (
	"api-mini-project/businesses/categories"
	controller "api-mini-project/controllers"
	"api-mini-project/controllers/categories/request"
	"api-mini-project/controllers/categories/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CategoryController struct {
	categoryUseCase categories.Usecase
}

func NewCategoryController(categoryUC categories.Usecase) *CategoryController {
	return &CategoryController{
		categoryUseCase: categoryUC,
	}
}

func (ctrl *CategoryController) GetAllCategories(c echo.Context) error {
	categoriesData := ctrl.categoryUseCase.GetAll()

	categories := []response.Category{}

	for _, category := range categoriesData {
		categories = append(categories, response.FromDomain(category))
	}

	return controller.NewResponse(c, http.StatusOK, "Success", "All Categories", categories)
}

func (ctrl *CategoryController) GetCategoriesByID(c echo.Context) error {
	var id string = c.Param("id")

	category := ctrl.categoryUseCase.GetByID(id)

	if category.ID == 0 {
		return controller.NewResponse(c, http.StatusNotFound, "Failed", "Category Not Found", "")
	}

	return controller.NewResponse(c, http.StatusOK, "Success", "Category Found", response.FromDomain(category))
}

func (ctrl *CategoryController) CreateCategories(c echo.Context) error {
	input := request.Category{}

	if err := c.Bind(&input); err != nil {
		return controller.NewResponse(c, http.StatusBadRequest, "Failed", "Validation Failed", "")
	}

	err := input.Validate()

	if err != nil {
		return controller.NewResponse(c, http.StatusBadRequest, "Failed", "Validation Failed", "")
	}

	sameCategories := ctrl.categoryUseCase.CheckData(input.ToDomain()).Name

	if input.Name == sameCategories {
		return controller.NewResponse(c, http.StatusBadRequest, "Failed", "Name Already Used!", "")
	}

	category := ctrl.categoryUseCase.Create(input.ToDomain())

	return controller.NewResponse(c, http.StatusCreated, "Success", "Category Created", response.FromDomain(category))
}

func (ctrl *CategoryController) UpdateCategories(c echo.Context) error {
	input := request.Category{}

	if err := c.Bind(&input); err != nil {
		return controller.NewResponse(c, http.StatusBadRequest, "Failed", "Validation Failed", "")
	}

	var id string = c.Param("id")

	err := input.Validate()

	if err != nil {
		return controller.NewResponse(c, http.StatusBadRequest, "Failed", "Validation Failed", "")
	}

	sameCategories := ctrl.categoryUseCase.CheckData(input.ToDomain()).Name

	if input.Name == sameCategories {
		return controller.NewResponse(c, http.StatusBadRequest, "Failed", "Name Already Used", "")
	}

	category := ctrl.categoryUseCase.Update(id, input.ToDomain())

	return controller.NewResponse(c, http.StatusOK, "Success", "Category Updated", response.FromDomain(category))
}

func (ctrl *CategoryController) DeleteCategories(c echo.Context) error {
	var categoryId string = c.Param("id")

	isSuccess := ctrl.categoryUseCase.Delete(categoryId)

	if !isSuccess {
		return controller.NewResponse(c, http.StatusNotFound, "Failed", "Category Not Found", "")
	}

	return controller.NewResponse(c, http.StatusOK, "Success", "Category Deleted", "")
}
