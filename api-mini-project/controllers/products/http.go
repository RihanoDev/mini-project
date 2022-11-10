package products

import (
	"api-mini-project/businesses/products"
	controller "api-mini-project/controllers"
	"api-mini-project/controllers/products/request"
	"api-mini-project/controllers/products/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ProductController struct {
	productUseCase products.Usecase
}

func NewProductController(productUC products.Usecase) *ProductController {
	return &ProductController{
		productUseCase: productUC,
	}
}

func (ctrl *ProductController) GetAllProducts(c echo.Context) error {
	productsData := ctrl.productUseCase.GetAll()

	products := []response.Product{}

	for _, product := range productsData {
		products = append(products, response.FromDomain(product))
	}

	return controller.NewResponse(c, http.StatusOK, "Success", "All Products", products)
}

func (ctrl *ProductController) GetProductByID(c echo.Context) error {
	var id string = c.Param("id")

	product := ctrl.productUseCase.GetByID(id)

	if product.ID == 0 {
		return controller.NewResponse(c, http.StatusNotFound, "Failed", "Product Not Found", "")
	}

	return controller.NewResponse(c, http.StatusOK, "Success", "Category Found", response.FromDomain(product))
}

func (ctrl *ProductController) CreateProduct(c echo.Context) error {
	input := request.Product{}

	if err := c.Bind(&input); err != nil {
		return controller.NewResponse(c, http.StatusBadRequest, "Failed", "Validation Failed", "")
	}

	err := input.Validate()

	if err != nil {
		return controller.NewResponse(c, http.StatusBadRequest, "Failed", "Validation Failed", "")
	}

	product := ctrl.productUseCase.Create(input.ToDomain())

	return controller.NewResponse(c, http.StatusCreated, "Success", "Product Created", response.FromDomain(product))
}

func (ctrl *ProductController) UpdateProduct(c echo.Context) error {
	input := request.Product{}

	if err := c.Bind(&input); err != nil {
		return controller.NewResponse(c, http.StatusBadRequest, "Failed", "Validation Failed", "")
	}

	var id string = c.Param("id")

	err := input.Validate()

	if err != nil {
		return controller.NewResponse(c, http.StatusBadRequest, "Failed", "Validation Failed", "")
	}

	product := ctrl.productUseCase.Update(id, input.ToDomain())

	return controller.NewResponse(c, http.StatusOK, "Success", "Product Updated", response.FromDomain(product))
}

func (ctrl *ProductController) DeleteProduct(c echo.Context) error {
	var productId string = c.Param("id")

	isSuccess := ctrl.productUseCase.Delete(productId)

	if !isSuccess {
		return controller.NewResponse(c, http.StatusNotFound, "Failed", "Product Not Found", "")
	}

	return controller.NewResponse(c, http.StatusOK, "Success", "Product Deleted", "")
}

func (ctrl *ProductController) RestoreProduct(c echo.Context) error {
	var productId string = c.Param("id")

	product := ctrl.productUseCase.Restore(productId)

	if product.ID == 0 {
		return controller.NewResponse(c, http.StatusNotFound, "Failed", "Product Not Found", "")
	}

	return controller.NewResponse(c, http.StatusOK, "Success", "Product Restored", response.FromDomain(product))
}

func (ctrl *ProductController) ForceDeleteProduct(c echo.Context) error {
	var productId string = c.Param("id")

	isSuccess := ctrl.productUseCase.ForceDelete(productId)

	if !isSuccess {
		return controller.NewResponse(c, http.StatusNotFound, "Failed", "Product Not Found", "")
	}

	return controller.NewResponse(c, http.StatusOK, "Success", "Product Deleted Permanently", "")
}
