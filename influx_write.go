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
	writeToInflux(payload, c.Param("serial"))
	c.AbortWithStatus(200)
}

func writeToInflux(payload Payload, serial string) {
	var pts = make([]client.Point, 2)
	pts[0] = client.Point{
		Measurement: "vaayu_device_readings",
		Tags: map[string]string{
			"serial_number": serial,
		},
		Fields: map[string]interface{}{
			"loc":              ParseFloat(payload.EventParameters["loc"], 32),
			"sid":              ParseFloat(payload.EventParameters["sid"], 32),
			"sig":              ParseFloat(payload.EventParameters["sig"], 32),
			"bat":              ParseFloat(payload.EventParameters["bat"], 32),
			"temperature":      ParseFloat(payload.EventParameters["a"], 32),
			"pressure":         ParseFloat(payload.EventParameters["b"], 32),
			"altitude":         ParseFloat(payload.EventParameters["c"], 32),
			"humidity":         ParseFloat(payload.EventParameters["d"], 32),
			"ammonia":          ParseFloat(payload.EventParameters["e"], 32),
			"carbon_monoxide":  ParseFloat(payload.EventParameters["f"], 32),
			"nitrogen_dioxide": ParseFloat(payload.EventParameters["g"], 32),
			"propane":          ParseFloat(payload.EventParameters["h"], 32),
			"butane":           ParseFloat(payload.EventParameters["i"], 32),
			"methane":          ParseFloat(payload.EventParameters["j"], 32),
			"hydrogen":         ParseFloat(payload.EventParameters["k"], 32),
			"ethanol":          ParseFloat(payload.EventParameters["l"], 32),
			"carbon_dioxide":   ParseFloat(payload.EventParameters["m"], 32),
			"voc":              ParseFloat(payload.EventParameters["n"], 32),
			"pm_2_5":           ParseFloat(payload.EventParameters["q"], 32),
			"pm_10":            ParseFloat(payload.EventParameters["r"], 32),
			"particle_0_3":     ParseFloat(payload.EventParameters["s"], 32),
			"particle_1":       ParseFloat(payload.EventParameters["t"], 32),
			"particle_2_5":     ParseFloat(payload.EventParameters["u"], 32),
			"particle_5":       ParseFloat(payload.EventParameters["v"], 32),
			"particle_10":      ParseFloat(payload.EventParameters["w"], 32),
			"pm_1":             ParseFloat(payload.EventParameters["x"], 32),
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
	if s == "" || s == " " {
		return 0
	}
	res, err := strconv.ParseFloat(s, bitSize)
	if err != nil {
		panic(err)
	}
	return res
}
