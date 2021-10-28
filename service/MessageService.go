package service

import (
	"context"

	pb "github.com/Echo-P-Chang/Lagom1-Dependency"
)

type MessageService struct {
	pb.UnimplementedGreeterServer
}

func (m MessageService) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {

	return &pb.HelloReply{Message: "Server say hello to " + in.GetName()}, nil
	//SayHello(context.Context, *HelloRequest) (*HelloReply, error)
}

// func (m MessageService) UnimplementedGreeterServer() {

// 	s := ""
// 	log.Println(s)
// 	//SayHello(context.Context, *HelloRequest) (*HelloReply, error)
// }
// func (m MessageService) mustEmbedUnimplementedGreeterServer() {

// 	s := ""
// 	log.Println(s)
// 	//SayHello(context.Context, *HelloRequest) (*HelloReply, error)
// }
