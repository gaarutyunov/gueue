package server

import (
	"context"
	"fmt"
	"github.com/gaarutyunov/gueue/pkg/encoding"
	"github.com/gaarutyunov/gueue/pkg/event"
	"github.com/gaarutyunov/gueue/pkg/types"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type (
	Server struct {
		event.UnimplementedConsumerServer
		event.UnimplementedProducerServer
		topics map[string]*types.Topic

		srv *grpc.Server
	}
)

func NewServer() *Server {
	return &Server{
		topics: map[string]*types.Topic{},
	}
}

func (s *Server) Start(ctx context.Context) error {
	ctx, cancel := context.WithCancelCause(ctx)
	stop := make(chan os.Signal, 1)

	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go func() {
		err := s.doStart(ctx)
		if err != nil {
			cancel(err)
		}
	}()

	for {
		select {
		case snl := <-stop:
			switch snl {
			case os.Interrupt:
				cancel(types.ErrShutdown.Format(snl))
			case syscall.SIGTERM:
				cancel(types.ErrKill.Format(snl))
			}
		case <-ctx.Done():
			return s.Stop(ctx)
		}
	}
}

func (s *Server) Stop(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second*10)

	defer cancel()

	switch cause := context.Cause(ctx).(type) {
	case *types.Error:
		switch cause.Code {
		case types.ErrShutdown:
			s.doStop(ctx)
			return nil
		default:
			return cause
		}
	case error:
		return types.ErrCritical.Format(cause)
	default:
		s.doStop(ctx)
		return nil
	}
}

func (s *Server) doStop(ctx context.Context) {
	log.Infof("Gracefully stopping")
	s.srv.GracefulStop()
	for _, topic := range s.topics {
		go func(topic *types.Topic) {
			topic.Flush()
		}(topic)
	}
	log.Infof("Gracefully stopped")
}

func (s *Server) doStart(ctx context.Context) error {
	var topics []*types.Topic

	err := encoding.UnmarshalKey("topics", &topics)
	if err != nil {
		return err
	}

	for _, topic := range topics {
		topic := topic

		s.topics[topic.Name] = topic
		go topic.Listen(ctx)
	}

	host := viper.GetString("server.host")
	port := viper.GetInt("server.port")

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		return err
	}

	s.srv = grpc.NewServer()
	event.RegisterConsumerServer(s.srv, s)
	event.RegisterProducerServer(s.srv, s)

	log.Infof("Server listening tcp on %s:%d", host, port)

	return s.srv.Serve(lis)
}
