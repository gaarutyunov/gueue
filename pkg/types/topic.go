package types

import (
	"context"
	"github.com/google/uuid"
	"github.com/spf13/cast"
	"sync"
)

type (
	Topic struct {
		Name   string `yaml:"name"`
		Buffer uint32 `yaml:"buffer"`

		producer  chan *Event
		consumers map[uuid.UUID]chan *Event

		mu sync.RWMutex
	}

	Topics []*Topic
)

func (t *Topic) UnmarshalMap(node map[string]interface{}) error {
	t.Name = node["name"].(string)
	t.Buffer = cast.ToUint32(node["buffer"])

	t.producer = make(chan *Event, t.Buffer)
	t.consumers = make(map[uuid.UUID]chan *Event)

	return nil
}

func (t *Topic) AddConsumer(id uuid.UUID, buffer uint32) <-chan *Event {
	if ch, ok := t.HasConsumer(id); ok {
		return ch
	}

	t.mu.Lock()
	defer t.mu.Unlock()

	ch := make(chan *Event, buffer)

	t.consumers[id] = ch

	return ch
}

func (t *Topic) Flush() {
	t.mu.Lock()
	defer t.mu.Unlock()

	close(t.producer)

	for _, ch := range t.consumers {
		close(ch)
	}
}

func (t *Topic) RemoveConsumer(id uuid.UUID) {
	ch, ok := t.hasConsumer(id)
	if !ok {
		return
	}

	t.mu.Lock()
	defer t.mu.Unlock()

	close(ch)

	delete(t.consumers, id)
}

func (t *Topic) HasConsumer(id uuid.UUID) (<-chan *Event, bool) {
	return t.hasConsumer(id)
}

func (t *Topic) hasConsumer(id uuid.UUID) (chan *Event, bool) {
	t.mu.RLock()
	defer t.mu.RUnlock()

	ch, ok := t.consumers[id]

	return ch, ok
}

func (t *Topic) Listen(ctx context.Context) {
	for event := range t.producer {
		event := event

		for _, ch := range t.consumers {
			go func(ch chan<- *Event, event *Event) {
				select {
				case ch <- event:
				case <-ctx.Done():
					return
				}
			}(ch, event)
		}
	}
}

func (t *Topic) Send(ctx context.Context, event *Event) {
	select {
	case t.producer <- event:
	case <-ctx.Done():
		return
	}
}
