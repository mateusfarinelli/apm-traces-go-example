package controller

import (
	"apm-trace-with-gin-gonic-example/interfaces"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	producUseCase interfaces.ProductUseCase
}

func NewProductController(usecase interfaces.ProductUseCase) *ProductController {
	return &ProductController{
		producUseCase: usecase,
	}
}

func (p *ProductController) GetProducts(ctx *gin.Context) {
	products, err := p.producUseCase.GetProducts(ctx)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
		return 
	}

	if products == nil || len(products) == 0 {
		ctx.JSON(http.StatusNoContent, nil)
		return 
	}

	ctx.JSON(http.StatusOK, products)

}