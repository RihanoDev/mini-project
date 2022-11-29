package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/steinfletcher/apitest"
)

func newApp() *echo.Echo {
	config.InitTestDB()
	app := echo.New()

	routes.SetupRoute(app)

	return app
}

func cleanup(res *http.Response, req *http.Request, apiTest *apitest.APITest) {
	if http.StatusOK == res.StatusCode || http.StatusCreated == res.StatusCode {
		config.CleanSeeders()
	}
}

func getJWTToken(t *testing.T) string {

	user := config.SeedUser()

	var userRequest *model.UserInput = &model.UserInput{
		Email:    user.Email,
		Password: user.Password,
	}

	var resp *http.Response = apitest.New().
		Handler(newApp()).
		Post("/api/v1/users/login").
		JSON(userRequest).
		Expect(t).
		Status(http.StatusOK).
		End().Response

	var response map[string]string = map[string]string{}

	json.NewDecoder(resp.Body).Decode(&response)

	var token string = response["token"]

	var JWT_TOKEN = "Bearer " + token

	return JWT_TOKEN
}

func TestRegister_Success(t *testing.T) {

	var userRequest *model.UserInput = &model.UserInput{
		Email:    "test@mail.com",
		Password: "123123",
	}

	apitest.New().
		Observe(cleanup).
		Handler(newApp()).
		Post("/api/v1/users/register").
		JSON(userRequest).
		Expect(t).
		Status(http.StatusCreated).
		End()
}

func TestRegister_ValidationFailed(t *testing.T) {
	var userRequest *model.UserInput = &model.UserInput{
		Email:    "",
		Password: "",
	}

	apitest.New().
		Handler(newApp()).
		Post("/api/v1/users/register").
		JSON(userRequest).
		Expect(t).
		Status(http.StatusBadRequest).
		End()
}

func TestLogin_Success(t *testing.T) {
	user := config.SeedUser()

	var userRequest *model.UserInput = &model.UserInput{
		Email:    user.Email,
		Password: user.Password,
	}

	apitest.New().
		Handler(newApp()).
		Post("/api/v1/users/login").
		JSON(userRequest).
		Expect(t).
		Status(http.StatusOK).
		End()
}

func TestLogin_ValidationFailed(t *testing.T) {
	var userRequest *model.UserInput = &model.UserInput{
		Email:    "",
		Password: "",
	}

	apitest.New().
		Handler(newApp()).
		Post("/api/v1/users/login").
		JSON(userRequest).
		Expect(t).
		Status(http.StatusBadRequest).
		End()
}

func TestLogin_Failed(t *testing.T) {
	var userRequest *model.UserInput = &model.UserInput{
		Email:    "lol@mail.com",
		Password: "123123",
	}

	apitest.New().
		Handler(newApp()).
		Post("/api/v1/users/login").
		JSON(userRequest).
		Expect(t).
		Status(http.StatusUnauthorized).
		End()
}

func TestCreateProduct_Success(t *testing.T) {
	category := config.SeedCategory()

	var productRequest *model.ProductInput = &model.ProductInput{
		Name:        "test",
		Price:       123,
		Description: "test desc",
		CategoryID:  category.ID,
		Stock:       1000,
	}

	var token string = getJWTToken(t)

	apitest.New().
		Observe(cleanup).
		Handler(newApp()).
		Post("/api/v1/products").
		Header("Authorization", token).
		JSON(productRequest).
		Expect(t).
		Status(http.StatusCreated).
		End()
}

func TestGetProducts_Success(t *testing.T) {
	var token string = getJWTToken(t)

	apitest.New().
		Observe(cleanup).
		Handler(newApp()).
		Get("/api/v1/products").
		Header("Authorization", token).
		Expect(t).
		Status(http.StatusOK).
		End()
}

func TestGetProduct_Success(t *testing.T) {
	var product model.Product = config.SeedProduct()

	productID := strconv.Itoa(int(product.ID))

	var token string = getJWTToken(t)

	apitest.New().
		Observe(cleanup).
		Handler(newApp()).
		Get("/api/v1/products/"+productID).
		Header("Authorization", token).
		Expect(t).
		Status(http.StatusOK).
		End()
}

func TestGetProduct_NotFound(t *testing.T) {
	var product model.Product = config.SeedProduct()

	productID := strconv.Itoa(int(product.ID))

	productID = "0"

	var token string = getJWTToken(t)

	apitest.New().
		Observe(cleanup).
		Handler(newApp()).
		Get("/api/v1/products/"+productID).
		Header("Authorization", token).
		Expect(t).
		Status(http.StatusNotFound).
		End()
}
