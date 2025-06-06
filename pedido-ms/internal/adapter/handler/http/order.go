package http

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"pedido-ms/internal/adapter/database"
	"pedido-ms/internal/adapter/dto"
	"pedido-ms/internal/core/domain"
	"pedido-ms/internal/core/port"
	"pedido-ms/internal/core/services"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

func CreateRoutes(r *chi.Mux) {
	r.Route("/pedidos", func(r chi.Router) {
		repository := database.CreateOrderRepository()
		service := services.NewOrderService(&repository)
		handler := NewOrderHandler(service)
		r.Post("/", handler.CreateOrder)
	})
}

type OrderHandler struct {
	service port.OrderService
}

func NewOrderHandler(svc port.OrderService) *OrderHandler {
	return &OrderHandler{
		service: svc,
	}
}

type (
	OrderRequest struct {
		*dto.OrderDTO
	}

	OrderResponse struct {
		Id          string  `json:"id" example:"abcde"`
		TotalAmount float64 `json:"totalAmount" example:"1.00"`
	}
)

func (or *OrderRequest) Bind(r *http.Request) error {
	log.Println("[OrderRequest.Bind] input fields:", or.OrderDTO)
	if or.OrderDTO == nil {
		return errors.New("missing required order fields")
	}

	return nil
}

func NewOrderResponse(o *domain.Order) *OrderResponse {
	resp := &OrderResponse{
		Id:          o.ID(),
		TotalAmount: o.TotalAmount(),
	}

	log.Println("[NewOrderResponse] response:", resp)

	return resp
}

func (rd *OrderResponse) Render(w http.ResponseWriter, r *http.Request) error {

	if rd.Id == "" || rd.TotalAmount == 0 {
		log.Println("Render: campos não preenchidos corretamente")
	}

	return nil
}

func (oh *OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	data := &OrderRequest{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
	}

	order := data.OrderDTO
	log.Println("[Create Order Handler] dto: ", order)
	ctx := r.Context()
	output, err := oh.service.Create(&ctx, order.ToEntity())

	if err != nil {
		ErrInvalidRequest(err)
	}

	response := NewOrderResponse(output)

	fmt.Println("response", response)

	render.Status(r, http.StatusCreated)
	render.Render(w, r, response)
}
