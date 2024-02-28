package main

import (
	"context"
	"log"
	"net"

	pb "github.com/khanghoang1210/smart-rent/proto"
	"google.golang.org/grpc"
)

type helloServer struct {
	pb.GreetServiceServer
}

func (*helloServer) SayHello(ctx context.Context, req *pb.HelloReq) (*pb.HelloResp, error) {
	log.Println("Hello word")
	resp := &pb.HelloResp{
		Message: "Hello" + req.GetName(),
	}
	return resp, nil
}
func main() {
	log.Println("Hello word")
	lis, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatalf("Fail to start server %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterGreetServiceServer(grpcServer, &helloServer{})
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Fail to start server %v", err)
	}
}
