package dbmanager

import (
	"encoding/json"
	"errors"
	"fmt"
	"os/exec"
	"reflect"
	"strings"
	"time"

	mqtt "database/mqttpub"

	"github.com/influxdata/influxdb/client/v2"
)

type InfluxDBAdmin struct {
	Host     string
	Port     string
	DbName   string
	Passwd   string
	UserName string
	Mq       mqtt.Mqtt
	buffer   []interface{}
}

// Init will start the InfluxDB and create an Admin User
func (ad *InfluxDBAdmin) Init() error {

	cmd := exec.Command("./database/script.sh")
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Failed to start influxdb Server, Error: %s\n", err)
		return err
	}

	ps := CheckPort(ad.Host, ad.Port)
	if !ps {
		return errors.New("InfluxDB server is not up")
	}

	c, err := CreateHttpClient(ad.Host, ad.Port, "", "")
	if err != nil {
		fmt.Printf("Failed to create influxdb connection, Error: %s\n", err)
		return err
	}

	res, err := CreateAdminUser(c, ad.UserName, ad.Passwd)
	if err == nil && res.Error() == nil {
		fmt.Printf("Successfully created admin user: %s\n", ad.UserName)
	} else {
		if res != nil && res.Error() != nil {
			fmt.Printf("User already exist")
		} else {
			fmt.Printf("Failed to create Admin user, Error: %v\n", err)
		}
		return err
	}

	c.Close()
	return nil
}

// CreateDatabase will create the database in InfluxDB
func (ad *InfluxDBAdmin) CreateDatabase() error {
	c, err := CreateHttpClient(ad.Host, ad.Port, ad.UserName, ad.Passwd)
	if err != nil {
		fmt.Printf("Failed to create influxdb connection, Error: %s\n", err)
		return err
	}
	defer c.Close()

	res, err := CreateDB(c, ad.DbName)
	if err == nil && res.Error() == nil {
		fmt.Printf("Successfully created Database: %s\n", ad.DbName)
	} else {
		fmt.Printf("Failed to create Database, Error: %s\n", res.Error())
		return err
	}
	return nil
}

// InsertData will write the data to InfluxDB
func (ad *InfluxDBAdmin) InsertData(msg []byte) {
	fmt.Println("received data:", string(msg))
	tags := make(map[string]string)
	field := make(map[string]interface{})
	data := make(map[string]interface{})

	err := json.Unmarshal(msg, &data)

	if err != nil {
		fmt.Printf("\n Not able to Parse data %s", err.Error())
	}

	for key, value := range data {
		if reflect.ValueOf(value).Type().Kind() == reflect.Float64 {
			field[key] = value
		} else if reflect.ValueOf(value).Type().Kind() == reflect.String {
			field[key] = value
		} else if reflect.ValueOf(value).Type().Kind() == reflect.Bool {
			field[key] = value
		} else if reflect.ValueOf(value).Type().Kind() == reflect.Int {
			field[key] = value
		}
	}

	Measurement := "demo"

	clientadmin, err := CreateHttpClient(ad.Host, ad.Port, ad.UserName, ad.Passwd)
	defer clientadmin.Close()

	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  ad.DbName,
		Precision: "ns",
	})

	if err != nil {
		fmt.Printf("\n Error in creating batch point %s", err.Error())
	}

	pt, err := client.NewPoint(Measurement, tags, field, time.Now())
	if err != nil {
		fmt.Printf("\n point error %s", err.Error())
	}

	bp.AddPoint(pt)

	if err := clientadmin.Write(bp); err != nil {
		fmt.Printf("\n Write Error %s", err.Error())
	}

	ad.buffer = append(ad.buffer, data)

}

// QueryDb will execute the query and send the results as output
func (ad *InfluxDBAdmin) QueryDb(query string) (map[string]interface{}, error) {

	clientadmin, err := CreateHttpClient(ad.Host, ad.Port, ad.UserName, ad.Passwd)
	defer clientadmin.Close()

	query = strings.ToLower(query)

	q := client.Query{
		Command:   query,
		Database:  ad.DbName,
		Precision: "s",
	}
	if response, err := clientadmin.Query(q); err == nil {
		if response.Error() != nil {
			fmt.Println("Response Error received:", response.Error())
		} else {
			if len(response.Results[0].Series) > 0 {
				output := response.Results[0].Series[0]
				fmt.Println(output)
				Output, err := json.Marshal(output)
				response := map[string]interface{}{"Data": string(Output)}
				return response, err
			}
		}
	}

	val := map[string]interface{}{"Data": ""}
	err = errors.New("Response is nil")
	return val, err
}

// LoadCache will publish the data store in the InfluxDB
func (ad *InfluxDBAdmin) LoadCache() bool {
	fmt.Println("data is:", ad.buffer)
	if len(ad.buffer) > 0 {
		cmd := make(map[string]interface{})
		cmd["load"] = ad.buffer
		bytedata, err := json.Marshal(cmd)
		if err != nil {
			fmt.Println("unable to marshal data:", err)
			return false
		}
		ad.Mq.Publisher(bytedata)
		ad.buffer = nil
		return true
	}
	return false

}
