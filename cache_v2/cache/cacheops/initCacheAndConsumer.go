package cacheops

func InitCacheAndConsumer() (*RedisConnect, *KafkaConsumer) {
	Init()
	NewRedisConnect()
	var conn *RedisConnect

	kafka := kafkaConfig{
		BrokerList: []string{"em-broker:9092"},
		Topic:      "load-cache",
	}

	kc := &KafkaConsumer{}
	kc.conf = &kafka
	kc.ConnectBrokers()
	kc.RegConsumer()

	return conn, kc
}
