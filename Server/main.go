package main

import (
	"context"
	"flag"
	"log"
	"net"
	"os"
	"time"

	"google.golang.org/grpc"
	pb "google.golang.org/grpc/examples/Go-Exercise/Protobuf"
)

type server struct {
	pb.UnimplementedTimeServer
}

// GetTime implements Time.Server
func (s *server) GetTime(ctx context.Context, in *pb.TimeRequest) (*pb.TimeReply, error) {
	log.Printf("Reqeust received\n")
	currentTime := time.Now()
	timeString := currentTime.Format("2006-01-02 15:04:05")
	return &pb.TimeReply{Message: timeString}, nil
}

func main() {
	flag.Parse()

	port := ":50052"
	if len(os.Args) > 1 {
		port = ":" + os.Args[1]
	}

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterTimeServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
