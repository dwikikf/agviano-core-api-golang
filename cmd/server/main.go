package main

import "github.com/dwikikf/agviano-core-api-golang/internal/delivery/http"

func main() {
	// Application entry point
	router := http.NewRouter()
	router.Run() // listen and serve on
}
