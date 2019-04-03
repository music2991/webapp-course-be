package routes

import "webapp-course-be/services/products"

func getProductRoutes() []routeModel {
	return []routeModel{
		{
			Path:    "/api/Product",
			Handler: products.Post,
			Method:  "POST",
		},
		{
			Path:    "/api/Product",
			Handler: products.Get,
			Method:  "GET",
		},
		{
			Path:    "/api/Product/{id}",
			Handler: products.GetById,
			Method:  "GET",
		},
		{
			Path:    "/api/Product",
			Handler: products.Put,
			Method:  "PUT",
		},
		{
			Path:    "/api/Product/{id}",
			Handler: products.Delete,
			Method:  "DELETE",
		},
	}
}
