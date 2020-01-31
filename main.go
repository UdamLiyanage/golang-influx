package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/influxdata/influxdb1-client"
	client "github.com/influxdata/influxdb1-client"
	_ "github.com/joho/godotenv/autoload"
	"log"
)

var influx *client.Client

func init() {
	connect()
}

func main() {
	router := gin.Default()

	router.POST("/payload", payloadHandler)
	log.Fatal(router.Run(":9000"))
}
