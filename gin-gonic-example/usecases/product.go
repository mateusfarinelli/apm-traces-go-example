package usecases

import (
	"apm-trace-with-gin-gonic-example/errs"
	"apm-trace-with-gin-gonic-example/interfaces"
	"apm-trace-with-gin-gonic-example/model"
	"context"
)

type ProductUseCase struct {
	productRepository interfaces.ProductRepository
}

func NewProductUseCase(productRepository interfaces.ProductRepository) interfaces.ProductUseCase {
	return &ProductUseCase{productRepository}
}

func (pu *ProductUseCase) GetProducts(ctx context.Context) ([]model.Product, *errs.Err) {
	products, err := pu.productRepository.GetProducts(ctx)

	if err != nil {
		return nil, errs.InternalServerError(err.Error())
	}

	return products, nil
}