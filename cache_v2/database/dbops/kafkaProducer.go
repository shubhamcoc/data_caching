package dbops

import (
	"errors"
	"fmt"

	"github.com/Shopify/sarama"
)

type kafkaConfig struct {
	BrokerList []string
	Topic      string
}

type KafkaProducer struct {
	Producer sarama.SyncProducer
	brokers  []*sarama.Broker
	conf     *kafkaConfig
}

func (kp *KafkaProducer) ConnectBrokers() {
	//brokerList := []string{"127.0.0.1:9092"}
	// invoke connect broker
	for _, b := range kp.conf.BrokerList {
		brokeraddr := kp.connectBroker(b)
		kp.brokers = append(kp.brokers, brokeraddr)
	}
}

// Connect Broker
func (kp *KafkaProducer) connectBroker(b string) *sarama.Broker {
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
		fmt.Printf("Please start the broker and try again\n")

		if err != nil {
			_ = fmt.Errorf("\n Broker not connected: %v", err)
		}
	}

	return broker
}

// Register Producer
func (kp *KafkaProducer) RegProducer() {
	// prepare the producer config
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	config.Version = sarama.V1_0_0_0

	// creating producer
	producer, err := sarama.NewSyncProducer(kp.conf.BrokerList, config)
	if err != nil {
		_ = fmt.Errorf("\n Unable to create producer: %v", err)
	}

	fmt.Printf("\n Producer created")
	kp.Producer = producer
}

func (kp *KafkaProducer) SendMessage(msg string) error {
	fmt.Printf("Received message in SendMessage: %v\n", msg)
	finalMsg := &sarama.ProducerMessage{
		Topic: kp.conf.Topic,
		Value: sarama.ByteEncoder(msg),
	}

	if kp.Producer != nil {
		kp.Producer.SendMessage(finalMsg)
		return nil
	}

	return errors.New("Producer not registered")
}

func (kp *KafkaProducer) Close() {
	kp.Producer.Close()

	for _, b := range kp.brokers {
		b.Close()
	}
}
