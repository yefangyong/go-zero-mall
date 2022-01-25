package logic

import (
	"context"
	"mall/service/product/rpc/productclient"

	"mall/service/product/api/internal/svc"
	"mall/service/product/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type RemoveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRemoveLogic(ctx context.Context, svcCtx *svc.ServiceContext) RemoveLogic {
	return RemoveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// 删除产品
func (l *RemoveLogic) Remove(req types.RemoveRequest) (resp *types.RemoveResponse, err error) {
	_, err = l.svcCtx.ProductRpc.Remove(l.ctx, &productclient.RemoveRequest{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	return &types.RemoveResponse{}, nil
}