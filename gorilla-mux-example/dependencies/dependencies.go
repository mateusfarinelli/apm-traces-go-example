package dependencies

import (
	"apm-trace-with-gorilla-mux-example/controller"
	"apm-trace-with-gorilla-mux-example/db"
	"apm-trace-with-gorilla-mux-example/interfaces"
	"apm-trace-with-gorilla-mux-example/repositories"
	"apm-trace-with-gorilla-mux-example/usecases"
)

func GetProductController() *controller.ProductController {
	return controller.NewProductController(GetProductUseCase())
}

func GetProductUseCase() interfaces.ProductUseCase {
	return usecases.NewProductUseCase(GetProductRepository())
}

func GetProductRepository() interfaces.ProductRepository {
	return repositories.NewProductRepository(db.Conn)
}