package mqttpub

import (
	"os"
	"fmt"
	MQTT "github.com/eclipse/paho.mqtt.golang"
)

type Mqtt struct {
	client   MQTT.Client
	Broker   string
	ClientId string
	Topic    string
}

// Init will initialise the connection to mqtt broker
func (m *Mqtt) Init() {

	opts := MQTT.NewClientOptions().AddBroker(m.Broker)
	opts.SetClientID(m.ClientId)

	c := MQTT.NewClient(opts)

	token := c.Connect()
	if token.Wait() && token.Error() != nil {
		fmt.Printf("Failed to create mqtt connection: %v\n", token.Error())
		os.Exit(-1)
	}

	m.client = c
}

// Publisher to publish the data to mqtt broker
func (m *Mqtt) Publisher(msg []byte) {
	fmt.Println("publishing msg:", string(msg))
	m.client.Publish(m.Topic, 0, false, msg)
}
