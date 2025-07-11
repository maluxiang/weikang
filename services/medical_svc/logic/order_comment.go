package logic

import (
	"context"
	"weikang/global"
	"weikang/models"
	"weikang/pkg"
	"weikang/services/medical_svc/proto/medical"
)

func (s Server) AddOrderComment(ctx context.Context, in *medical.AddOrderCommentRequest) (*medical.AddOrderCommentResponse, error) {
	//检查订单是否存在且属于该用户
	var order models.Order
	if err := global.DB.Where("id = ? AND user_id = ?", in.OrderId, in.UserId).First(&order).Error; err != nil {
		return &medical.AddOrderCommentResponse{Success: false, Message: "订单不存在或无权限"}, nil
	}

	if order.UserID != in.UserId || order.PaymentStatus != 1 {
		return &medical.AddOrderCommentResponse{Success: false, Message: "评论失败,没有权限去评论"}, nil
	}

	//违规评论判断
	violation := pkg.Violation(in.Content)
	if violation.ConclusionType != 1 {
		return &medical.AddOrderCommentResponse{Success: false, Message: "评论内容存在不当言论，评论失败"}, nil
	}

	//评论次数
	if !pkg.CanComment(in.UserId) {
		return &medical.AddOrderCommentResponse{
			Success: false,
			Message: "评论过于频繁，请稍后再试",
		}, nil
	}

	//创建评论
	comment := models.OrderComment{
		OrderID: in.OrderId,
		UserID:  in.UserId,
		Content: in.Content,
	}
	if err := global.DB.Create(&comment).Error; err != nil {
		return &medical.AddOrderCommentResponse{Success: false, Message: "评论失败"}, nil
	}

	return &medical.AddOrderCommentResponse{Success: true, Message: "评论成功"}, nil
}
