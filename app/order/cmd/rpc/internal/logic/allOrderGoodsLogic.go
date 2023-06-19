package logic

import (
	"context"

	"github.com/myboran/zshop/app/order/cmd/rpc/internal/svc"
	"github.com/myboran/zshop/app/order/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AllOrderGoodsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAllOrderGoodsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AllOrderGoodsLogic {
	return &AllOrderGoodsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AllOrderGoodsLogic) AllOrderGoods(in *pb.OrderGoodsReq) (*pb.OrderGoodsResp, error) {
	orderGoods, err := l.svcCtx.OrderGoodsModel.FindOrderBySn(l.ctx, in.OrderSn)
	if err != nil {
		return nil, err
	}
	resp := &pb.OrderGoodsResp{Orders: make([]*pb.OrderGoods, len(orderGoods))}
	for i, v := range orderGoods {
		resp.Orders[i] = new(pb.OrderGoods)
		resp.Orders[i].Id = v.Id
		resp.Orders[i].Goods = v.Goods
		resp.Orders[i].OrderSn = v.OrderSn
		resp.Orders[i].Nums = v.Nums
		resp.Orders[i].Price = float32(v.Price)
	}
	return resp, nil
}
