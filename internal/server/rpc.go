package server

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/lcrownover/unproc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type unprocServer struct {
	pb.UnimplementedUnprocServer
}

func (s *unprocServer) Retrieve(ctx context.Context, u *pb.User) (*pb.UserReply, error) {
	log.Printf("recieved a user request: %v\n", u.Name)
	count := GetProcessCountforUser(u.Name)
	return &pb.UserReply{Name: u.Name, ProcessCount: count}, nil
}

func NewTCPListener(host string, port string) *net.Listener {
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", host, port))
	// Check for errors
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	return &lis
}

func NewGRPCServer() *grpc.Server {
	// Instantiate the server
	srv := grpc.NewServer()

	pb.RegisterUnprocServer(srv, &unprocServer{})

	reflection.Register(srv)
	return srv
}
