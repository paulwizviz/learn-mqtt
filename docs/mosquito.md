# Eclipse Mosquitto

Eclipse Mosquitto is an open source (EPL/EDL licensed) message broker that implements the MQTT protocol versions 5.0, 3.1.1 and 3.1. Mosquitto is lightweight and is suitable for use on all devices from low power single board computers to full servers (see [offical Ddcs](https://mosquitto.org/).

## Mosquitto client

* [Go module](https://github.com/eclipse/paho.mqtt.golang.git)

## Mosquitto broker

* [mosquitto.conf](https://mosquitto.org/man/mosquitto-conf-5.html)

## Working examples

<u>Example 1</u>

This example demonstrates an IoT device sending data payload to a backend service via Mosquitto (MQTT broker). The demonstrator is simulated using docker compose. The data payload is serialized using CBOR. The working example set up is here:

* [docker compose network](../deployment/ex1/docker-compose.yml)
* [operational script](../scripts/ex1.sh)