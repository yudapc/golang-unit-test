package services

import "fmt"

type Product struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type productService struct {
	db string
}

func NewProductService(p string) *productService {
	return &productService{
		db: p,
	}
}

func (ps *productService) GetProduct(page string) ([]Product, error) {
	err := ps.ValidateID(page)
	if err != nil {
		return nil, err
	}
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
	return result, nil
}

func (ps *productService) ValidateID(id string) error {
	if id == "1" {
		return nil
	}
	return fmt.Errorf("something happened")
}
