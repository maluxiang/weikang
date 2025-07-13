package logic

import (
	"context"
	"weikang/global"
	"weikang/models"
	"weikang/pkg"
	"weikang/services/medical_svc/proto/medical"
)

// AddOrderComment 添加订单评论 - 支持无限评论
func (s Server) AddOrderComment(ctx context.Context, in *medical.AddOrderCommentRequest) (*medical.AddOrderCommentResponse, error) {
	//检查订单是否存在且属于该用户
	var order models.Order
	if err := global.DB.Where("id = ? AND user_id = ?", in.OrderId, in.UserId).First(&order).Error; err != nil {
		return &medical.AddOrderCommentResponse{Success: false, Message: "订单不存在或无权限"}, nil
	}

	// 检查订单支付状态 - 只有已支付的订单才能评论
	if order.PaymentStatus != 1 {
		return &medical.AddOrderCommentResponse{Success: false, Message: "只有已支付的订单才能评论"}, nil
	}

	//违规评论判断
	violation := pkg.Violation(in.Content)
	if violation.ConclusionType != 1 {
		return &medical.AddOrderCommentResponse{Success: false, Message: "评论内容存在不当言论，评论失败"}, nil
	}

	// 如果是回复评论，检查回复的评论是否存在
	if in.ReplyTo > 0 {
		var replyComment models.OrderComment
		if err := global.DB.Where("id = ? AND order_id = ?", in.ReplyTo, in.OrderId).First(&replyComment).Error; err != nil {
			return &medical.AddOrderCommentResponse{Success: false, Message: "回复的评论不存在"}, nil
		}
	}

	//创建评论
	comment := models.OrderComment{
		OrderID:     in.OrderId,
		UserID:      in.UserId,
		Content:     in.Content,
		Rating:      int(in.Rating),
		Images:      in.Images,
		ReplyTo:     in.ReplyTo,
		IsAnonymous: in.IsAnonymous,
		Status:      1, // 正常状态
		Likes:       0,
	}

	if err := global.DB.Create(&comment).Error; err != nil {
		return &medical.AddOrderCommentResponse{Success: false, Message: "评论失败"}, nil
	}

	// 获取用户信息
	var user models.Users
	global.DB.Where("id = ?", in.UserId).First(&user)

	// 构建返回的评论信息
	commentInfo := &medical.CommentInfo{
		Id:          comment.ID,
		OrderId:     comment.OrderID,
		UserId:      comment.UserID,
		Content:     comment.Content,
		Rating:      int32(comment.Rating),
		Images:      comment.Images,
		ReplyTo:     comment.ReplyTo,
		IsAnonymous: comment.IsAnonymous,
		Status:      int32(comment.Status),
		Likes:       int32(comment.Likes),
		CreatedAt:   comment.CreatedAt.Format("2006-01-02 15:04:05"),
		UserName:    user.Username,
		IsLiked:     false,
	}

	return &medical.AddOrderCommentResponse{
		Success: true,
		Message: "评论成功",
		Comment: commentInfo,
	}, nil
}

// GetOrderComments 获取订单评论列表
func (s Server) GetOrderComments(ctx context.Context, in *medical.GetOrderCommentsRequest) (*medical.GetOrderCommentsResponse, error) {
	// 检查订单是否存在
	var order models.Order
	if err := global.DB.Where("id = ?", in.OrderId).First(&order).Error; err != nil {
		return &medical.GetOrderCommentsResponse{}, nil
	}

	// 获取评论列表
	var commentModel models.OrderComment
	comments, total, err := commentModel.GetOrderComments(in.OrderId, int(in.Page), int(in.PageSize))
	if err != nil {
		return &medical.GetOrderCommentsResponse{}, err
	}

	// 构建返回数据
	var commentInfos []*medical.CommentInfo
	for _, comment := range comments {
		commentInfo := &medical.CommentInfo{
			Id:          comment.ID,
			OrderId:     comment.OrderID,
			UserId:      comment.UserID,
			Content:     comment.Content,
			Rating:      int32(comment.Rating),
			Images:      comment.Images,
			ReplyTo:     comment.ReplyTo,
			IsAnonymous: comment.IsAnonymous,
			Status:      int32(comment.Status),
			Likes:       int32(comment.Likes),
			CreatedAt:   comment.CreatedAt.Format("2006-01-02 15:04:05"),
			UserName:    comment.User.Username,
			IsLiked:     false, // 这里可以根据当前用户ID判断是否点赞
		}
		commentInfos = append(commentInfos, commentInfo)
	}

	return &medical.GetOrderCommentsResponse{
		Comments: commentInfos,
		Total:    int32(total),
		Page:     in.Page,
		PageSize: in.PageSize,
	}, nil
}

