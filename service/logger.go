package service

import (
	"context"
	"fmt"
	"log"

	pb "github.com/dnataraj/luna/logger"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/google/uuid"
	"github.com/nsqio/go-nsq"
)

type LoggerServer struct {
	Config  *nsq.Config
	Clients map[string]*nsq.Producer
}

// AddLogger initializes the logging service and returns a client or tenant id that identifies
// this particular client instance
func (s *LoggerServer) AddLogger(ctx context.Context, cfg *pb.Config) (*pb.ClientId, error) {
	id := uuid.New().String()
	c := nsq.NewConfig()
	prod, err := nsq.NewProducer("127.0.0.1:4150", c)
	if err != nil {
		log.Printf("error initializing NSQ producer for client: %s: %v", id, err)
		return nil, fmt.Errorf("unable to initialize client: %w", err)
	}
	s.Clients[id] = prod
	return &pb.ClientId{Id: id}, nil
}

// LogLine publishes a log record to a Log queue for further processing
func (s *LoggerServer) LogLine(ctx context.Context, log *pb.Log) (*empty.Empty, error) {
	return nil, nil
}
