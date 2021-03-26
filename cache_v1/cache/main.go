package main

import (
	mqtt "cache/mqttsub"
	redis "cache/redismanager"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	topic  = "test"
	broker = "em-broker:1883"
)

var mp mqtt.Mqtt
var conn redis.RedisConnect

// Main function
func main() {
	redis.Init()
	redis.NewRedisConnect()
	mp.Broker = broker
	mp.Topic = topic
	mp.ClientId = "sub"
	mp.Init()

	http.HandleFunc("/get", fetch)
	http.HandleFunc("/set", save)
	http.HandleFunc("/search", search)
	// Start the Http Server at port 4000
	http.ListenAndServe(":4000", nil)
}

// Http handler to /get endpoint
func fetch(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Received a get request\n")
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		http.Error(w, http.StatusText(405), 405)
		return
	}

	key := r.URL.Query().Get("key")
	fmt.Printf("Key in get request: %v\n", key)
	if key == "" {
		http.Error(w, http.StatusText(400), 400)
		return
	}
	// Read the key from Redis
	val, err := conn.Read(key)
	if err != nil {
		fmt.Fprintf(w, "error in retrieving Key:", err)
		return
	}
	fmt.Fprintf(w, val)
}

// Http handler to /set endpoint
func save(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Received a post request\n")
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, http.StatusText(405), 405)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	fmt.Printf("Post request body: %s\n", string(body))
	var data1 map[string]string
	err = json.Unmarshal(body, &data1)
	if err != nil {
		fmt.Printf("Error while json Unmarshal: %v\n", err.Error())
		fmt.Fprintf(w, "error in storing data")
		return
	}
	fmt.Printf("Data parsed as: %v\n", data1)
	// Save the data in Cache
	val, err := conn.Store(data1["key"], []byte(data1["value"]))
	if err != nil {
		fmt.Fprintf(w, "error in storing data")
		return
	}
	fmt.Fprintf(w, val)
}

// Http handler to /search endpoint
func search(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Received a post request\n")
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, http.StatusText(405), 405)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	fmt.Printf("Post request body: %s\n", string(body))
	var data1 map[string]interface{}
	err = json.Unmarshal(body, &data1)
	if err != nil {
		fmt.Printf("Error while json Unmarshal: %v\n", err.Error())
		fmt.Fprintf(w, "error in storing data")
		return
	}
	fmt.Printf("Data parsed as: %v\n", data1)

	// Read the 10 data from list
	val, err := conn.Lrange(data1["key"].(string), int64(data1["start"].(float64)), int64(data1["stop"].(float64)))
	if err != nil {
		fmt.Fprintf(w, "error in storing data")
		return
	}
	rd := make(map[string]string)
	for _, v := range val {
		value, _ := conn.Read(v)
		rd[v] = value
	}
	fd, _ := json.Marshal(rd)
	fmt.Fprintf(w, string(fd))
}
