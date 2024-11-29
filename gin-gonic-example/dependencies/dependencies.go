package dependencies

import (
	"apm-trace-with-gin-gonic-example/controller"
	"apm-trace-with-gin-gonic-example/db"
	"apm-trace-with-gin-gonic-example/interfaces"
	"apm-trace-with-gin-gonic-example/repositories"
	"apm-trace-with-gin-gonic-example/usecases"
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