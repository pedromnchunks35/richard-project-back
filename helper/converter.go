package helper

import (
	"richard-project-back/dtos"
	"richard-project-back/repositories"
)

func ConvertListDtoProd(products []dtos.Product) *[]repositories.Product {
	var newproducts []repositories.Product
	for _, product := range products {
		newproducts = append(newproducts, *ConvertProductDtoToRepository(&product))

	}

	return &newproducts
}

func ConvertProductRepositoryToDto(product *repositories.Product) *dtos.Product {

	return &dtos.Product{
		Id:   uint(product.Id),
		Name: product.Name,
	}

}

func ConvertProductDtoToRepository(dtoProduct *dtos.Product) *repositories.Product {

	return &repositories.Product{
		Id:         int64(dtoProduct.Id),
		Name:       dtoProduct.Name,
		IdCategory: 0,
		IdImage:    int64(dtoProduct.Image.Id),
	}

}
func ConvertDetailRepositoryToDto(detail *repositories.Detail) *dtos.Detail {

	return &dtos.Detail{
		Id:          uint(detail.Id),
		Name:        detail.Name,
		Value:       detail.Value,
		Description: detail.Description,
	}
}
func ConvertImageRepositoryToDto(image *repositories.Image) *dtos.Image {

	return &dtos.Image{
		Id:     uint(image.Id),
		Name:   image.Name,
		Base64: image.Base64,
	}
}
