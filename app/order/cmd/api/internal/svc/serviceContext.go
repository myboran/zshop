package svc

import (
	"github.com/myboran/zshop/app/order/cmd/api/internal/config"
	"github.com/myboran/zshop/app/order/cmd/rpc/order"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config

	OrderRpc order.Order
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,

		OrderRpc: order.NewOrder(zrpc.MustNewClient(c.OrderRpcConf)),
	}
}
