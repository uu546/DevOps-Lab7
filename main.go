package main

import (
	"fmt"
	"net"
	"net/http"

	"google.golang.org/grpc"

	bookgrpc "example.com/book-service/grpc"
	bookhttp "example.com/book-service/http"
	"example.com/book-service/implementation"
	"example.com/book-service/pb"
	"example.com/book-service/repository"
	"example.com/book-service/transport"
)

func main() {

	repo := repository.NewMemoryRepository()
	svc := implementation.NewService(repo)
	endpoints := transport.MakeEndpoints(svc)

	// HTTP на :8080
	httpHandler := bookhttp.NewService(endpoints)
	go func() {
		fmt.Println("HTTP server on :8080")
		if err := http.ListenAndServe(":8080", httpHandler); err != nil {
			fmt.Println("HTTP error:", err)
		}
	}()

	// gRPC на :9090
	grpcServer := bookgrpc.NewGRPCServer(endpoints)

	listener, err := net.Listen("tcp", ":9090")
	if err != nil {
		fmt.Println("gRPC listen error:", err)
		return
	}

	s := grpc.NewServer()
	pb.RegisterBookServiceServer(s, grpcServer)

	fmt.Println("gRPC server on :9090")
	if err := s.Serve(listener); err != nil {
		fmt.Println("gRPC error:", err)
	}
}