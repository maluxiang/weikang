package logic

import (
	"context"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"weikang/models"
	"weikang/services/user_svc/proto/user"
)

func (Server) TransferCompensation(ctx context.Context, in *user.TransferCompensationRequest) (*user.TransferCompensationResponse, error) {
	// 实现转账逻辑
	fromUser := models.Account{
		UserID: in.FromID,
	}
	// 检查转账账户是否存在
	transfer, err := fromUser.GetInfo()
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return &user.TransferCompensationResponse{}, status.Errorf(codes.Aborted, "转账账户不存在")
	}
	if err != nil {
		return &user.TransferCompensationResponse{}, status.Errorf(codes.Internal, "转账账户查询失败")
	}
	toUser := models.Account{
		UserID: in.ToID,
	}
	
	// 检查收款账户是否存在
	receivables, err := toUser.GetInfo()
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return &user.TransferCompensationResponse{}, status.Errorf(codes.Aborted, "收款账户不存在")
	}
	
	if err != nil {
		return &user.TransferCompensationResponse{}, status.Errorf(codes.Internal, "收款账户查询失败")
	}
	// 检查转账账户余额是否足够
	if transfer.Balance < in.Amount {
		return &user.TransferCompensationResponse{}, status.Errorf(codes.Aborted, "转账账户余额不足")
	}
	
	// 更新转账账户余额
	transfer.Balance += in.Amount
	err = transfer.UpdateBalance()
	if err != nil {
		return &user.TransferCompensationResponse{}, status.Errorf(codes.Internal, "转账账户余额更新失败")
	}
	
	// 更新收款账户余额
	receivables.Balance -= in.Amount
	
	err = receivables.UpdateBalance()
	
	if err != nil {
		return &user.TransferCompensationResponse{}, status.Errorf(codes.Internal, "收款账户余额更新失败")
	}
	
	return &user.TransferCompensationResponse{Success: true}, nil
}
