package svc

import (
	"github.com/reation/home_user_service/internal/config"
	"github.com/zeromicro/go-zero/zrpc"
)
import "github.com/reation/home_order_service/order"

type ServiceContext struct {
	Config config.Config
	OrderRpc order.Order
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		OrderRpc: order.NewOrder(zrpc.MustNewClient(c.OrderRpc)),
	}
}
