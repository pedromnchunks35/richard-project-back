package services

import (
	"richard-project-back/dtos"
	"richard-project-back/helper"
	"richard-project-back/repositories"
)

type ProductServiceImpl struct{}

func NewProductImpl() *ProductServiceImpl {
	return &ProductServiceImpl{}
}

func (h ProductServiceImpl) Teste() string {
	return "GOOD"
}

func (h ProductServiceImpl) GetProduct(Id int64) *dtos.Product {
	return helper.ConvertProductRepositoryToDto(repositories.GetProduct(Id))
}
func (h ProductServiceImpl) GetDetail(Id int64) *dtos.Detail {
	return helper.ConvertDetailRepositoryToDto(repositories.GetDetail(Id))
}
func (h ProductServiceImpl) GetImage(Id int64) *dtos.Image {
	return helper.ConvertImageRepositoryToDto(repositories.GetImage(Id))
}
func (h ProductServiceImpl) InsertProduct(product *repositories.Product) bool {
	if product == nil {
		return false
	}

	return true
}
func (h ProductServiceImpl) UpdateProduct(Id int64, product *repositories.Product) bool {
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
