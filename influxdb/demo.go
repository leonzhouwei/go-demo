package main

import (
	"fmt"
	"log"

	"github.com/influxdb/influxdb/client/v2"
)

func main() {
	// Make client
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr: "http://localhost:8086",
	})
	if err != nil {
		fmt.Println("Error creating InfluxDB Client: ", err.Error())
	}
	defer c.Close()

	q := client.NewQuery("SELECT * FROM pandora_agent", "qiniu", "")
	response, err := c.Query(q)
	if err != nil {
		log.Fatal(err)
	}
	if response.Error() != nil {
		log.Fatal(response.Error())
	}
	fmt.Println(response.Results)
}