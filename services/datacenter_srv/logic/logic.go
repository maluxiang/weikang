package logic

import (
	"context"
	"weikang/services/datacenter_srv/proto/datacenter"
)

type Server struct {
	datacenter.UnimplementedDatacenterServer
}

func (s Server) Ping(ctx context.Context, request *datacenter.Request) (*datacenter.Response, error) {
	//TODO implement me
	panic("implement me")
}

func (s Server) mustEmbedUnimplementedDatacenterServer() {
	//TODO implement me
	panic("implement me")
}
