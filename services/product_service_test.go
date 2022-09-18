package services

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetProductWithTableTest(t *testing.T) {
	result := []Product{
		{
			ID:   "1",
			Name: "Product 1",
		},
		{
			ID:   "2",
			Name: "Product 2",
		},
	}
	p := NewProductService(mock.Anything)
	tests := []struct {
		name           string
		page           string
		mockFunc       func()
		expectedResult []Product
		expectingErr   bool
	}{
		{
			name:           "when success get products",
			page:           "1",
			mockFunc:       func() {},
			expectedResult: result,
			expectingErr:   false,
		},
		{
			name:           "when failed get products",
			page:           "2",
			mockFunc:       func() {},
			expectedResult: nil,
			expectingErr:   false,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(tt *testing.T) {
			tc.mockFunc()
			result, err := p.GetProduct(tc.page)
			assert.Equal(tt, tc.expectedResult, result)
			if tc.expectingErr {
				assert.Error(tt, err)
			}
		})
	}
}

func TestValidateID(t *testing.T) {
	p := NewProductService("")
	tests := []struct {
		name           string
		productID      string
		expectedResult error
		expectingErr   bool
	}{
		{
			name:           "when success pass validated",
			productID:      "1",
			expectedResult: nil,
			expectingErr:   false,
		},
		{
			name:           "when error validated",
			productID:      "0",
			expectedResult: fmt.Errorf("something happened"),
			expectingErr:   true,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(tt *testing.T) {
			err := p.ValidateID(tc.productID)
			assert.Equal(tt, tc.expectedResult, err)
			if tc.expectingErr {
				assert.Error(tt, err)
			}
		})
	}
}
