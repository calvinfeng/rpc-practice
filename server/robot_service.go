package server

import (
	"context"
	"math"
	"rpc-practice/protogens/go/robot"
)

type RobotMovementService struct{}

func (srv *RobotMovementService) Move(ctx context.Context, req *robot.MoveRequest) (*robot.MoveResponse, error) {
	dist := math.Pow(req.Target.X-req.Origin.X, 2) + math.Pow(req.Target.Y-req.Origin.Y, 2)
	return &robot.MoveResponse{
		Distance: math.Sqrt(dist),
	}, nil
}
