package server

import (
	"context"
	"fmt"
	"net"
	"rpc-practice/cprotos/golang/robot"
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
	robot.RegisterMovementServiceServer(s, &RobotMovementService{})

	go s.Serve(lis)
}

func TestRobotMovementService(t *testing.T) {
	port, addr := GetFreeAddr(t)
	SetupGRPCServer(t, addr)

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	hostname := "localhost"
	url := fmt.Sprintf("%s:%d", hostname, port)

	conn, err := grpc.DialContext(ctx, url, opts...)
	if err != nil {
		t.Errorf("fail to dial: %v", err)
	}
	defer conn.Close()

	client := robot.NewMovementServiceClient(conn)
	res, err := client.Move(ctx, &robot.MoveRequest{})
	if err != nil {
		t.Error(err)
	}

	fmt.Println(res)
}
