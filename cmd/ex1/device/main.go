package main

import (
	"flag"
	"fmt"
	"net/url"
	"time"

	"github.com/fxamacker/cbor/v2"
	"github.com/paulwizviz/learn-mqtt/internal/agent"
	"github.com/paulwizviz/learn-mqtt/internal/nmodel"
)

func main() {

	urlPtr := flag.String("url", "tcp://mosquitto-server:1883/", "location of broker")
	topicPtr := flag.String("topic", "test/hello", "topic name")
	flag.Parse()

	s := *urlPtr
	topic := *topicPtr

	uri, _ := url.Parse(s)

	// Get a publishing client
	client := agent.Connect("pub", uri)
	// Create a timer that ticks every second
	timer := time.NewTicker(1 * time.Second)
	for t := range timer.C {
		d := nmodel.Data{
			ID:      nmodel.CreateID(),
			Content: []byte(fmt.Sprintf("Time: %v", t)),
		}
		// Serialize model
		cborD, _ := cbor.Marshal(d)
		// Publish
		client.Publish(topic, 0, false, cborD)
	}
}
