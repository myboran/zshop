// Code generated by goctl. DO NOT EDIT.
// Source: goods.proto

package server

import (
	"context"

	"github.com/myboran/zshop/app/goods/cmd/rpc/internal/logic"
	"github.com/myboran/zshop/app/goods/cmd/rpc/internal/svc"
	"github.com/myboran/zshop/app/goods/cmd/rpc/pb"
)

type GoodsServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedGoodsServer
}

func NewGoodsServer(svcCtx *svc.ServiceContext) *GoodsServer {
	return &GoodsServer{
		svcCtx: svcCtx,
	}
}

func (s *GoodsServer) GoodsList(ctx context.Context, in *pb.GoodsListReq) (*pb.GoodsListResp, error) {
	l := logic.NewGoodsListLogic(ctx, s.svcCtx)
	return l.GoodsList(in)
}

func (s *GoodsServer) GoodsDetail(ctx context.Context, in *pb.GoodDetailReq) (*pb.GoodsInfo, error) {
	l := logic.NewGoodsDetailLogic(ctx, s.svcCtx)
	return l.GoodsDetail(in)
}