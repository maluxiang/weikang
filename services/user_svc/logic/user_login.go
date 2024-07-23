package logic

import (
	"context"
	"errors"
	"weikang/common"
	"weikang/models"
	"weikang/services/user_svc/proto/user"
)

func (Server) Login(ctx context.Context, in *user.LoginRequest) (*user.LoginResponse, error) {
	u := models.User{
		Username: in.User.Username,
		Password: common.MD5(in.User.Password),
	}
	info, err := u.GetInfo()
	if err != nil {
		return nil, err
	}
	if info.ID == 0 {
		return &user.LoginResponse{}, errors.New("用户名或密码错误")
	}
	return &user.LoginResponse{UserId: int64(info.ID)}, nil
}
