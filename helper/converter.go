package helper

import (
	"encoding/json"
	"richard-project-back/dtos"
	repositories "richard-project-back/repositories/Models"

	"github.com/jackc/pgx/v4"
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

func RowsToJSON(rows pgx.Rows) ([]byte, error) {
	defer rows.Close()

	results := []map[string]interface{}{}

	cols := rows.FieldDescriptions() // Info das colunas
	colNames := make([]string, len(cols))
	for i, col := range cols {
		colNames[i] = string(col.Name)
	}

	for rows.Next() {
		values := make([]interface{}, len(cols))
		valuePtrs := make([]interface{}, len(cols))
		for i := range values {
			valuePtrs[i] = &values[i]
		}

		err := rows.Scan(valuePtrs...)
		if err != nil {
			return nil, err
		}

		rowMap := make(map[string]interface{})
		for i, colName := range colNames {
			val := values[i]

			if b, ok := val.([]byte); ok {
				rowMap[colName] = string(b)
			} else {
				rowMap[colName] = val
			}
		}

		results = append(results, rowMap)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return json.Marshal(results)
}
func JsonToProductList(jsonData []byte) ([]repositories.Product, error) {
	var products []repositories.Product
	err := json.Unmarshal(jsonData, &products)
	if err != nil {
		return nil, err
	}
	return products, nil
}
