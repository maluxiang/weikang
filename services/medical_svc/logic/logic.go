package logic

import (
	"weikang/services/medical_svc/proto/medical"
)

type Server struct {
	medical.UnimplementedMedicalServiceServer
}
