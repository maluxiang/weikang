package logic

import (
	"golang.org/x/net/context"
	"weikang/services/points_svc/proto/points"
)

type Server struct {
	points.UnimplementedPointsServer
}

func (s Server) DeletePoints(ctx context.Context, in *points.DeletePointsRequest) (*points.DeletePointsResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s Server) mustEmbedUnimplementedPointsServer() {
	//TODO implement me
	panic("implement me")
}
