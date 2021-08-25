package repo

import (
	"github.com/danisbagus/golang-hexagon-architecture/internal/core/domain"
	"github.com/danisbagus/golang-hexagon-architecture/internal/core/port"
)

type ProductRepo struct{}

var ProductList = []domain.Product{
	{ProductID: "1001", ProductName: "Lenovo A1000", ProductCategory: "ELECTRONIC", Quantity: 1000},
	{ProductID: "1002", ProductName: "POCO X1", ProductCategory: "ELECTRONIC", Quantity: 1000},
	{ProductID: "1003", ProductName: "Samsung A52", ProductCategory: "ELECTRONIC", Quantity: 1000},
}

func NewProductRepo() port.IProductRepo {
	return &ProductRepo{}
}

func (r ProductRepo) FetchAll() ([]domain.Product, error) {
	return ProductList, nil
}

func (r ProductRepo) FetchOne(ProductID string) (*domain.Product, error) {
	product := domain.Product{}
	for _, v := range ProductList {
		if v.ProductID == ProductID {
			product = v
		}
	}
	return &product, nil
}
