package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/influxdata/influxdb1-client"
	client "github.com/influxdata/influxdb1-client"
	_ "github.com/joho/godotenv/autoload"
	"os"
	"strconv"
)

type Payload struct {
	EventParameters map[string]string `json:"eventParameters"`
	DeviceID        string            `json:"deviceId"`
}

func payloadHandler(c *gin.Context) {
	var payload Payload
	err := json.NewDecoder(c.Request.Body).Decode(&payload)
	if err != nil {
		fmt.Println("Error", err)
		c.AbortWithStatus(500)
	}
	writeToInflux(payload)
	c.AbortWithStatus(200)
}

func writeToInflux(payload Payload) {
	var pts = make([]client.Point, 2)
	pts[0] = client.Point{
		Measurement: "device_readings",
		Tags: map[string]string{
			"serial_number": payload.DeviceID,
		},
		Fields: map[string]interface{}{
			"Temp1":   ParseFloat(payload.EventParameters["Temp1"], 32),
			"Temp2":   ParseFloat(payload.EventParameters["Temp2"], 32),
			"ACPower": ParseFloat(payload.EventParameters["ACPower"], 32),
			"CH1":     ParseFloat(payload.EventParameters["CH1"], 32),
			"CH2":     ParseFloat(payload.EventParameters["CH2"], 32),
			"CH3":     ParseFloat(payload.EventParameters["CH3"], 32),
			"CH4":     ParseFloat(payload.EventParameters["CH4"], 32),
			"RSSI":    ParseFloat(payload.EventParameters["RSSI"], 32),
			"BAT":     ParseFloat(payload.EventParameters["BAT"], 32),
			"STAT":    payload.EventParameters["STAT"],
		},
	}
	bps := client.BatchPoints{
		Points:   pts,
		Database: os.Getenv("INFLUX_DATABASE"),
	}

	_, err := influx.Write(bps)
	if err != nil {
		panic(err)
	}
}

func ParseFloat(s string, bitSize int) float64 {
	res, err := strconv.ParseFloat(s, bitSize)
	if err != nil {
		panic(err)
	}
	return res
}

//ec2-34-203-213-57.compute-1.amazonaws.com
//34.203.213.57
