package logic

import (
	"context"
	"weikang/common"
	"weikang/models"
	"weikang/pkg"
	"weikang/services/user_svc/proto/user"
)

func (Server) Register(ctx context.Context, in *user.RegisterRequest) (*user.RegisterResponse, error) {
	// 实现用户注册的逻辑
	u := models.Users{
		Username: in.User.Username,
		Password: common.MD5(in.User.Password),
		Email:    in.User.Email,
		Phone:    in.User.Phone,
	}
	err := u.Create()
	if err != nil {
		return nil, err
	}

	pkg.Publish("2210a", u.Username)

	return &user.RegisterResponse{Success: true}, nil
}
