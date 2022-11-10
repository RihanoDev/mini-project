package products_test

import (
	"api-mini-project/businesses/categories"
	"api-mini-project/businesses/products"
	_productMock "api-mini-project/businesses/products/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	productRepository _productMock.Repository
	productService    products.Usecase

	productDomain products.Domain
)

func TestMain(m *testing.M) {
	productService = products.NewProductUsecase(&productRepository)

	categoryDomain := categories.Domain{
		Name: "test",
	}

	productDomain = products.Domain{
		Name:        "test",
		Price:       2000,
		Description: "test product",
		CategoryID:  categoryDomain.ID,
		Stock:       200,
	}

	m.Run()
}

func TestGetAll(t *testing.T) {
	t.Run("Get All | Valid", func(t *testing.T) {
		productRepository.On("GetAll").Return([]products.Domain{productDomain}).Once()

		result := productService.GetAll()

		assert.Equal(t, 1, len(result))
	})

	t.Run("Get All | InValid", func(t *testing.T) {
		productRepository.On("GetAll").Return([]products.Domain{}).Once()

		result := productService.GetAll()

		assert.Equal(t, 0, len(result))
	})
}

func TestGeByID(t *testing.T) {
	t.Run("Get By ID | Valid", func(t *testing.T) {
		productRepository.On("GetByID", "1").Return(productDomain).Once()

		result := productService.GetByID("1")

		assert.NotNil(t, result)
	})

	t.Run("Get By ID | InValid", func(t *testing.T) {
		productRepository.On("GetByID", "-1").Return(products.Domain{}).Once()

		result := productService.GetByID("-1")

		assert.NotNil(t, result)
	})
}

func TestCreate(t *testing.T) {
	t.Run("Create | Valid", func(t *testing.T) {
		productRepository.On("Create", &productDomain).Return(productDomain).Once()

		result := productService.Create(&productDomain)

		assert.NotNil(t, result)
	})

	t.Run("Create | InValid", func(t *testing.T) {
		productRepository.On("Create", &productDomain).Return(productDomain).Once()

		result := productService.Create(&productDomain)

		assert.NotNil(t, result)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Update | Valid", func(t *testing.T) {
		productRepository.On("Update", "1", &productDomain).Return(productDomain).Once()

		result := productService.Update("1", &productDomain)

		assert.NotNil(t, result)
	})

	t.Run("Update | InValid", func(t *testing.T) {
		productRepository.On("Update", "1", &products.Domain{}).Return(products.Domain{}).Once()

		result := productService.Update("1", &products.Domain{})

		assert.NotNil(t, result)
	})
}

func TestDelete(t *testing.T) {
	t.Run("Delete | Valid", func(t *testing.T) {
		productRepository.On("Delete", "1").Return(true).Once()

		result := productService.Delete("1")

		assert.True(t, result)
	})

	t.Run("Delete | InValid", func(t *testing.T) {
		productRepository.On("Delete", "-1").Return(false).Once()

		result := productService.Delete("-1")

		assert.False(t, result)
	})
}

func TestRestore(t *testing.T) {
	t.Run("Restore | Valid", func(t *testing.T) {
		productRepository.On("Restore", "1").Return(productDomain).Once()

		result := productService.Restore("1")

		assert.NotNil(t, result)
	})

	t.Run("Restore | InValid", func(t *testing.T) {
		productRepository.On("Restore", "-1").Return(products.Domain{}).Once()

		result := productService.Restore("-1")

		assert.NotNil(t, result)
	})
}

func TestForceDelete(t *testing.T) {
	t.Run("ForceDelete | Valid", func(t *testing.T) {
		productRepository.On("ForceDelete", "1").Return(true).Once()

		result := productService.ForceDelete("1")

		assert.True(t, result)
	})

	t.Run("ForceDelete | InValid", func(t *testing.T) {
		productRepository.On("ForceDelete", "-1").Return(false).Once()

		result := productService.ForceDelete("-1")

		assert.False(t, result)
	})
}
