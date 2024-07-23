package logic

import (
	"context"
	"weikang/models"
	"weikang/services/datacenter_srv/proto/datacenter"
)

func (s Server) ReportHealthDataDelete(ctx context.Context, request *datacenter.ReportHealthDataDeleteRequest) (*datacenter.ReportHealthDataDeleteResponse, error) {
	id := request.Id
	err := models.DataDelete(id)
	if err != nil {
		return &datacenter.ReportHealthDataDeleteResponse{
			Code:    500,
			Message: "数据删除失败",
		}, nil
	}
	return &datacenter.ReportHealthDataDeleteResponse{
		Code:    200,
		Message: "数据删除成功",
	}, nil
}
