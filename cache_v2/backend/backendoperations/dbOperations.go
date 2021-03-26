package backendoperations

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var store map[string]string

// fetchFromDb common func to send the query to MariaDB
func fetchFromDb(key string) (string, error) {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	url := dep + "/read?employee_id=" + key
	res, err := client.Get(url)
	if err != nil {
		fmt.Printf("Failed to make request: %s", err)
		return "", err
	}

	bd, err := ioutil.ReadAll(res.Body)
	fmt.Println("response received:", string(bd))

	return string(bd), nil
}

// saveToDb will save the data in InfluxDB
func saveToDb(data map[string]string) bool {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	dbd := make(map[string]string)
	for k, v := range data {
		if k == "EmployeeID" {
			dbd["employee_id"] = v
		}
		if k == "EmployeeName" {
			dbd["employee_name"] = v
		}
	}
	fmt.Println("influxdb data is:", dbd)
	reqBody, err := json.Marshal(dbd)
	if err != nil {
		fmt.Printf("Error marshalling json : %s", err)
		return false
	}

	// Calling the write api to save the data in DB
	_, err = client.Post(dep+"/write", "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		fmt.Printf("Failed to make request: %s", err)
		return false
	}

	return true
}
