package logic

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"weikang/common"
	"weikang/models"
	"weikang/services/user_svc/proto/user"
)

func (Server) CreateAccount(ctx context.Context, in *user.CreateAccountRequest) (*user.CreateAccountResponse, error) {
	// 实现创建用户账号的逻辑
	userId := models.Account{
		UserID: in.UserID,
	}
	info, err := userId.GetInfo()
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return &user.CreateAccountResponse{}, err
	}
	if info.ID != 0 {
		return &user.CreateAccountResponse{}, errors.New("帐户已存在")
	}
	account := models.Account{
		UserID:        in.UserID,
		AccountNumber: common.GenerateRandomString(),
		Balance:       in.Balance,
		Currency:      in.Currency,
	}
	err = account.Create()
	if err != nil {
		return &user.CreateAccountResponse{}, err
	}
	
	return &user.CreateAccountResponse{Success: true}, nil
}
