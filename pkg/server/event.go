package server

import (
	"context"
	"github.com/gaarutyunov/gueue/pkg/event"
	"github.com/gaarutyunov/gueue/pkg/types"
	"github.com/google/uuid"
	"reflect"
)

func (s *Server) Publish(ctx context.Context, request *event.EventRequest) (*event.EventResponse, error) {
	if t, ok := s.topics[request.GetTopic()]; !ok {
		return nil, types.ErrTopicNotFound.Format(request.GetTopic())
	} else {
		ev := types.NewEvent(request.Topic, request.CorrelationId, request.Message)

		go t.Send(ctx, ev)

		return &event.EventResponse{Id: ev.ID.String()}, nil
	}
}

func (s *Server) Consume(request *event.TopicRequest, server event.Consumer_ConsumeServer) error {
	chs, err := s.addConsumer(request)
	if err != nil {
		return err
	}

	var cases []reflect.SelectCase

	for _, ch := range chs {
		cases = append(cases, reflect.SelectCase{
			Dir:  reflect.SelectRecv,
			Chan: reflect.ValueOf(ch),
		})
	}

	cases = append(cases, reflect.SelectCase{
		Dir:  reflect.SelectRecv,
		Chan: reflect.ValueOf(server.Context().Done()),
	})

	closed := make(map[int]struct{})

	for {
		idx, val, ok := reflect.Select(cases)

		if len(closed) == len(chs) {
			return nil
		}

		if idx <= len(chs)-1 && !ok {
			closed[idx] = struct{}{}
		} else if idx != len(cases)-1 && ok {
			e, ok := val.Interface().(*types.Event)
			if !ok {
				continue
			}
			err = server.Send(&event.Event{
				Id:            e.ID.String(),
				Topic:         e.Topic,
				CorrelationId: e.CorrelationID,
				Message:       e.Message,
			})
			if err != nil {
				return err
			}
		} else if idx == len(cases)-1 {
			for _, topic := range request.GetNames() {
				topic := topic

				s.topics[topic].RemoveConsumer(uuid.MustParse(request.GetConsumerId()))
			}
			return nil
		}
	}
}

func (s *Server) Unbind(ctx context.Context, unbind *event.UnbindRequest) (*event.StatusResponse, error) {
	for _, topic := range s.topics {
		topic.RemoveConsumer(uuid.MustParse(unbind.GetConsumerId()))
	}

	return &event.StatusResponse{
		Status: 200,
	}, nil
}

func (s *Server) addConsumer(request *event.TopicRequest) (ch []<-chan *types.Event, err error) {
	for _, name := range request.GetNames() {
		if t, ok := s.topics[name]; !ok {
			err = types.ErrTopicNotFound.Format(name)
			break
		} else {
			ch = append(ch, t.AddConsumer(uuid.MustParse(request.GetConsumerId()), request.GetBuffer()))
		}
	}

	return
}
