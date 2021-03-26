package backendoperations

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

const (
	cep = "http://em-cache:4000"
	dep = "http://em-database:7000"
)

// ReadID func is handler to /search api
func ReadID(key string) (string, error) {
	// Get the data from Cache
	status, val := fetchFromCache(key)
	if status {
		fmt.Printf("Value recieved from cache: %v\n", val)
		return val, nil
	}

	// Get the data from InfluxDB
	val1, err := fetchFromDb(key)
	if err != nil {
		return "", err
	}

	if val1 != "" {
		// Save the result in the Cache
		saveInCache(store)
		fmt.Printf("Value recieved from database: %v\n", val1)
		return val1, nil
	}
	return "No record found", nil
}

// SearchResult func will fetch the 10 records from offset
func SearchResult(offset int, empName string) (string, error) {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	loadPayload := make(map[string]string)

	loadPayload["employee_id"] = strconv.Itoa(offset)
	loadPayload["employee_name"] = empName

	loadreq, err := json.Marshal(loadPayload)
	if err != nil {
		return "", err
	}

	_, err = client.Post(dep+"/load", "application/json", bytes.NewBuffer((loadreq)))
	if err != nil {
		return "", err
	}

	time.Sleep(10 * time.Second)
	body := make(map[string]interface{})

	body["key"] = "search"
	body["start"] = strconv.Itoa(offset)
	body["stop"] = strconv.Itoa(offset + 10)

	reqBody, err := json.Marshal(body)
	if err != nil {
		return "", err
	}

	// fetch the unique resuts in 10 search result from cache
	fmt.Println("send body:", string(reqBody))
	res, err := client.Post(cep+"/search", "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return "", err
	}

	bd, _ := ioutil.ReadAll(res.Body)
	return string(bd), nil
}

// Submit func take the data and save it to database and cache
func Submit(data map[string]string) (bool, bool) {
	fmt.Println("Data received in submit method:", data)
	cacheSet := false
	dbSet := false
	//Save the data in Redis cahce
	success := saveInCache(data)
	if success {
		cacheSet = true
	}

	//Save the data in influxDB
	success = saveToDb(data)
	if success {
		dbSet = true
	}

	return cacheSet, dbSet
}
