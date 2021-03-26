package backendoperations

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func fetchFromCache(key string) (bool, string) {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	// Calling the get api to fetch the key value
	// from cache
	url := cep + "/get?employee_id=" + key
	fmt.Println("url is:", url)
	res, err := client.Get(url)
	bd, err := ioutil.ReadAll(res.Body)

	fmt.Println("response from url is:", string(bd))
	if err != nil {
		return false, ""
	}

	if strings.Contains(string(bd), "Key Not Found") {
		return false, ""
	}
	return true, string(bd)
}

// saveInCache will save the data in cache
func saveInCache(data map[string]string) bool {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	cd := make(map[string]string)
	for k, v := range data {
		if k == "EmployeeID" {
			cd["employee_id"] = v
		}
		if k == "EmployeeName" {
			cd["employee_name"] = v
		}
	}

	reqBody, err := json.Marshal(cd)
	if err != nil {
		fmt.Printf("Error marshalling json : %s", err)
		return false
	}

	// Calling the set api to save the key: value
	// in cache
	res, err := client.Post(cep+"/write", "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		fmt.Printf("Failed to make request: %s", err)
		return false
	}
	bd, err := ioutil.ReadAll(res.Body)
	val := string(bd)
	for _, v := range data {
		if v == val {
			return true
		}
	}
	return false
}
