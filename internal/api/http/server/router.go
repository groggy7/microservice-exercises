package server

import (
	"microservice/internal/api/http/handlers"
	"microservice/internal/api/http/middleware"

	"github.com/gorilla/mux"
)

type Router struct {
	handler handlers.ProductHandler
}

func NewProductRouter(handler *handlers.ProductHandler) Router {
	return Router{
		handler: *handler,
	}
}

func (r *Router) InitRouter() (router *mux.Router) {
	router = mux.NewRouter()

	listRouter := router.Methods("GET").Subrouter()
	listRouter.HandleFunc("/list", r.handler.GetProducts)

	addRouter := router.Methods("POST").Subrouter()
	addRouter.HandleFunc("/add", r.handler.AddProduct)
	addRouter.Use(middleware.ValidateAddProduct)

	return
}
