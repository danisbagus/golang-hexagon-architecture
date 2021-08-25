package service

import (
	"github.com/danisbagus/golang-hexagon-architecture/internal/core/port"
	"github.com/danisbagus/golang-hexagon-architecture/internal/dto"
)

type ProductService struct {
	repo port.IProductRepo
}

func NewProductService(repo port.IProductRepo) port.IProductService {
	return &ProductService{
		repo: repo,
	}
}

func (r ProductService) GetList() (*dto.ProductListResponse, error) {
	productList, _ := r.repo.FetchAll()
	return dto.NewGetListProductResponse(productList), nil
}

func (r ProductService) GetDetail(productID string) (*dto.ProductResponse, error) {
	product, _ := r.repo.FetchOne(productID)
	return dto.GetDetailProductResponse(product), nil
}
