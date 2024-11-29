package dependencies

import (
	"apm-trace-worker-example/db"
	"apm-trace-worker-example/interfaces"
	"apm-trace-worker-example/repositories"
	"apm-trace-worker-example/usecases"
)

func GetProductUseCase() interfaces.ProductUseCase {
	return usecases.NewProductUseCase(GetProductRepository())
}

func GetProductRepository() interfaces.ProductRepository {
	return repositories.NewProductRepository(db.Conn)
}
