package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "google.golang.org/grpc/examples/Go-Exercise/Protobuf"
	"log"
	"net"
	"os"
	"strings"
)

type server struct {
	pb.UnimplementedTimeServer
}

var ip []string
var port []string
var count int

// GetTime implements Time.Server
func (s *server) GetTime(ctx context.Context, in *pb.TimeRequest) (*pb.TimeReply, error) {

	//Connessione al server:
	serverAddress := fmt.Sprintf("%s:%s", ip[count], port[count])
	conn, err := grpc.Dial(serverAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	//Chiudo connessione
	defer conn.Close()

	// Creo client gRPC associato alla connessione, in grado di invocare operazioni definite nel protobuf del servizio Time.
	c := pb.NewTimeClient(conn)

	//Richiedo servizio:
	r, err := c.GetTime(ctx, &pb.TimeRequest{})
	if err != nil {
		log.Fatalf("Error on request: %v", err)
	}
	log.Printf("%d", len(ip))
	count = (count + 1) % len(ip)

	log.Printf("Send Time to client")
	return &pb.TimeReply{Message: r.GetMessage()}, nil
}

func main() {
	flag.Parse()
	portLB := ":50051"
	lis, err := net.Listen("tcp", portLB)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterTimeServer(s, &server{})
	log.Printf("Load Balancer listening at %v", lis.Addr())

	//Andiamo a recuperare i vari Ip e # di porta dei nostri server
	//Scan del file di configurazione:
	f, err := os.Open("configuration.txt")
	if err != nil {
		log.Fatalf("Error", err)
	}

	//Chiusura file
	defer f.Close()

	//Lettura File
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)

	//Inseriamo Ip e # di porta
	for scanner.Scan() {
		word := scanner.Text()
		if strings.Contains(word, "Ip:") {
			scanner.Scan()
			ip = append(ip, scanner.Text())
			//log.Printf("Ip: %s", ip)
		}
		if strings.Contains(word, "Port") {
			scanner.Scan()
			port = append(port, scanner.Text())
			//log.Printf("Port: %s", port)

		}

	}

	//Avviamo servizio:
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
