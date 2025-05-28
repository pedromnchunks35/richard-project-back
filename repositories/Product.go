package repositories

type Product struct {
	Id         int64  `json:"id"`
	Name       string `json:"name"`
	IdCategory int64  `json:"id_category"`
	IdImage    int64  `json:"id_image"`
}

func NewProduct() *Product {
	return &Product{}
}
func GetProduct(id int64) *Product {
	return NewProduct()
}

func GetProducts() *[]Product {
	return &[]Product{}
}
