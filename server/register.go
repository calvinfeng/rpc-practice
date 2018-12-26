package server

import (
	"rpc-practice/cprotos/golang/robot"

	"google.golang.org/grpc"
)

func Register(s *grpc.Server) {
	robot.RegisterMovementServiceServer(s, &RobotMovementService{})
}
