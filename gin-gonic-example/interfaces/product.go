package interfaces

import (
	"apm-trace-with-gin-gonic-example/errs"
	"apm-trace-with-gin-gonic-example/model"
	"context"
)

type ProductUseCase interface {
	GetProducts(ctx context.Context) ([]model.Product, *errs.Err)
}

type ProductRepository interface {
	GetProducts(ctx context.Context) ([]model.Product, error)
}