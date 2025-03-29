package types

import (
	"context"

	"github.com/daioru/order-management/services/pkg/genproto/orders"
)

type OrderService interface {
	CreateOrder(context.Context, *orders.Order) error
	GetOrders(context.Context) []*orders.Order
}
