package services
import "richard-project-back/repositories"
type ProductServiceImpl struct{}

func NewProductImpl() *ProductServiceImpl {
	return &ProductServiceImpl{}
}

func(h ProductServiceImpl) Teste() string{
	return "GOOD"
}

func (h ProductServiceImpl) GetProduct(Id int64) repositories.Product {
	return repositories.Product.NewProduct()
}

func(h ProductServiceImpl) InsertProduct(product repositories.Product) bool{
	if product == nil{
		return false
	}


	return true
}
func(h ProductServiceImpl) UpdateProduct(Id int64, product repositories.Product ) bool{
	if product == nil{
		return false
	}
	if Id == nil {
		return false
	}
	if Id == 0{
		return false
	}
	
	
	return true
}
func(h ProductServiceImpl) RemoveProduct(Id int64) bool{
	if Id == nil {
		return false
	}
	if Id == 0{
		return false
	}

	return true
}