package server

import pb "github.com/yadav-shubh/base-backend/generated/grpc/modules/health/grpc"

type HealthServer struct {
	pb.UnimplementedHealthServer
}

func NewHealthServer() *HealthServer {
	return &HealthServer{
		UnimplementedHealthServer: pb.UnimplementedHealthServer{},
	}
}
