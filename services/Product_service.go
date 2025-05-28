package services

import (
	"richard-project-back/convert"
	"richard-project-back/dtos"
	"richard-project-back/repositories"
)

type ProductServiceImpl struct{}

func NewProductImpl() *ProductServiceImpl {
	return &ProductServiceImpl{}
}

func (h ProductServiceImpl) Teste() string {
	return "GOOD"
}

func (h ProductServiceImpl) GetProduct(Id int64) dtos.Product {
	return convert.onvertProductRepositoryToDto(repositories.NewProduct())
}
func (h ProductServiceImpl) GetDetail(Id int64) dtos.Detail {
	return convert.ConvertDetailRepositoryToDto(repositories.NewDetail())
}
func (h ProductServiceImpl) GetImage(Id int64) dtos.Image {
	return convert.ConvertImageRepositoryToDto(repositories.NewImage())
}
func (h ProductServiceImpl) InsertProduct(product *dtos.Product) bool {
	if product == nil {
		return false
	}

	return true
}
func (h ProductServiceImpl) UpdateProduct(Id int64, product *dtos.Product) bool {
	if product == nil {
		return false
	}
	if Id == 0 {
		return false
	}

	return true
}
func (h ProductServiceImpl) RemoveProduct(Id int64) bool {
	if Id == 0 {
		return false
	}

	return true
}
