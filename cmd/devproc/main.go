package main

import (
	"flag"
	"fmt"
	"net/url"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/paulwizviz/learn-mqtt/internal/agent"
	"github.com/paulwizviz/learn-mqtt/internal/payload"
)

func listen(uri *url.URL, topic string, c chan string) {
	client := agent.Connect("sub", uri)
	client.Subscribe(topic, 0, func(_ mqtt.Client, msg mqtt.Message) {
		p := msg.Payload()
		i := payload.MustDeserialize(p)
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
