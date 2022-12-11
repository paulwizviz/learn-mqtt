package main

import (
	"flag"
	"fmt"
	"net/url"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/fxamacker/cbor/v2"
	"github.com/paulwizviz/learn-mqtt/internal/agent"
	"github.com/paulwizviz/learn-mqtt/internal/nmodel"
)

func listen(uri *url.URL, topic string, c chan string) {
	client := agent.Connect("sub", uri)
	client.Subscribe(topic, 0, func(_ mqtt.Client, msg mqtt.Message) {
		p := msg.Payload()
		// Unmarshal payload
		var d nmodel.Data
		cbor.Unmarshal(p, &d)
		c <- fmt.Sprintf("* [%s] +----+ %v\n", msg.Topic(), string(d.Content))
	})
}

func main() {

	urlPtr := flag.String("url", "tcp://mosquitto-server:1883/", "location of broker")
	topicPtr := flag.String("topic", "test/hello", "topic name")
	flag.Parse()

	s := *urlPtr
	topic := *topicPtr

	uri, _ := url.Parse(s)

	// Listen to a topic
	c := make(chan string)
	go listen(uri, topic, c)

	for {
		fmt.Println(<-c)
	}

}
