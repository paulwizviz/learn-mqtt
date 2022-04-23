package main

import (
	"flag"
	"fmt"
	"net/url"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/paulwizviz/mqtt-mock/internal/agent"
	"github.com/paulwizviz/mqtt-mock/internal/payload"
)

func listen(uri *url.URL, topic string, c chan string) {
	client := agent.Connect("sub", uri)
	client.Subscribe(topic, 0, func(client mqtt.Client, msg mqtt.Message) {
		p := msg.Payload()
		i := payload.MistDeserialize(p)
		c <- fmt.Sprintf("* [%s] +----+ %v\n", msg.Topic(), i)
	})
}

func main() {

	urlPtr := flag.String("url", "tcp://mosquitto-server:1883/", "location of broker")
	topicPtr := flag.String("topic", "test/hello", "topic name")
	flag.Parse()

	s := *urlPtr
	topic := *topicPtr

	uri, _ := url.Parse(s)

	c := make(chan string)
	go listen(uri, topic, c)

	for {
		fmt.Println(<-c)
	}

}
