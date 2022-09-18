package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockProduct struct {
	mock.Mock
}

func (m *mockProduct) GetProduct() ([]Product, error) {
	result := []Product{
		{
			ID:   "1",
			Name: "Product 1",
		},
	}

	return result, nil
}

func TestGetProduct(t *testing.T) {
	p := NewProductService("test")
	result, _ := p.GetProduct()
	actual := len(result)
	expected := 2
	assert.Equal(t, expected, actual)
}

func TestFailedGetProduct(t *testing.T) {
	p := new(mockProduct)
	mockResult := []Product{
		{
			ID:   "1",
			Name: "Product 1",
		},
	}
	p.On("GetProduct").Return(mockResult, nil)

	result, _ := p.GetProduct()
	actual := len(result)
	expected := 1
	assert.Equal(t, expected, actual)
}
