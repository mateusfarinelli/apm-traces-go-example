package interfaces

import (
	"apm-trace-worker-example/errs"
	"apm-trace-worker-example/model"
	"context"
)

type ProductUseCase interface {
	GetProducts(ctx context.Context) ([]model.Product, *errs.Err)
}

type ProductRepository interface {
	GetProducts(ctx context.Context) ([]model.Product, error)
}
