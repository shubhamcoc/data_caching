package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"io/ioutil"
	"time"
	"bytes"
	"strings"
)

const (
	cep = "http://em_cache:4000"
	dep = "http://em_database:7000"
)

var store map[string]string

func main() {

	http.HandleFunc("/api/submit", submit)
	http.HandleFunc("/api/searchbyid", readId)
	http.HandleFunc("/api/search", readRecent)
	http.HandleFunc("/api/search/next", readNext)

	// Starting the Http Server at port 6000
        http.ListenAndServe(":6000", nil)
}

// Http handler to /api/searchbyid endpoint
func readId (w http.ResponseWriter, r *http.Request) {
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

	// Get the data from Cache
	status, val := fetchFromCache(key)
	if status {
	    fmt.Fprintf(w, val)
	    return
	}

	// Get the data from InfluxDB
	val1, err := fetchIdFromDb(key)
	if err != nil {
		http.Error(w, http.StatusText(400), 400)
		return
	}
	if val1 != "" {
		fmt.Fprintf(w, val1)
		// Save the result in the Cache
		saveInCache(store)
		return
	}
	fmt.Fprintf(w, "No record found")
}

// Http handler to /api/search endpoint
func readRecent(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodGet {
                w.Header().Set("Allow", http.MethodGet)
                http.Error(w, http.StatusText(405), 405)
                return
        }

	client := &http.Client{
                        Timeout: 10 * time.Second,
        }

	_, err := client.Get(dep + "/load")
        if err != nil {
                fmt.Printf("Failed to make request: %s", err)
		return
        }

	time.Sleep(10 * time.Second)
	body := make(map[string]interface{})

	body["key"] = "search"
	body["start"] = 1
	body["stop"] = 10

	reqBody, err := json.Marshal(body)
	if err != nil {
                fmt.Printf("Error marshalling json : %s", err)
		return
        }

	// fetch the unique resuts in 10 search result from cache
	fmt.Println("send body:", reqBody)
	res, err := client.Post(cep + "/search", "application/json", bytes.NewBuffer(reqBody))
        if err != nil {
                fmt.Printf("Failed to make request: %s", err)
		return
        }

	bd, _ := ioutil.ReadAll(res.Body)
	fmt.Fprintf(w, string(bd))
}

func readNext (w http.ResponseWriter, r *http.Request) {
        client := &http.Client{
                        Timeout: 10 * time.Second,
        }

	body := make(map[string]interface{})

        body["key"] = "search"
        body["start"] = 11
        body["stop"] = 20

        reqBody, err := json.Marshal(body)
        if err != nil {
                fmt.Printf("Error marshalling json : %s", err)
        }

        // fetch the unique resuts in 11-20 search result from cache
        fmt.Println("send body:", reqBody)
        res, err := client.Post(cep + "/search", "application/json", bytes.NewBuffer(reqBody))
        if err != nil {
                fmt.Printf("Failed to make request: %s", err)
        }

        bd, _ := ioutil.ReadAll(res.Body)

        fmt.Fprintf(w, string(bd))
}

// Http handler to /api/submit endpoint
func submit (w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Received a post request\n")
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, http.StatusText(405), 405)
		return
	}
        body, err := ioutil.ReadAll(r.Body)
	var data map[string]string
	err = json.Unmarshal(body, &data)
        if err != nil {
		http.Error(w, http.StatusText(400), 400)
                fmt.Printf("Error while json Unmarshal: %v\n", err.Error())
                return
        }
	fmt.Printf("Data parsed as: %v\n", data)

	//Save the data in Redis cahce 
	success := saveInCache(data)
	if success {
	       fmt.Fprintf(w, "successfully set in cache\n")
	}else{
	       fmt.Fprintf(w, "failed to set in cache\n")
	}

	//Save the data in influxDB
	success = saveToDb(data)
	if success {
	       fmt.Fprintf(w, "successfully set in db\n")
	}else{
	       fmt.Fprintf(w, "failed to set in db\n")
	}
}

func fetchFromCache(key string) (bool, string) {
	client := &http.Client{
                        Timeout: 10 * time.Second,
        }

	// Calling the get api to fetch the key value
	// from cache
	url := cep + "/get?key=" + key
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
	reqBody, err := json.Marshal(data)
        if err != nil {
                fmt.Printf("Error marshalling json : %s", err)
		return false
        }

	// Calling the set api to save the key: value 
	// in cache
	res, err := client.Post(cep + "/set", "application/json", bytes.NewBuffer(reqBody))
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

// saveToDb will save the data in InfluxDB
func saveToDb(data map[string]string) bool {
	client := &http.Client{
                        Timeout: 10 * time.Second,
        }
	dbd := make(map[string]string)
	for k, v := range data {
		if k == "key" {
		    dbd["id"] = v
		}
		if k == "value" {
		    dbd["name"] = v
		}
	}
	fmt.Println("influxdb data is:", dbd)
	reqBody, err := json.Marshal(dbd)
        if err != nil {
                fmt.Printf("Error marshalling json : %s", err)
                return false
        }

	// Calling the write api to save the data in DB
	_, err = client.Post(dep + "/write", "application/json", bytes.NewBuffer(reqBody))
        if err != nil {
                fmt.Printf("Failed to make request: %s", err)
                return false
        }

	return true
}

// fetchIdFromDb will fetch the data from InfluxDB with the provided key
// It get the results from query and extract the required value from result
func fetchIdFromDb(key string) (string, error) {

	query := fmt.Sprintf("SELECT \"name\" FROM DEMO WHERE \"id\" = '%s'", key)
	fmt.Println("query formed", query)

	// Calling the fetchFromDb with the created query
	val, err := fetchFromDb(query)
	if err != nil {
		return "", err
	}
	str := val["Data"].(interface{})
	str1 := str.(string)
	var data map[string]interface{}
	err = json.Unmarshal([]byte(str1), &data)
	if err != nil {
		fmt.Println("unable to unmarshal data:", err)
		return "", err
	}

	// extract the name from the query results
	str = data["values"].(interface{})
	str3 := str.([]interface{})
	str4 := str3[0].([]interface{})

	store = make(map[string]string)

	store["key"] = key
	store["value"] = str4[1].(string)
	return str4[1].(string), nil
}

// fetchFromDb common func to send the query to InfluxDB
func fetchFromDb(query string) (map[string]interface{}, error) {
	client := &http.Client{
                        Timeout: 10 * time.Second,
        }
	emptyres := make(map[string]interface{})
	reqBody := []byte(query)
	res, err := client.Post(dep + "/read", "application/text", bytes.NewBuffer(reqBody))
        if err != nil {
                fmt.Printf("Failed to make request: %s", err)
                return emptyres, err
        }

        bd, err := ioutil.ReadAll(res.Body)
	fmt.Println("response received:",  string(bd))
	data := make(map[string]interface{})
	err = json.Unmarshal(bd, &data)
        if err != nil {
                fmt.Printf("Failed to Unmarshal request: %s", err)
                return emptyres, err
        }

        return data, nil
}
