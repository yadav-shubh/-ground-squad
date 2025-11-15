package main

import (
	"fmt"
	"github.com/yadav-shubh/base-backend/generated/grpc/modules/health/grpc"
	healthGrpc "github.com/yadav-shubh/base-backend/modules/health/server"
	"github.com/yadav-shubh/base-backend/utils"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {

	addr := fmt.Sprintf("0.0.0.0:%d", 8000)
	tcp, _ := net.ResolveTCPAddr("tcp", addr)
	listener, err := net.ListenTCP("tcp", tcp)
	if err != nil {
		utils.Log.Error("Unable to resolve tcp", zap.String("address", addr))
	}
	grpcServer := grpc.NewServer()

	reflection.Register(grpcServer)

	health.RegisterHealthServer(grpcServer, healthGrpc.NewHealthServer())

	err = grpcServer.Serve(listener)
	if err != nil {
		utils.Log.Error("Unable to serve", zap.String("address", addr))
	}
	utils.Log.Info("Server started", zap.String("address", addr))

}
