package client

import (
	"context"
	"errors"
	"fmt"
	pb "github.com/gaarutyunov/gueue/pkg/event"
	"github.com/gaarutyunov/gueue/pkg/types"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
)

type (
	Consumer struct {
		ID       uuid.UUID
		addr     string
		insecure bool
		client   pb.ConsumerClient
		conn     *grpc.ClientConn
	}
)

func NewConsumer(host string, port uint, insecure bool) *Consumer {
	addr := fmt.Sprintf("%s:%d", host, port)

	return &Consumer{ID: uuid.New(), addr: addr, insecure: insecure}
}

func (c *Consumer) Dial(ctx context.Context, opts ...grpc.DialOption) error {
	if c.insecure {
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}

	conn, err := grpc.DialContext(ctx, c.addr, opts...)
	if err != nil {
		return nil
	}

	c.client = pb.NewConsumerClient(conn)
	c.conn = conn

	return nil
}

func (c *Consumer) Close(ctx context.Context) error {
	_, err := c.client.Unbind(ctx, &pb.UnbindRequest{ConsumerId: c.ID.String()})
	if err != nil {
		return err
	}

	return c.conn.Close()
}

func (c *Consumer) Consume(ctx context.Context, buffer uint32, names ...string) (<-chan *types.Event, <-chan error, error) {
	stream, err := c.client.Consume(ctx, &pb.TopicRequest{
		Names:      names,
		ConsumerId: c.ID.String(),
		Buffer:     buffer,
	})
	if err != nil {
		return nil, nil, err
	}

	out := make(chan *types.Event, buffer)
	errCh := make(chan error)

	go func() {
		for {
			ev, err := stream.Recv()
			if err != nil {
				if errors.Is(err, io.EOF) {
					errCh <- stream.CloseSend()
				}

				errCh <- err
			}

			select {
			case out <- types.NewEvent(ev.GetTopic(), ev.GetCorrelationId(), ev.GetMessage()):
			case <-ctx.Done():
				close(out)
				return
			}
		}
	}()

	return out, errCh, nil
}
