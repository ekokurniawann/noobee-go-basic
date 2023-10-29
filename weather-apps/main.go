package main

import (
	"fmt"
	"log"
	"net/http"
	"weather-apps/app"
)

func main() {

	port := ":4444"

	http.HandleFunc("/weather", app.GetCurrentWeather)

	log.Println("server running at port", port)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}
