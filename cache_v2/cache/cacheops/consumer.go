package cacheops

import (
	"encoding/json"
	"fmt"

	"github.com/Shopify/sarama"
)

type kafkaConfig struct {
	BrokerList []string
	Topic      string
}

type KafkaConsumer struct {
	Consumer sarama.PartitionConsumer
	brokers  []*sarama.Broker
	conf     *kafkaConfig
}

type employee struct {
	EmployeeID   string
	EmployeeName string
}

func (kc *KafkaConsumer) ConnectBrokers() {
	for _, b := range kc.conf.BrokerList {
		brokeraddr := kc.connectBroker(b)
		kc.brokers = append(kc.brokers, brokeraddr)
	}
}

// Connect Broker
func (kc *KafkaConsumer) connectBroker(b string) *sarama.Broker {
	config := sarama.NewConfig()
	config.Version = sarama.V1_0_0_0
	broker := sarama.NewBroker(b)

	for {
		err := broker.Open(config)
		if err != nil {
			_ = fmt.Errorf("Not able to connect Broker, error is: %v", err)
		}
		fmt.Printf("Checking broker connectivity\n")
		status, err := broker.Connected()
		fmt.Printf("Broker connection status: %v\n", status)
		if status {
			break
		}
		fmt.Printf("Please Start the Kafka Broker and try again \n")

		if err != nil {
			_ = fmt.Errorf("\n Broker not connected: %v", err)
		}
	}

	return broker
}

// Register Producer
func (kc *KafkaConsumer) RegConsumer() {
	// Prepare the consumer config
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config.Version = sarama.V1_0_0_0

	// creating Consumer
	consumer, err := sarama.NewConsumer(kc.conf.BrokerList, config)
	if err != nil {
		_ = fmt.Errorf("\n Unable to create consumer: %v", err)
	}

	fmt.Printf("Consumer Created %v\n", consumer)

	kc.Consumer, err = consumer.ConsumePartition(kc.conf.Topic, 0, sarama.OffsetOldest)
	if err != nil {
		_ = fmt.Errorf("\n Error in creating partition Consumer %v", err)
	}
}

func (kc *KafkaConsumer) ReadMessage() {
	rc, _ := InitCacheAndConsumer()
	recMsg := kc.Consumer.Messages()
	fmt.Printf("In Kafka Read message\n")
	for rmsg := range recMsg {
		fmt.Printf("Message received is: %v\n", string(rmsg.Value))
		var tempStore []employee

		err := json.Unmarshal(rmsg.Value, &tempStore)
		if err != nil {
			fmt.Printf("Error in Unmarshal in ReadMessage func: %v\n", err)
		}

		fmt.Printf("Message received in tempstore: %v\n", tempStore)
		for _, v := range tempStore {

			// redis client
			_, err = rc.Store(v.EmployeeID, []byte(v.EmployeeName))
			if err != nil {
				fmt.Printf("Error in writing data to cache: %v\n", err)
			}

			_ = rc.Lpush("search", v.EmployeeName)
		}

	}

}

func (kc *KafkaConsumer) Close() {
	kc.Consumer.Close()

	for _, b := range kc.brokers {
		b.Close()
	}
}