// GetCommentReplies 获取评论回复
func (s Server) GetCommentReplies(ctx context.Context, in *medical.GetCommentRepliesRequest) (*medical.GetCommentRepliesResponse, error) {
	// 检查评论是否存在
	var comment models.OrderComment
	if err := global.DB.Where("id = ?", in.CommentId).First(&comment).Error; err != nil {
		return &medical.GetCommentRepliesResponse{}, nil
	}

	// 获取回复列表
	var commentModel models.OrderComment
	replies, total, err := commentModel.GetCommentReplies(in.CommentId, int(in.Page), int(in.PageSize))
	if err != nil {
		return &medical.GetCommentRepliesResponse{}, err
	}

	// 构建返回数据
	var replyInfos []*medical.CommentInfo
	for _, reply := range replies {
		replyInfo := &medical.CommentInfo{
			Id:          reply.ID,
			OrderId:     reply.OrderID,
			UserId:      reply.UserID,
			Content:     reply.Content,
			Rating:      int32(reply.Rating),
			Images:      reply.Images,
			ReplyTo:     reply.ReplyTo,
			IsAnonymous: reply.IsAnonymous,
			Status:      int32(reply.Status),
			Likes:       int32(reply.Likes),
			CreatedAt:   reply.CreatedAt.Format("2006-01-02 15:04:05"),
			UserName:    reply.User.Username,
			IsLiked:     false,
		}
		replyInfos = append(replyInfos, replyInfo)
	}

	return &medical.GetCommentRepliesResponse{
		Replies:  replyInfos,
		Total:    int32(total),
		Page:     in.Page,
		PageSize: in.PageSize,
	}, nil
}

// LikeComment 点赞评论
func (s Server) LikeComment(ctx context.Context, in *medical.LikeCommentRequest) (*medical.LikeCommentResponse, error) {
	// 检查评论是否存在
	var comment models.OrderComment
	if err := global.DB.Where("id = ?", in.CommentId).First(&comment).Error; err != nil {
		return &medical.LikeCommentResponse{Success: false, Message: "评论不存在"}, nil
	}

	// 执行点赞操作
	var likeModel models.CommentLike
	err := likeModel.LikeComment(in.CommentId, in.UserId)
	if err != nil {
		return &medical.LikeCommentResponse{Success: false, Message: "点赞失败"}, nil
	}

	// 检查当前点赞状态
	isLiked := likeModel.IsLiked(in.CommentId, in.UserId)

	// 获取最新的点赞数
	global.DB.Where("id = ?", in.CommentId).First(&comment)

	return &medical.LikeCommentResponse{
		Success:    true,
		Message:    "操作成功",
		IsLiked:    isLiked,
		LikesCount: int32(comment.Likes),
	}, nil
}

// DeleteComment 删除评论
func (s Server) DeleteComment(ctx context.Context, in *medical.DeleteCommentRequest) (*medical.DeleteCommentResponse, error) {
	// 检查评论是否存在且属于该用户
	var comment models.OrderComment
	if err := global.DB.Where("id = ? AND user_id = ?", in.CommentId, in.UserId).First(&comment).Error; err != nil {
		return &medical.DeleteCommentResponse{Success: false, Message: "评论不存在或无权限删除"}, nil
	}

	// 软删除评论
	if err := global.DB.Model(&comment).Update("status", 3).Error; err != nil {
		return &medical.DeleteCommentResponse{Success: false, Message: "删除失败"}, nil
	}

	return &medical.DeleteCommentResponse{Success: true, Message: "删除成功"}, nil
}
