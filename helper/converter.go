package convert

import (
	"richard-project-back/dtos"
	"richard-project-back/repositories"
)

func ConvertProductRepositoryToDto(product *repositories.Product, detail *repositories.Detail) dtos.Product {

	if product == nil {
		return dtos.Product{}
	} else {
		return dtos.Product{
			Id:   uint(product.Id),
			Name: product.Name,
		}
	}
}
func ConvertProductDtoToRepository(dtoProduct *dtos.Product) repositories.Product {

	if dtoProduct == nil {
		return repositories.NewProduct()
	}
	return repositories.Product{
		Id:         int64(dtoProduct.Id),
		Name:       dtoProduct.Name,
		IdCategory: 0,
		IdImage:    int64(dtoProduct.Image.Id),
	}

}
func ConvertDetailRepositoryToDto(detail *repositories.Detail) dtos.Detail {

	if detail == nil {
		return dtos.Detail{}
	}
	return dtos.Detail{
		Id:          uint(detail.Id),
		Name:        detail.Name,
		Value:       detail.Value,
		Description: detail.Description,
	}
}
func ConvertImageRepositoryToDto(image *repositories.Image) dtos.Image {
	if image == nil {
		return dtos.Image{}
	}
	return dtos.Image{
		Id:     uint(image.Id),
		Name:   image.Name,
		Base64: image.Base64,
	}
}
