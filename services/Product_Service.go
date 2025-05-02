package services

type ProductServiceImpl struct{}

func (h ProductServiceImpl) Teste() string {
	return "Teste"
}

func NewProductService() *ProductServiceImpl {
	return &ProductServiceImpl{}
}
