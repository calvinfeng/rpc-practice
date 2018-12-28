package server

import (
	"context"
	"fmt"
	"net"
	"rpc-practice/protogens/go/robot"
	"testing"
	"time"

	"google.golang.org/grpc"
)

// GetFreeAddr returns a free address and port ready to use.
func GetFreeAddr(t *testing.T) (int, string) {
	addr := &net.TCPAddr{
		IP:   net.IPv4(127, 0, 0, 1).To4(),
		Port: 0,
	}

	list, err := net.ListenTCP("tcp", addr)
	if err != nil {
		t.Error(err)
	}

	defer list.Close()

	port := list.Addr().(*net.TCPAddr).Port

	if err != nil {
		t.Fatal(err)
	}

	return port, fmt.Sprintf(":%d", port)
}

func SetupGRPCServer(t *testing.T, addr string) {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		t.Error(err)
	}

	s := grpc.NewServer()
	robot.RegisterNavigationServer(s, NewRobotNavigationServer())

	go s.Serve(lis)
}

func TestRobotMovementService(t *testing.T) {
	port, addr := GetFreeAddr(t)
	SetupGRPCServer(t, addr)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	hostname := "localhost"
	url := fmt.Sprintf("%s:%d", hostname, port)

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())

	conn, err := grpc.DialContext(ctx, url, opts...)
	if err != nil {
		t.Errorf("fail to dial: %v", err)
	}
	defer conn.Close()

	client := robot.NewNavigationClient(conn)

	t.Run("NavigateSuccess", func(t *testing.T) {
		res, err := client.Navigate(ctx, &robot.NavRequest{
			Robot: "freight100-001",
			Destination: &robot.Position{
				X: 15,
				Y: 15,
			},
		})

		if err != nil {
			t.Error(err)
		}

		if res.DistanceTraveled == 0 {
			t.Error("distance should not be zero valued")
		}
	})

	t.Run("NavigateFail", func(t *testing.T) {
		res, err := client.Navigate(ctx, &robot.NavRequest{
			Robot: "freight100-002",
			Destination: &robot.Position{
				X: 15,
				Y: 15,
			},
		})

		if err == nil {
			t.Error("gRPC response should return an error")
		}

		fmt.Println(err)

		if res != nil {
			t.Error("gRPC response should be nil")
		}
	})
}
