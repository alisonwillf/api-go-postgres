package usecase

import (
	"github.com/alisonferreira/APIGO/model"
	"github.com/alisonferreira/APIGO/repository"
)

type ProductUsecase struct {
	repository *repository.ProductRepository
	//repository
}

func NewProductUsecase(repo *repository.ProductRepository) *ProductUsecase {

	return &ProductUsecase{
		repository: repo,
	}
}

func (pu *ProductUsecase) GetProducts() ([]model.Product, error) {
	return pu.repository.GetProducts()
}

func (pu *ProductUsecase) CreateProduct(product model.Product) (model.Product, error) {

	productId, err := pu.repository.CreateProduct(product)
	if err != nil {
		return model.Product{}, err
	}

	product.ID = productId
	return product, nil
}

func (pu *ProductUsecase) GetProductsById(id_product int) (*model.Product, error) {

	product, err := pu.repository.GetProductsById(id_product)
	if err != nil {
		return nil, err
	}
	return product, nil

}
