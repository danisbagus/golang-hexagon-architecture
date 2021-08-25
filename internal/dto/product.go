package dto

import (
	"github.com/danisbagus/golang-hexagon-architecture/internal/core/domain"
)

type ProductResponse struct {
	ProductID       string `json:"product_id"`
	ProductName     string `json:"product_name"`
	ProductCategory string `json:"product_category"`
	Quantity        int64  `json:"quantity"`
}

type ProductListResponse struct {
	Data []ProductResponse `json:"data"`
}

func NewGetListProductResponse(data []domain.Product) *ProductListResponse {
	productList := make([]ProductResponse, 0)
	for _, v := range data {
		product := ProductResponse{
			ProductID:       v.ProductID,
			ProductName:     v.ProductName,
			ProductCategory: v.ProductCategory,
			Quantity:        v.Quantity,
		}
		productList = append(productList, product)
	}
	return &ProductListResponse{productList}
}

func GetDetailProductResponse(data *domain.Product) *ProductResponse {
	product := ProductResponse{
		ProductID:       data.ProductID,
		ProductName:     data.ProductName,
		ProductCategory: data.ProductCategory,
		Quantity:        data.Quantity,
	}
	return &product
}
