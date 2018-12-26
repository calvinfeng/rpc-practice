package server

import (
	"context"
	"rpc-practice/cprotos/golang/robot"
)

type RobotMovementService struct{}

func (srv *RobotMovementService) Move(ctx context.Context, req *robot.MoveRequest) (*robot.MoveResponse, error) {
	return &robot.MoveResponse{}, nil
}
