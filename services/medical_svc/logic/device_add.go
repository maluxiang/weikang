package logic

import (
	"context"
	"weikang/global"
	"weikang/models"
	"weikang/pkg"
	"weikang/services/medical_svc/proto/medical"
)

func (s Server) AddDevice(ctx context.Context, in *medical.AddDeviceRequest) (*medical.AddDeviceResponse, error) {
	device := &models.MedicalDevice{
		Name:        in.Name,
		Brand:       in.Brand,
		Model:       in.Model,
		Description: in.Description,
		Stock:       int(in.Stock),
		Price:       float64(in.Price),
	}
	if err := global.DB.Create(device).Error; err != nil {
		return &medical.AddDeviceResponse{Success: false, Message: err.Error()}, err
	}
	if err := pkg.SyncDeviceToES(device); err != nil {
		return &medical.AddDeviceResponse{Success: false, Message: err.Error()}, err
	}
	return &medical.AddDeviceResponse{
		Success: true,
		Message: "添加成功",
	}, nil
}
