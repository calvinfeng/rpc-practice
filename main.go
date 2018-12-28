package main

import (
	"fmt"
	"net"
	"rpc-practice/protogens/go/robot"
	"rpc-practice/server"

	"google.golang.org/grpc"
)

func RunGRPCServer() error {
	lis, err := net.Listen("tcp", ":3000")
	if err != nil {
		return err
	}

	s := grpc.NewServer()
	robot.RegisterMovementServiceServer(s, &server.RobotMovementService{})

	go s.Serve(lis)
	return nil
}

func main() {
	fmt.Println("Let's do gRPC")
}
