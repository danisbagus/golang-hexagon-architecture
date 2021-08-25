package port

import (
	"github.com/danisbagus/golang-hexagon-architecture/internal/core/domain"
	"github.com/danisbagus/golang-hexagon-architecture/internal/dto"
)

type IProductRepo interface {
	FetchAll() ([]domain.Product, error)
	FetchOne(ProductID string) (*domain.Product, error)
}

type IProductService interface {
	GetList() (*dto.ProductListResponse, error)
	GetDetail(ProductID string) (*dto.ProductResponse, error)
}
