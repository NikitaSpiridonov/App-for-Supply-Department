package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/google/uuid"

	orderV1 "github.com/NikitaSpiridonov/App-for-Supply-Department/shared/pkg/openapi/order/v1"
)

const (
	httpPort          = "8080"
	readHeaderTimeout = 5 * time.Second
	shutdownTimeout   = 10 * time.Second
)

type PaymentStatus string

const (
	StatusPending  PaymentStatus = "PENDING_PAYMENT"
	StatusPaid     PaymentStatus = "PAID"
	StatusCanceled PaymentStatus = "CANCELED"
)

type Order struct {
	user_uuid        uuid.UUID
	part_uuids       []uuid.UUID
	total_price      float64
	transaction_uuid *uuid.UUID
	payment_method   *string
	status           PaymentStatus
}

type OrderStorage struct {
	orders map[uuid.UUID]Order
	mu     sync.Mutex
}

func newOrderStorage() *OrderStorage {
	return &OrderStorage{
		orders: make(map[uuid.UUID]Order),
	}
}

type OrderHandler struct {
	storage *OrderStorage
}

func NewOrderHandler(storage *OrderStorage) *OrderHandler {
	return &OrderHandler{
		storage: storage,
	}
}

func (h *OrderHandler) CreateNewOrder(ctx context.Context, req *orderV1.CreateNewOrderReq) (r orderV1.CreateNewOrderRes, _ error) {
	orderUUID := uuid.New()
	partUUIDs := make([]uuid.UUID, len(req.PartUuids))
	for i, pu := range req.PartUuids {
		partUUIDs[i] = uuid.MustParse(string(pu))
	}

}
