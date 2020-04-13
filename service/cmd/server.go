package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/dnataraj/luna/logger"
	"github.com/dnataraj/luna/service"
	"github.com/nsqio/go-nsq"
	"google.golang.org/grpc"
)

var (
	port     = flag.Int("port", 5050, "The server port")
	nsq_addr = flag.String("nsq_addr", "localhost:4150", "The address for NSQ")
)

func main() {
	flag.Parse()
	log.SetPrefix("server >> ")
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption

	server := grpc.NewServer(opts...)
	ls := &service.LoggerServer{
		NsqAddr: *nsq_addr,
		Config:  nsq.NewConfig(),
		Clients: make(map[string]*nsq.Producer, 0),
	}
	log.Println("registering logger service...")
	pb.RegisterLoggerServer(server, ls)
	log.Printf("starting server on: %d\n", *port)
	server.Serve(lis)
}
