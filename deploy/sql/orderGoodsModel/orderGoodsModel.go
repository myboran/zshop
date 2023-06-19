package orderGoodsModel

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ OrderGoodsModel = (*customOrderGoodsModel)(nil)

type (
	// OrderGoodsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customOrderGoodsModel.
	OrderGoodsModel interface {
		orderGoodsModel
	}

	customOrderGoodsModel struct {
		*defaultOrderGoodsModel
	}
)

// NewOrderGoodsModel returns a model for the database table.
func NewOrderGoodsModel(conn sqlx.SqlConn) OrderGoodsModel {
	return &customOrderGoodsModel{
		defaultOrderGoodsModel: newOrderGoodsModel(conn),
	}
}
