package main

import (
	"fmt"
	"net"
	"rpc-practice/protogens/go/robot"
	"rpc-practice/server"
	"rpc-practice/taskflow"
	"time"

	"google.golang.org/grpc"
)

func RunGRPCServer() error {
	lis, err := net.Listen("tcp", ":3000")
	if err != nil {
		return err
	}

	gRPCServer := grpc.NewServer()
	robot.RegisterNavigationServer(gRPCServer, server.NewRobotNavigationServer())

	go gRPCServer.Serve(lis)

	fmt.Println("Let's do gRPC")
	time.Sleep(10 * time.Second)

	gRPCServer.Stop()

	return nil
}

func main() {
	todo := taskflow.NewTodo("things to do tonight")
	t1 := todo.AddTask("shower")
	t2 := todo.AddTask("dry hair")
	todo.AddRelation(t1, t2)

	todo.Traverse()
}
