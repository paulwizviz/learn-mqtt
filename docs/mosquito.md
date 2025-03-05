# Eclipse Mosquitto

This section discusses techniques to provision and developing applications against Mosquito.  


* [Go module](https://github.com/eclipse/paho.mqtt.golang.git)
* [mosquitto.conf](https://mosquitto.org/man/mosquitto-conf-5.html)


## Provision Examples

The following include examples of mosquito setup.

### Solo

This example showcase a setup with a solo broker with minimu (no password or any security setup).

Here is the:

* [Docker configuration](../deployment/solo/docker-compose.yaml)
* [Mosquito configuration](../deployment/solo/mosquitto.conf)
* [Operational script](../scripts/broker.sh) - To run the solo network, use the script and run the following commands `./scripts/broker.sh solo [start or stop]`

## Application Development Working Examples

### Example 1

This example demonstrates a simple pub-sub application.
Here is the [source](../examples/ex1/main.go)



## References

* [Eclipse Mosquitto](https://mosquitto.org/).