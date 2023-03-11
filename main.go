package main

import (
	"context"
	"log"
	"net"
	"net/http"

	todo_listv1 "github.com/aalindg/boilerplate/proto/todo_list/v1"
	user "github.com/aalindg/boilerplate/proto/user/v1"
	"github.com/aalindg/boilerplate/server"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8080"
	addr = "0.0.0.0:10000"
)

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalln("Failed to listen at port: ", port)
	}

	gs := grpc.NewServer(
		grpc.Creds(insecure.NewCredentials()),
	)

	user.RegisterUserServiceServer(gs, server.New())
	todo_listv1.RegisterTodoListServiceServer(gs, server.New())

	log.Println("Serving gRPC on https://", addr)
	go func() {
		log.Fatal(gs.Serve(lis))
	}()

	conn, err := grpc.DialContext(
		context.Background(),
		"0.0.0.0:8080",
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalln("Failed to dial serer: ", err)
	}

	gwmux := runtime.NewServeMux()

	err = user.RegisterUserServiceHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway: ", err)
	}

	gwServer := &http.Server{
		Addr:    ":8090",
		Handler: gwmux,
	}

	log.Println("Serving gRPC-Gateway on http://0.0.0.0:8090")
	log.Fatalln(gwServer.ListenAndServe())

	// err = gateway.Run("dns:///" + addr)
	// log.Fatalln(err)
}
