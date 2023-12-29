package middleware

import (
	"context"
	"encoding/json"
	"log"
	"microservice/pkg/dto"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func ValidateAddProduct(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		product := dto.AddProductRequest{}
		if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}

		validator := validator.New()

		if err := validator.Struct(product); err != nil {
			log.Printf("Validation error: %s", err.Error())
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), dto.AddProductRequest{}, product)
		r = r.WithContext(ctx)
		next.ServeHTTP(rw, r)
	})
}
