package logic

import (
	"context"
	"mall/service/order/model"
	"mall/service/product/rpc/productclient"
	"mall/service/user/rpc/userclient"

	"google.golang.org/grpc/status"

	"mall/service/order/rpc/internal/svc"
	"mall/service/order/rpc/order"

	"github.com/tal-tech/go-zero/core/logx"
)

type CreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateLogic) Create(in *order.CreateRequest) (*order.CreateResponse, error) {
	// 查询用户是否存在
	_, err := l.svcCtx.UserRpc.UserInfo(l.ctx, &userclient.UserInfoRequest{Id: in.Uid})
	if err != nil {
		return nil, err
	}

	// 查询产品是否存在
	productRes, err := l.svcCtx.ProductRpc.Detail(l.ctx, &productclient.DetailRequest{Id: in.Pid})

	if err != nil {
		return nil, err
	}

	// 判断库存是否存在
	if productRes.Stock <= 0 {
		return nil, status.Error(500, "库存不足")
	}

	newOrder := model.Order{
		Uid:    in.Uid,
		Pid:    in.Pid,
		Amount: in.Amount,
		Status: 0,
	}

	// 创建订单
	res, err := l.svcCtx.OrderModel.Insert(&newOrder)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	newOrder.Id, err = res.LastInsertId()
	if err != nil {
		return nil, err
	}

	// 更新产品的库存
	_, err = l.svcCtx.ProductRpc.Update(l.ctx, &productclient.UpdateRequest{
		Id:     productRes.Id,
		Name:   productRes.Name,
		Desc:   productRes.Desc,
		Stock:  productRes.Stock - 1,
		Amount: productRes.Amount,
		Status: productRes.Status,
	})

	if err != nil {
		return nil, err
	}
	return &order.CreateResponse{
		Id: newOrder.Id,
	}, nil
}
