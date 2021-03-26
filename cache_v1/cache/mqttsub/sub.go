package mqttpub

import (
	redis "cache/redismanager"
	"encoding/json"
	"fmt"
	"os"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

type Mqtt struct {
	//client   MQTT.Client
	Broker   string
	ClientId string
	Topic    string
}

var conn redis.RedisConnect

var store []interface{}

//define a function for the default message handler
var mh MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
	var temp map[string]interface{}
	data := make(map[string]string)
	err := json.Unmarshal(msg.Payload(), &temp)
	if err != nil {
		fmt.Println("error in unmarshal", err)
	}

	_ = conn.Remove("search")
	for k, v := range temp {
		if k == "load" {
			for _, v := range v.([]interface{}) {
				for key, value := range v.(map[string]interface{}) {
					if key == "id" {
						data["key"] = value.(string)
					} else {
						data["value"] = value.(string)
					}
				}
				fmt.Println(data)
				_, err := conn.Store(data["key"], []byte(data["value"]))
				if err != nil {
					fmt.Println("failed to store value:")
				}
				_ = conn.Lpush("search", data["key"])
			}
		}
	}
}

// Init will initailize the mqtt client and create a connection with mqtt broker
func (m *Mqtt) Init() {

	opts := MQTT.NewClientOptions().AddBroker(m.Broker)
	opts.SetClientID(m.ClientId)
	opts.SetDefaultPublishHandler(mh)

	c := MQTT.NewClient(opts)

	token := c.Connect()
	if token.Wait() && token.Error() != nil {
		fmt.Printf("Failed to create mqtt connection: %v\n", token.Error())
		os.Exit(-1)
	}

	if token := c.Subscribe(m.Topic, 0, nil); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}
}
