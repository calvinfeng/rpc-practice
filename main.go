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

func runGRPCServer() error {
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
	t3 := todo.AddTask("brush teeth")
	t4 := todo.AddTask("termination")

	todo.AddRelation(t1, t2, taskflow.Success)
	todo.AddRelation(t3, t2, taskflow.Success)
	todo.AddRelation(t2, t4, taskflow.Done)

	taskflow.RunTodo(todo)
}
