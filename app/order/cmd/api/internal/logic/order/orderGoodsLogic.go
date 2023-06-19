package order

import (
	"context"
	"github.com/myboran/zshop/app/order/cmd/rpc/pb"

	"github.com/myboran/zshop/app/order/cmd/api/internal/svc"
	"github.com/myboran/zshop/app/order/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderGoodsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOrderGoodsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderGoodsLogic {
	return &OrderGoodsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OrderGoodsLogic) OrderGoods(req *types.OrderGoodsReq) (resp *types.OrderGoodsResp, err error) {
	orders, err := l.svcCtx.OrderRpc.AllOrderGoods(l.ctx, &pb.OrderGoodsReq{OrderSn: req.OrderSn})
	if err != nil {
		return nil, err
	}
	resp = &types.OrderGoodsResp{Goods: make([]types.OrderGoods, len(orders.Orders))}
	for i, v := range orders.Orders {
		resp.Goods[i].Id = v.Id
		resp.Goods[i].OrderSn = v.OrderSn
		resp.Goods[i].Goods = v.Goods
		resp.Goods[i].Price = float64(v.Price)
		resp.Goods[i].Nums = v.Nums
	}
	return resp, nil
}
