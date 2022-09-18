package services

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

func (ps *productService) GetProduct() ([]Product, error) {
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
