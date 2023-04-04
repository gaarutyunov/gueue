package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/gaarutyunov/gueue/pkg/client"
	log "github.com/sirupsen/logrus"
	"strings"
)

var (
	topic  = flag.String("topic", "test.topic.1", "Topics to consume (split by ',')")
	buffer = flag.Uint("buffer", 100, "Buffer size")
)

func main() {
	flag.Parse()
	consumer := client.NewConsumer("localhost", 8001, true)

	ctx := context.Background()

	err := consumer.Dial(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	defer func(consumer *client.Consumer) {
		_ = consumer.Close(ctx)
	}(consumer)

	topics := strings.Split(*topic, ",")

	ch, errCh, err := consumer.Consume(ctx, uint32(*buffer), topics...)
	if err != nil {
		log.Fatalln(err)
	}

	for {
		select {
		case val, ok := <-ch:
			if !ok {
				fmt.Printf("Closed on %s\n", consumer.ID)
			} else {
				fmt.Printf("Consumed value on %s from %s: %s\n", consumer.ID, val.Topic, val.Message)
			}
		case err := <-errCh:
			log.Fatalf("Error on %s: %s\n", consumer.ID, err)
		}
	}
}
