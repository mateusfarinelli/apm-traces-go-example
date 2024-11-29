package controller

import (
	"apm-trace-with-gorilla-mux-example/handles"
	"apm-trace-with-gorilla-mux-example/interfaces"
	"context"
	"net/http"
)

type ProductController struct {
	producUseCase interfaces.ProductUseCase
}

func NewProductController(usecase interfaces.ProductUseCase) *ProductController {
	return &ProductController{
		producUseCase: usecase,
	}
}

func (p *ProductController) GetProducts(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithCancel(r.Context())
	defer cancel()

	products, err := p.producUseCase.GetProducts(ctx)

	if err != nil {
		handles.Error(w, err)
		return
	}

	if products == nil || len(products) == 0 {
		handles.Success(w, http.StatusNoContent, nil)
		return 
	}

	handles.Success(w, http.StatusOK, products)

}