package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"pedido-ms/internal/adapter/broker"
	"pedido-ms/internal/adapter/config"
	"pedido-ms/internal/adapter/database"
	handlers "pedido-ms/internal/adapter/handler/http"
	"pedido-ms/internal/core/services"

	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

func main() {
	config, err := config.New()
	if err != nil {
		slog.Error("Error loading environment variables", "error", err)
		os.Exit(1)
	}

	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	if err := database.CreateConnection(ctx, config.MongoConnectionUrl); err != nil {
		slog.Error("Error creating database connection", "error", err)
		os.Exit(1)
	}

	if err := broker.CreateConnection(config.RabbitHost); err != nil {
		slog.Error("Error creating broker connection", "error", err)
		os.Exit(1)
	}

	err = broker.B.CreatePublishers([]string{"pedido-criado-out-0"})
	if err != nil {
		slog.Error("Error creating publisher", "error", err)
		os.Exit(1)
	}

	// @TODO: Configurar Logger
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(render.SetContentType(render.ContentTypeJSON))
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Wellcome"))
	})

	handlers.CreateRoutes(r, LoadDependencies())

	http.ListenAndServe(":8085", r)
}

func LoadDependencies() handlers.HandlerDependency {
	repository := database.CreateOrderRepository()
	orderOutput := broker.CreateOutputImp(broker.B)
	orderService := services.NewOrderService(repository, orderOutput)
	orderHandler := handlers.NewOrderHandler(orderService)

	d := handlers.HandlerDependency{
		OrderHandler: orderHandler,
	}

	return d
}
