package svc

import (
	"github.com/myboran/zshop/app/order/cmd/rpc/internal/config"
	"github.com/myboran/zshop/app/order/orderGoodsModel"
	"github.com/myboran/zshop/app/order/orderInfoModel"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config

	OrderGoodsModel orderGoodsModel.OrderGoodsModel
	OrderInfoModel  orderInfoModel.OrderInfoModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,

		OrderGoodsModel: orderGoodsModel.NewOrderGoodsModel(sqlx.NewMysql(c.DB.DataSource)),
		OrderInfoModel:  orderInfoModel.NewOrderInfoModel(sqlx.NewMysql(c.DB.DataSource)),
	}
}
