package logic

import (
	"context"

	"github.com/myboran/zshop/app/goods/cmd/rpc/internal/svc"
	"github.com/myboran/zshop/app/goods/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GoodsListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGoodsListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GoodsListLogic {
	return &GoodsListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GoodsListLogic) GoodsList(in *pb.GoodsListReq) (*pb.GoodsListResp, error) {
	goods, err := l.svcCtx.GoodsModel.FindList(l.ctx, in.GetPageNum(), in.GetPageSize())
	if err != nil {
		return nil, err
	}
	resp := &pb.GoodsListResp{
		Goods: make([]*pb.GoodsInfo, len(goods)),
	}
	for i, v := range goods {
		resp.Goods[i] = new(pb.GoodsInfo)
		resp.Goods[i].Id = v.Id
		resp.Goods[i].Name = v.Name
		resp.Goods[i].Price = float32(v.Price)
	}
	return resp, nil
}
