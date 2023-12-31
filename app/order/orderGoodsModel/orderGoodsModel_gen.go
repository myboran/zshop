// Code generated by goctl. DO NOT EDIT.

package orderGoodsModel

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	orderGoodsFieldNames          = builder.RawFieldNames(&OrderGoods{})
	orderGoodsRows                = strings.Join(orderGoodsFieldNames, ",")
	orderGoodsRowsExpectAutoSet   = strings.Join(stringx.Remove(orderGoodsFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	orderGoodsRowsWithPlaceHolder = strings.Join(stringx.Remove(orderGoodsFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	orderGoodsModel interface {
		Insert(ctx context.Context, data *OrderGoods) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*OrderGoods, error)
		Update(ctx context.Context, data *OrderGoods) error
		Delete(ctx context.Context, id int64) error
	}

	defaultOrderGoodsModel struct {
		conn  sqlx.SqlConn
		table string
	}

	OrderGoods struct {
		Id      int64   `db:"id"`
		OrderSn string  `db:"order_sn"` // 订单号
		Goods   int64   `db:"goods"`    // 商品id
		Price   float64 `db:"price"`    // 价格
		Nums    int64   `db:"nums"`     // 商品数量
	}
)

func newOrderGoodsModel(conn sqlx.SqlConn) *defaultOrderGoodsModel {
	return &defaultOrderGoodsModel{
		conn:  conn,
		table: "`order_goods`",
	}
}

func (m *defaultOrderGoodsModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultOrderGoodsModel) FindOne(ctx context.Context, id int64) (*OrderGoods, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", orderGoodsRows, m.table)
	var resp OrderGoods
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultOrderGoodsModel) Insert(ctx context.Context, data *OrderGoods) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, orderGoodsRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.OrderSn, data.Goods, data.Price, data.Nums)
	return ret, err
}

func (m *defaultOrderGoodsModel) Update(ctx context.Context, data *OrderGoods) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, orderGoodsRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.OrderSn, data.Goods, data.Price, data.Nums, data.Id)
	return err
}

func (m *defaultOrderGoodsModel) tableName() string {
	return m.table
}
