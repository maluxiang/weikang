package logic

import (
	"context"
	"weikang/services/medical_svc/proto/medical"
)

type Server struct {
	medical.UnimplementedMedicalServiceServer
}

// 获取消息列表
func (s *Server) GetMessageList(ctx context.Context, req *medical.GetMessageListRequest) (*medical.GetMessageListResponse, error) {
	msgs, err := GetSiteMessages(req.UserId)
	if err != nil {
		return nil, err
	}
	var pbMsgs []*medical.SiteMessage
	for _, m := range msgs {
		pbMsgs = append(pbMsgs, &medical.SiteMessage{
			Id:        m.ID,
			UserId:    m.UserID,
			Title:     m.Title,
			Content:   m.Content,
			IsRead:    int32(m.IsRead),
			CreatedAt: m.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}
	return &medical.GetMessageListResponse{Messages: pbMsgs}, nil
}

// 标记已读
func (s *Server) MarkMessageRead(ctx context.Context, req *medical.MarkMessageReadRequest) (*medical.MarkMessageReadResponse, error) {
	err := MarkMessageRead(req.Id)
	if err != nil {
		return &medical.MarkMessageReadResponse{Msg: "操作失败"}, err
	}
	return &medical.MarkMessageReadResponse{Msg: "已标记为已读"}, nil
}
