package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	pb "github.com/dnataraj/luna/logger"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var (
	serverAddr = flag.String("addr", "localhost:5050", "The logger server address in the form of host:port")
)

func main() {
	flag.Parse()
	log.SetPrefix("client >> ")
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithBlock())

	log.Println("connecting to log service...")
	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		log.Fatal("unable to connect to logger server: ", err)
	}
	defer conn.Close()

	client := pb.NewLoggerClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	log.Println("adding client logger...")
	id, err := client.AddLogger(ctx, &pb.Config{})
	if err != nil {
		fmt.Println("failed to initialize logger client: ", err)
	}

	log.Println("initialized logger client: ", id.GetId())
}
