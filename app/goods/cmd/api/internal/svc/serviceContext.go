package svc

import (
	"github.com/myboran/zshop/app/goods/cmd/api/internal/config"
	"github.com/myboran/zshop/app/goods/cmd/rpc/goods"
	"github.com/myboran/zshop/app/inventory/cmd/rpc/inventory"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config

	GoodsRpc     goods.Goods
	InventoryRpc inventory.Inventory
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,

		GoodsRpc:     goods.NewGoods(zrpc.MustNewClient(c.GoodsRpcConf)),
		InventoryRpc: inventory.NewInventory(zrpc.MustNewClient(c.InventoryRpcConf)),
	}
}
