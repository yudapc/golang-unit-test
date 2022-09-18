package services

import "fmt"

type Category struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// private function
func repoCategories() ([]Category, error) {
	return []Category{
		{
			ID:   "1",
			Name: "Thomas",
		},
		{
			ID:   "2",
			Name: "Mark",
		},
	}, nil
}

var categoryRepo = repoCategories

func GetCategories(page int) ([]Category, error) {
	if page > 0 {
		result, err := categoryRepo()
		if err != nil {
			return nil, err
		}
		return result, err
	}
	return nil, fmt.Errorf("something happened")
}

var validateCategoryID = validateID

func GetCategoryByID(id string) (*Category, error) {
	err := validateCategoryID(id)
	if err != nil {
		return nil, err
	}
	return &Category{
		ID:   "1",
		Name: "Category 1",
	}, nil
}

// private function
func validateID(id string) error {
	if id == "1" {
		return nil
	}
	return fmt.Errorf("something happened")
}
