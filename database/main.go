package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	dbadmin "database/dbadmin"
	mqtt    "database/mqttpub"
)

const (
	dbHost = "localhost"
	dbPort = "8086"
	dbName = "empinfo"
	userName = "admin"
	passwrd = "admin"
	broker = "em_broker:1883"
)

var id dbadmin.InfluxDBAdmin
var mp mqtt.Mqtt

// It will start the InfluxDB and create the Database
func startDB() {
        id.Host = dbHost
	id.Port = dbPort
	id.DbName = dbName
	id.UserName = userName
	id.Passwd = passwrd

	id.Init()
	id.CreateDatabase()
}

// It will Initialize the mqtt client connection
func startPublisher() {
	mp.Broker = broker
	mp.Topic = "test"
	mp.ClientId = "pub"
	mp.Init()
        id.Mq = mp

}

// To start the Http server
func startServer() {
	http.HandleFunc("/write", writer)
        http.HandleFunc("/read", reader)
	http.HandleFunc("/load", loader)

	// Starting Http Server at port 7000
	fmt.Println("Starting Http server")
        http.ListenAndServe(":7000", nil)
}

// Http Handler to /write endpoint
func writer(w http.ResponseWriter, r *http.Request) {
	fmt.Println("received POST request")
	if r.Method != http.MethodPost {
                w.Header().Set("Allow", http.MethodPost)
                http.Error(w, http.StatusText(405), 405)
                return
        }

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, http.StatusText(400), 400)
                fmt.Printf("unable to read body: %v\n", err.Error())
                return
        }
        id.InsertData(body)
}

// Http Handler to /read endpoint
func reader(w http.ResponseWriter, r *http.Request) {
        fmt.Println("received POST request")
        if r.Method != http.MethodPost {
                w.Header().Set("Allow", http.MethodPost)
                http.Error(w, http.StatusText(405), 405)
                return
        }

        body, err := ioutil.ReadAll(r.Body)
        if err != nil {
                http.Error(w, http.StatusText(400), 400)
                fmt.Printf("unable to read body: %v\n", err.Error())
                return
        }
        fmt.Println("received query:", string(body))
	value, err := id.QueryDb(string(body))
        if err != nil {
                fmt.Fprintf(w, err.Error())
                fmt.Printf("query influx failed: %v\n", err.Error())
                return
	}

	res, _ := json.Marshal(value)
        fmt.Fprintf(w, string(res))
}

// Http Handler to /load endpoint
func loader(w http.ResponseWriter, r *http.Request) {
        fmt.Println("received POST request")
        if r.Method != http.MethodGet {
                w.Header().Set("Allow", http.MethodGet)
                http.Error(w, http.StatusText(405), 405)
                return
        }
	status := id.LoadCache()
	if status {
	     fmt.Fprintf(w, "Done")
	}
	fmt.Fprintf(w, "failed to load cache")

}

// Main method
func main() {
	done := make(chan bool)
	startDB()
	startPublisher()
	startServer()
	<-done
}
