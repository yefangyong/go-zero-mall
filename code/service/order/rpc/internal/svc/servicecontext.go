package svc

import (
	"mall/service/order/model"
	"mall/service/order/rpc/internal/config"
	"mall/service/product/rpc/productclient"
	"mall/service/user/rpc/userclient"

	"github.com/tal-tech/go-zero/zrpc"

	"github.com/tal-tech/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config     config.Config
	OrderModel model.OrderModel
	UserRpc    userclient.User
	ProductRpc productclient.Product
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:     c,
		OrderModel: model.NewOrderModel(conn, c.CacheRedis),
		UserRpc:    userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
		ProductRpc: productclient.NewProduct(zrpc.MustNewClient(c.ProductRpc)),
	}
}
