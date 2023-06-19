package svc

import (
	"github.com/myboran/zshop/app/goods/cmd/rpc/internal/config"
	"github.com/myboran/zshop/app/goods/goodsModel"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config

	GoodsModel goodsModel.GoodsModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,

		GoodsModel: goodsModel.NewGoodsModel(sqlx.NewMysql(c.DB.DataSource)),
	}
}
