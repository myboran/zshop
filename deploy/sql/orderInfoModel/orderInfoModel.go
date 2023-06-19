package orderInfoModel

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ OrderInfoModel = (*customOrderInfoModel)(nil)

type (
	// OrderInfoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customOrderInfoModel.
	OrderInfoModel interface {
		orderInfoModel
	}

	customOrderInfoModel struct {
		*defaultOrderInfoModel
	}
)

// NewOrderInfoModel returns a model for the database table.
func NewOrderInfoModel(conn sqlx.SqlConn) OrderInfoModel {
	return &customOrderInfoModel{
		defaultOrderInfoModel: newOrderInfoModel(conn),
	}
}
