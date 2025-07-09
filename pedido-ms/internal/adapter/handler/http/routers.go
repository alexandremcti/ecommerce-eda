package http

import (
	"github.com/go-chi/chi/v5"
)

type HandlerDependency struct {
	OrderHandler *OrderHandler
}

func CreateRoutes(r *chi.Mux, hd HandlerDependency) {
	r.Route("/pedidos", func(r chi.Router) {
		r.Post("/", hd.OrderHandler.CreateOrder)
	})
}
