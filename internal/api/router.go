package api

import (
	"github.com/go-chi/chi"
	"github.com/kopjenmbeng/kanggo_rest_test/internal/application/authentication"
	"github.com/kopjenmbeng/kanggo_rest_test/internal/application/products"
	"github.com/kopjenmbeng/kanggo_rest_test/internal/application/order"
	// "github.com/kopjenmbeng/kanggo_rest_test/internal/application/tes"
)

func routes(r *chi.Mux) {
	r.Route("/v1", func(r chi.Router) {
		r.Mount("/authentication", authentication.Routes())
		r.Mount("/products", products.Routes())
		r.Mount("/order", order.Routes())
	})
}
