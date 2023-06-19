package svc

import (
	"github.com/myboran/zshop/app/inventory/cmd/rpc/internal/config"
	"github.com/myboran/zshop/app/inventory/inventoryModel"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config

	InventoryModel inventoryModel.InventoryModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,

		InventoryModel: inventoryModel.NewInventoryModel(sqlx.NewMysql(c.DB.DataSource)),
	}
}
