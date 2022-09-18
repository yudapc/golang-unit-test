package services

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCategoriesWithTestCases(t *testing.T) {
	tests := []struct {
		name           string
		page           int
		mockFunc       func()
		expectedResult []Category
		expectingErr   bool
	}{
		{
			name: "should return categories without error",
			page: 1,
			mockFunc: func() {
				categoryRepo = func() ([]Category, error) {
					return []Category{
						{
							ID:   "1",
							Name: "Category 1",
						},
					}, nil
				}
			},
			expectedResult: []Category{
				{
					ID:   "1",
					Name: "Category 1",
				},
			},
			expectingErr: false,
		},
		{
			name: "should return nil with error",
			page: 1,
			mockFunc: func() {
				categoryRepo = func() ([]Category, error) {
					return nil, fmt.Errorf("something happened")
				}
			},
			expectedResult: nil,
			expectingErr:   true,
		},
		{
			name:           "should return nil with error",
			page:           0,
			mockFunc:       func() {},
			expectedResult: nil,
			expectingErr:   true,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(tt *testing.T) {
			tc.mockFunc()
			result, err := GetCategories(tc.page)
			assert.Equal(tt, tc.expectedResult, result)
			if tc.expectingErr {
				assert.Error(tt, err)
			}
		})
	}
}

func TestGetCategoryByID(t *testing.T) {
	tests := []struct {
		name           string
		CategoryID     string
		mockFunc       func()
		expectedResult *Category
		expectingErr   bool
	}{
		{
			name:       "test GetCategoryByID return with data",
			CategoryID: "1",
			mockFunc: func() {
				validateCategoryID = func(id string) error {
					return nil
				}
			},
			expectedResult: &Category{
				ID:   "1",
				Name: "Category 1",
			},
			expectingErr: false,
		},
		{
			name:       "test GetCategoryByID return error",
			CategoryID: "1",
			mockFunc: func() {
				validateCategoryID = func(id string) error {
					return fmt.Errorf("something happened")
				}
			},
			expectedResult: nil,
			expectingErr:   true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(tt *testing.T) {
			tc.mockFunc()
			result, err := GetCategoryByID(tc.CategoryID)
			assert.Equal(tt, tc.expectedResult, result)
			if tc.expectingErr {
				assert.Error(tt, err)
			}
		})
	}
}

func Test_validateID(t *testing.T) {
	tests := []struct {
		name           string
		CategoryID     string
		expectedResult error
		expectingErr   bool
	}{
		{
			name:           "test validateID return true with error",
			CategoryID:     "2",
			expectedResult: fmt.Errorf("something happened"),
			expectingErr:   true,
		},
		{
			name:           "test validateID return false with no error",
			CategoryID:     "1",
			expectedResult: nil,
			expectingErr:   false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(tt *testing.T) {
			err := validateID(tc.CategoryID)
			if err != nil {
				assert.Error(tt, err)
			} else {
				assert.Equal(tt, tc.expectedResult, err)
			}
		})
	}
}

func Test_repoCategory(t *testing.T) {
	tests := []struct {
		name           string
		expectedResult error
		expectingErr   bool
	}{
		{
			name:           "test repoCategory return false with no error",
			expectedResult: nil,
			expectingErr:   false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(tt *testing.T) {
			result, err := repoCategories()
			if err != nil {
				assert.Error(tt, err)
			} else {
				assert.Equal(tt, tc.expectedResult, result)
			}
		})
	}
}
