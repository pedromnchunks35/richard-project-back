package services

import (
	"richard-project-back/dtos"
)

type ProductService interface {
	Teste() string
	GetProduct(Id int64) dtos.Product
	GetImage(Id int64) dtos.Image
	GetDetail(Id int64) dtos.Detail
	InsertProduct(product *dtos.Product) bool
	RemoveProduct(Id int64) bool
	UpdateProduct(Id int64, product *dtos.Product) bool
}
