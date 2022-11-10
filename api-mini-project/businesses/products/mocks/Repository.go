// Code generated by mockery v2.14.1. DO NOT EDIT.

package mocks

import (
	products "api-mini-project/businesses/products"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// Create provides a mock function with given fields: productDomain
func (_m *Repository) Create(productDomain *products.Domain) products.Domain {
	ret := _m.Called(productDomain)

	var r0 products.Domain
	if rf, ok := ret.Get(0).(func(*products.Domain) products.Domain); ok {
		r0 = rf(productDomain)
	} else {
		r0 = ret.Get(0).(products.Domain)
	}

	return r0
}

// Delete provides a mock function with given fields: id
func (_m *Repository) Delete(id string) bool {
	ret := _m.Called(id)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// ForceDelete provides a mock function with given fields: id
func (_m *Repository) ForceDelete(id string) bool {
	ret := _m.Called(id)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// GetAll provides a mock function with given fields:
func (_m *Repository) GetAll() []products.Domain {
	ret := _m.Called()

	var r0 []products.Domain
	if rf, ok := ret.Get(0).(func() []products.Domain); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]products.Domain)
		}
	}

	return r0
}

// GetByID provides a mock function with given fields: id
func (_m *Repository) GetByID(id string) products.Domain {
	ret := _m.Called(id)

	var r0 products.Domain
	if rf, ok := ret.Get(0).(func(string) products.Domain); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(products.Domain)
	}

	return r0
}

// Restore provides a mock function with given fields: id
func (_m *Repository) Restore(id string) products.Domain {
	ret := _m.Called(id)

	var r0 products.Domain
	if rf, ok := ret.Get(0).(func(string) products.Domain); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(products.Domain)
	}

	return r0
}

// Update provides a mock function with given fields: id, productDomain
func (_m *Repository) Update(id string, productDomain *products.Domain) products.Domain {
	ret := _m.Called(id, productDomain)

	var r0 products.Domain
	if rf, ok := ret.Get(0).(func(string, *products.Domain) products.Domain); ok {
		r0 = rf(id, productDomain)
	} else {
		r0 = ret.Get(0).(products.Domain)
	}

	return r0
}

type mockConstructorTestingTNewRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewRepository creates a new instance of Repository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRepository(t mockConstructorTestingTNewRepository) *Repository {
	mock := &Repository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}