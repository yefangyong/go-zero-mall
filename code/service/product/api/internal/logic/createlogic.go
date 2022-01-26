package logic

import (
	"context"
	"mall/service/product/rpc/productclient"

	"mall/service/product/api/internal/svc"
	"mall/service/product/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type CreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) CreateLogic {
	return CreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// 创建产品
func (l *CreateLogic) Create(req types.CreateRequest) (resp *types.CreateResponse, err error) {
	res, err := l.svcCtx.ProductRpc.Create(l.ctx, &productclient.CreateRequest{
		Name:   req.Name,
		Desc:   req.Desc,
		Stock:  req.Stock,
		Amount: req.Amount,
		Status: req.Status,
	})
	if err != nil {
		return nil, err
	}

	return &types.CreateResponse{
		Id: res.Id,
	}, nil
}
