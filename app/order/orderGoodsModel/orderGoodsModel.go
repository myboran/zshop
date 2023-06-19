package orderGoodsModel

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ OrderGoodsModel = (*customOrderGoodsModel)(nil)

type (
	// OrderGoodsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customOrderGoodsModel.
	OrderGoodsModel interface {
		orderGoodsModel
		FindOrderBySn(ctx context.Context, sn string) ([]*OrderGoods, error)
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

func (m *defaultOrderGoodsModel) FindOrderBySn(ctx context.Context, sn string) ([]*OrderGoods, error) {
	query := fmt.Sprintf("select %s from %s where `order_sn` = ?", orderGoodsRows, m.table)
	var resp []*OrderGoods
	err := m.conn.QueryRowsCtx(ctx, &resp, query, sn)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
