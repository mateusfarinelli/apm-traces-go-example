package routes

import (
	"apm-trace-with-gorilla-mux-example/controller"
	"net/http"
)

func GetProductRoutes(controller *controller.ProductController) []Route {
	baseUrl := "/products"
	return []Route {
		{
			URI: baseUrl,
			Method: http.MethodGet,
			Function: controller.GetProducts,
		},
	}
}