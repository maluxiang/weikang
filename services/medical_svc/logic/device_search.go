package logic

import (
	"context"
	"weikang/pkg"
	"weikang/services/medical_svc/proto/medical"
)

func (s Server) SearchDevice(ctx context.Context, in *medical.SearchDeviceRequest) (*medical.SearchDeviceResponse, error) {
	devices, err := pkg.SearchDeviceFromES(in.Keyword)
	if err != nil {
		return &medical.SearchDeviceResponse{}, err
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
	return &medical.SearchDeviceResponse{
		Devices: respDevices,
		Total:   uint32(len(respDevices)),
	}, nil
}
