package dbops

import "fmt"

func initKafkaConfig() *kafkaConfig {
	kafka := kafkaConfig{
		BrokerList: []string{"em-broker:9092"},
		Topic:      "load-cache",
	}

	return &kafka
}

func initDBCOnfig() *Dbconnection {
	dbConfig := &Dbconnection{
		dbuser: "root",
		passwd: "",
		dbhost: "127.0.0.1",
		dbport: "3306",
		dbname: "employeerecord",
	}

	return dbConfig
}

func Init() (*Dbconnection, *KafkaProducer) {
	kp := &KafkaProducer{}
	dbc := &Dbconnection{}

	kp.conf = initKafkaConfig()
	kp.ConnectBrokers()
	kp.RegProducer()

	dbc = initDBCOnfig()
	fmt.Printf("Initialized Config is: %v\n", dbc)
	err := dbc.InitDB()
	if err != nil {
		fmt.Printf("Failed to open connection\n")
	}

	return dbc, kp
}
