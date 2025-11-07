package main

import (
	"fmt"
	"log"

	"github.com/dwikikf/agviano-core-api-golang/internal/config"
	"github.com/dwikikf/agviano-core-api-golang/internal/delivery/http"
)

func main() {
	// Application entry point
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatalf("❌ Gagal memuat konfigurasi: %v", err)
	}

	fmt.Printf("🚀 Memulai %s di lingkungan %s pada port %s, %s\n", config.AppName, config.AppEnv, config.AppPort, config.Message)

	router := http.NewRouter()
	router.Run(":" + config.AppPort)
}
