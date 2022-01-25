package logic

import (
	"context"
	"mall/common/cryptx"
	"mall/service/user/model"

	"google.golang.org/grpc/status"

	"mall/service/user/rpc/internal/svc"
	"mall/service/user/rpc/user"

	"github.com/tal-tech/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 注册用户
func (l *RegisterLogic) Register(in *user.RegisterRequest) (*user.RegisterResponse, error) {
	_, err := l.svcCtx.UserModel.FindOneByMobile(in.Mobile)
	if err == nil {
		return nil, status.Error(100, "该用户已经存在")
	}

	if err == model.ErrNotFound {
		newUser := model.User{
			Name:     in.Name,
			Mobile:   in.Mobile,
			Gender:   in.Gender,
			Password: cryptx.PasswordEncrypt(l.svcCtx.Config.Salt, in.Password),
		}
		res, err := l.svcCtx.UserModel.Insert(&newUser)

		if err != nil {
			return nil, status.Error(500, err.Error())
		}

		newUser.Id, err = res.LastInsertId()

		if err != nil {
			return nil, status.Error(500, err.Error())
		}

		return &user.RegisterResponse{
			Id:     newUser.Id,
			Name:   newUser.Name,
			Gender: newUser.Gender,
			Mobile: newUser.Mobile,
		}, nil
	}

	return nil, status.Error(500, err.Error())
}
