package server

import (
	"context"
	"math"
	"rpc-practice/protogens/go/robot"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func NewRobotNavigationServer() *RobotNavigationServer {
	return &RobotNavigationServer{
		positions: map[string]robot.Position{
			"freight100-001": robot.Position{X: 0, Y: 0},
		},
	}
}

type RobotNavigationServer struct {
	positions map[string]robot.Position
}

func (srv *RobotNavigationServer) Navigate(ctx context.Context, req *robot.NavRequest) (*robot.NavResponse, error) {
	origin, ok := srv.positions[req.Robot]
	if !ok {
		return nil, status.Error(codes.InvalidArgument, "cannot find robot")
	}

	if req.Destination == nil {
		return nil, status.Error(codes.InvalidArgument, "destination is not provided")
	}

	dist := math.Pow(req.Destination.X-origin.X, 2) + math.Pow(req.Destination.Y-origin.Y, 2)
	srv.positions[req.Robot] = *req.Destination

	return &robot.NavResponse{
		DistanceTraveled: math.Sqrt(dist),
	}, nil
}
