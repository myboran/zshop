package orderInfoModel

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ OrderInfoModel = (*customOrderInfoModel)(nil)

type (
	// OrderInfoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customOrderInfoModel.
	OrderInfoModel interface {
		orderInfoModel

		FindOrderByUser(ctx context.Context, user int64) ([]*OrderInfo, error)
		FindOrders(ctx context.Context) ([]*OrderInfo, error)
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

func (m *defaultOrderInfoModel) FindOrderByUser(ctx context.Context, user int64) ([]*OrderInfo, error) {
	query := fmt.Sprintf("select %s from %s where `user` = ?", orderInfoRows, m.table)
	var resp []*OrderInfo
	err := m.conn.QueryRowsCtx(ctx, &resp, query, user)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *defaultOrderInfoModel) FindOrders(ctx context.Context) ([]*OrderInfo, error) {
	query := fmt.Sprintf("select %s from %s", orderInfoRows, m.table)
	var resp []*OrderInfo
	err := m.conn.QueryRowsCtx(ctx, &resp, query)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
