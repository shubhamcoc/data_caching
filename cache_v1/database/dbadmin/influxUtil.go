package dbmanager

import (
	"bytes"
	"fmt"
	"net"
	"time"

	"github.com/influxdata/influxdb/client/v2"
)

//CreateHttpClient to open a connection to InfluxDB
func CreateHttpClient(host, port, userName, passwd string) (client.Client, error) {

	var buff bytes.Buffer

	fmt.Fprintf(&buff, "http://%s:%s", host, port)
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     buff.String(),
		Username: userName,
		Password: passwd})

	return c, err
}

//CreateAdminUser to create an admin user in InfluxDB
func CreateAdminUser(c client.Client, userName, passwd string) (*client.Response, error) {

	var buff bytes.Buffer

	fmt.Fprintf(&buff, "CREATE USER %s WITH PASSWORD '%s' WITH ALL PRIVILEGES", userName, passwd)
	q := client.NewQuery(buff.String(), "", "")
	res, err := c.Query(q)
	return res, err
}

//CreateDB to create a database in InfluxDB
func CreateDB(c client.Client, dbName string) (*client.Response, error) {

	var buff bytes.Buffer

	fmt.Fprintf(&buff, "CREATE DATABASE %s", dbName)
	q := client.NewQuery(buff.String(), "", "")
	res, err := c.Query(q)
	return res, err
}

//CreateSubs to create a subscription to a db in InfluxDB
func CreateSubs(c client.Client, sName, dbName, host, port string) (*client.Response, error) {

	var buff bytes.Buffer

	fmt.Fprintf(&buff, "CREATE SUBSCRIPTION %s ON \"%s\".\"autogen\" DESTINATIONS ALL 'http://%s:%s'", sName, dbName, host, port)
	q := client.NewQuery(buff.String(), "", "")
	res, err := c.Query(q)
	return res, err
}

//CheckPort to check if InfluxDB server is up
func CheckPort(host, port string) bool {
	retry := 100
	count := 0
	status := false
	fmt.Printf("Waiting for Port: %s on hostname: %s \n", port, host)
	for count < retry {
		conn, _ := net.DialTimeout("tcp", net.JoinHostPort(host, port), (5 * time.Second))
		if conn != nil {
			fmt.Printf("Port: %s on hostname: %s is up.", port, host)
			conn.Close()
			status = true
			break
		}
		time.Sleep(100 * time.Millisecond)
		count++
	}
	return status
}
