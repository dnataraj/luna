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
	NsqAddr string
	Config  *nsq.Config
	Clients map[string]*nsq.Producer
}

func init() {
	log.SetPrefix("logger >> ")
}

// AddLogger initializes the logging service and returns a client or tenant id that identifies
// this particular client instance
func (s *LoggerServer) AddLogger(ctx context.Context, cfg *pb.Config) (*pb.ClientId, error) {
	id := uuid.New().String()
	c := nsq.NewConfig()
	prod, err := nsq.NewProducer(s.NsqAddr, c)
	if err != nil {
		log.Printf("error initializing NSQ producer for client: %s: %v", id, err)
		return nil, fmt.Errorf("unable to initialize client: %w", err)
	}
	s.Clients[id] = prod
	return &pb.ClientId{Id: id}, nil
}

// LogLine publishes a log record to a Log queue for further processing
func (s *LoggerServer) LogLine(ctx context.Context, ln *pb.Log) (*empty.Empty, error) {
	id := ln.GetClientId().GetId()
	if len(id) == 0 {
		return &empty.Empty{}, fmt.Errorf("invalid client id provided")
	}
	// format the log message
	// severity  timestamp  id  message
	m := fmt.Sprintf("%s %s", ln.GetSeverity().String(), ln.GetLogMessage())
	log.Printf("client [%s]: logging message: %s\n", id, m)
	p := s.Clients[id]
	err := p.Publish(id, []byte(m))
	if err != nil {
		return &empty.Empty{}, fmt.Errorf("error writing log: %w", err)
	}
	return &empty.Empty{}, nil
}
