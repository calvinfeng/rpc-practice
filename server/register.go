package server

import (
	"rpc-practice/protogens/go/robot"

	"google.golang.org/grpc"
)

func Register(s *grpc.Server) {
	robot.RegisterMovementServiceServer(s, &RobotMovementService{})
}
