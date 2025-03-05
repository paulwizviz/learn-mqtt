package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/fxamacker/cbor/v2"
	"github.com/paulwizviz/learn-mqtt/internal/nmodel"
)

func main() {

	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s", "localhost:1883"))
	opts.SetClientID("clientId")
	topic := "test/hello"

	client := mqtt.NewClient(opts)
	// Wait for connection to broker to establish
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatal(token.Error())
	}

	go func() {
		// Create a timer that ticks every second
		timer := time.NewTicker(1 * time.Second)
		for t := range timer.C {
			d := nmodel.Data{
				ID:      nmodel.CreateID(),
				Content: []byte(fmt.Sprintf("Time: %v", t)),
			}
			// Serialize model
			payload, _ := cbor.Marshal(d)
			// Publish
			if token := client.Publish(topic, 0, false, payload); token.Wait() && token.Error() != nil {
				log.Printf("Error publishing to topic %s: %v", topic, token.Error())
				continue
			}
			log.Printf("Published to topic %s: %s", topic, payload)
		}
	}()

	messages := make(chan mqtt.Message)
	go func() {
		if token := client.Subscribe(topic, 0, func(client mqtt.Client, msg mqtt.Message) {
			messages <- msg
		}); token.Wait() && token.Error() != nil {
			log.Fatalf("Error subscribing: %v", token.Error())
		}
	}()

	go func() {
		for message := range messages {
			var m nmodel.Data
			err := cbor.Unmarshal(message.Payload(), &m)
			if err != nil {
				log.Println(err)
				continue
			}
			fmt.Println("--> ", m.ID)
			fmt.Println("--> ", string(m.Content))
		}
	}()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM)
	<-sc
	log.Println("Received termination signal. Shutting down ...")
}
