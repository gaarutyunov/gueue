package client

import (
	"context"
	"fmt"
	pb "github.com/gaarutyunov/gueue/pkg/event"
	"github.com/gaarutyunov/gueue/pkg/types"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type (
	Producer struct {
		addr     string
		insecure bool
		conn     *grpc.ClientConn
		client   pb.ProducerClient
	}
)

func NewProducer(host string, port uint, insecure bool) *Producer {
	addr := fmt.Sprintf("%s:%d", host, port)

	return &Producer{addr: addr, insecure: insecure}
}

func (c *Producer) Dial(ctx context.Context, opts ...grpc.DialOption) error {
	if c.insecure {
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}

	conn, err := grpc.DialContext(ctx, c.addr, opts...)
	if err != nil {
		return nil
	}

	c.client = pb.NewProducerClient(conn)
	c.conn = conn

	return nil
}

func (c *Producer) Close() error {
	return c.conn.Close()
}

func (c *Producer) Publish(ctx context.Context, event *types.Event) error {
	res, err := c.client.Publish(ctx, &pb.EventRequest{
		Topic:         event.Topic,
		CorrelationId: event.CorrelationID,
		Message:       event.Message,
	})
	if err != nil {
		return err
	}

	event.ID = uuid.MustParse(res.Id)

	return nil
}
