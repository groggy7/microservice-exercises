package server

import (
	"microservice/internal/api/http/handlers"
	"microservice/internal/api/http/middlewares"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
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

	opts := middleware.RedocOpts{SpecURL: "../../../../swagger.yaml"}
	docs := middleware.Redoc(opts, nil)

	getRouter := router.Methods("GET").Subrouter()
	getRouter.HandleFunc("/list", r.handler.GetProducts)
	getRouter.Handle("/docs", docs)
	getRouter.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))

	postRouter := router.Methods("POST").Subrouter()
	postRouter.HandleFunc("/add", r.handler.AddProduct)
	postRouter.Use(middlewares.ValidateAddProduct)

	return
}
