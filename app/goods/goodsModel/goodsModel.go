package goodsModel

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ GoodsModel = (*customGoodsModel)(nil)

type (
	// GoodsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customGoodsModel.
	GoodsModel interface {
		goodsModel
		FindList(ctx context.Context, pageNum, pageSize int64) ([]*Goods, error)
	}

	customGoodsModel struct {
		*defaultGoodsModel
	}
)

// NewGoodsModel returns a model for the database table.
func NewGoodsModel(conn sqlx.SqlConn) GoodsModel {
	return &customGoodsModel{
		defaultGoodsModel: newGoodsModel(conn),
	}
}

func (m *defaultGoodsModel) FindList(ctx context.Context, pageNum, pageSize int64) ([]*Goods, error) {
	query := fmt.Sprintf("select %s from %s limit ?,?", goodsRows, m.table)
	var resp []*Goods
	err := m.conn.QueryRowsCtx(ctx, &resp, query, pageNum, pageSize)
	return resp, err
}
