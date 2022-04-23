package main

import (
	"flag"
	"net/url"
	"time"

	"github.com/paulwizviz/mqtt-mock/internal/agent"
	"github.com/paulwizviz/mqtt-mock/internal/payload"
)

func main() {

	urlPtr := flag.String("url", "tcp://mosquitto-server:1883/", "location of broker")
	topicPtr := flag.String("topic", "test/hello", "topic name")
	flag.Parse()

	s := *urlPtr
	topic := *topicPtr

	uri, _ := url.Parse(s)

	client := agent.Connect("pub", uri)
	timer := time.NewTicker(1 * time.Second)
	for t := range timer.C {
		item := payload.NewData([]byte(t.String()))
		p := payload.MustSerialize(item)
		client.Publish(topic, 0, false, p)
	}
}
