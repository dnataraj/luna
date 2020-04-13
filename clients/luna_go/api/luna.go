package api

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/dnataraj/luna/logger"
	"google.golang.org/grpc"
)

type Client struct {
	serverAddr string
	conn       *grpc.ClientConn
	id         *pb.ClientId
	logger     pb.LoggerClient
}

// New initializes and returns a reference to a logger client. addr is the address (host:port)
// of the logger service
func New(addr string) (*Client, error) {
	log.SetPrefix("client >> ")
	c := &Client{serverAddr: addr}
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	log.Println("connecting to log service...")
	//TODO: Dial or DialContext?
	conn, err := grpc.Dial(addr, opts...)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to logger service: %w", err)
	}
	c.conn = conn
	c.logger = pb.NewLoggerClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	log.Println("adding client logger...")
	id, err := c.logger.AddLogger(ctx, &pb.Config{})
	if err != nil {
		return nil, fmt.Errorf("faled to initialize logger service: %w", err)
	}
	c.id = id

	return c, nil
}

// Write writes out the provided bytes to the luna logger service with INFO severity level
// If the first bytes indicate severity, then that is used.
func (c *Client) Write(p []byte) (int, error) {
	log.Printf("generating log for %d  bytes\n", len(p))
	l := &pb.Log{
		ClientId: c.id,
		//Treat standard log.PrintXX() as INFO severity
		Severity:   pb.Severity_INFO,
		LogMessage: string(p),
	}
	//TODO: Figure out how to build a context in this case, if it is a case that we support
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if c.logger == nil {
		log.Fatal("luna logger not initialized")
	}
	//TODO: Use WaitForReady semantics in call options?
	_, err := c.logger.LogLine(ctx, l)
	if err != nil {
		return 0, fmt.Errorf("unable to log message: %w", err)
	}
	return len(p), nil

}

func (c *Client) Close() error {
	return c.conn.Close()
}
