package logic

import (
	"context"

	"github.com/myboran/zshop/app/goods/cmd/rpc/internal/svc"
	"github.com/myboran/zshop/app/goods/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GoodsDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGoodsDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GoodsDetailLogic {
	return &GoodsDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GoodsDetailLogic) GoodsDetail(in *pb.GoodDetailReq) (*pb.GoodsInfo, error) {
	goodsInfo, err := l.svcCtx.GoodsModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}
	return &pb.GoodsInfo{
		Id:    goodsInfo.Id,
		Name:  goodsInfo.Name,
		Price: float32(goodsInfo.Price),
	}, nil
}
