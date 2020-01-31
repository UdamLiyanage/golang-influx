package main

import (
	_ "github.com/influxdata/influxdb1-client"
	client "github.com/influxdata/influxdb1-client"
	"net/url"
	"os"
)

func connect() {
	u, err := url.Parse(os.Getenv("INFLUX_URL"))
	if err != nil {
		panic(err)
	}
	influx, err = client.NewClient(client.Config{URL: *u})
	if err != nil {
		panic(err)
	}
}
