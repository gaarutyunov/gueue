package main

import (
	"context"
	"flag"
	"github.com/gaarutyunov/gueue/pkg/client"
	"github.com/gaarutyunov/gueue/pkg/types"
	log "github.com/sirupsen/logrus"
)

var (
	topic = flag.String("topic", "test.topic.1", "Topic to consume")
	num   = flag.Int("num", 1000, "Num events")
)

func main() {
	flag.Parse()
	producer := client.NewProducer("localhost", 8001, true)

	ctx := context.Background()

	err := producer.Dial(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	defer func(producer *client.Producer) {
		_ = producer.Close()
	}(producer)

	for i := 0; i < *num; i++ {
		request := map[string]interface{}{
			"idx": i,
		}

		err := producer.Publish(ctx, types.FromJSON(*topic, "", request))
		if err != nil {
			log.Fatalln(err)
		}
	}
}
