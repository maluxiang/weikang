package logic

import (
	"context"
	"weikang/global"
	"weikang/models"
	"weikang/services/medical_svc/proto/medical"
)

func (s Server) ListDevice(ctx context.Context, in *medical.ListDeviceRequest) (*medical.ListDeviceResponse, error) {
	var devices []models.MedicalDevice
	err := global.DB.Find(&devices).Error
	if err != nil {
		return &medical.ListDeviceResponse{}, err
	}
	var respDevices []*medical.MedicalDevice
	for _, d := range devices {
		respDevices = append(respDevices, &medical.MedicalDevice{
			Id:          d.ID,
			Name:        d.Name,
			Brand:       d.Brand,
			Model:       d.Model,
			Stock:       int64(d.Stock),
			Price:       float32(d.Price),
			Description: d.Description,
		})
	}
	return &medical.ListDeviceResponse{
		Devices: respDevices,
		Total:   uint32(len(respDevices)),
	}, nil
}
