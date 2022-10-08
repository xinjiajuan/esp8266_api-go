package main

import (
	config "DHT22-temperature_databases-go/config"
	"DHT22-temperature_databases-go/http"
	"os"
)

func main() {
	Config := config.ReadConfig(os.Args[1])
	http.RunHttpAPIServer(http.MakeHttpServer(Config))
}
