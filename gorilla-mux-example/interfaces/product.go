package interfaces

import (
	"apm-trace-with-gorilla-mux-example/errs"
	"apm-trace-with-gorilla-mux-example/model"
	"context"
)

type ProductUseCase interface {
	GetProducts(ctx context.Context) ([]model.Product, *errs.Err)
}

type ProductRepository interface {
	GetProducts(ctx context.Context) ([]model.Product, error)
}