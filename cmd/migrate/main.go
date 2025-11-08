package main

import (
	"fmt"
	"log"
	"os"

	"github.com/dwikikf/agviano-core-api-golang/internal/config"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatalf("❌ Gagal memuat konfigurasi: %v", err)
	}

	// fmt.Printf("🚀 Memulai migrasi database untuk %s di lingkungan %s\n", config.AppName, config.AppEnv)
	// fmt.Printf(" user : %s, Pwd : %s, host : %s, post : %s, db : %s \n", config.Database.User, config.Database.Password, config.Database.Host, config.Database.Port, config.Database.Name)

	dbURL := fmt.Sprintf("mysql://%s:%s@tcp(%s:%s)/%s?multiStatements=true&parseTime=true", config.Database.User, config.Database.Password, config.Database.Host, config.Database.Port, config.Database.Name)

	// Kode migrasi database di sini
	if len(os.Args) < 2 {
		log.Fatalf("❌ --> Pemakaian: go run cmd/migrate/main.go [up|down|version|force <ver>]")
	}

	action := os.Args[1]

	m, err := migrate.New(
		"file://database/migrations",
		dbURL,
	)
	if err != nil {
		log.Fatalf("❌ Gagal membuat instance migrasi: %v", err)
	}

	switch action {
	case "up":
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatalf("❌ Gagal menjalankan migrasi: %v", err)
		}
		fmt.Println("✅ Migrasi berhasil dijalankan.")

	case "down":
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			log.Fatalf("❌ Gagal menjalankan rollback migrasi: %v", err)
		}
		fmt.Println("✅ Rollback migrasi berhasil dijalankan.")

	case "version":
		version, dirty, err := m.Version()
		if err != nil {
			log.Fatalf("❌ Gagal mendapatkan versi migrasi: %v", err)
		}
		fmt.Printf("🔢 Versi migrasi saat ini: %d, Dirty: %v\n", version, dirty)

	case "force":
		if len(os.Args) < 3 {
			log.Fatalf("❌ --> Pemakaian: go run cmd/migrate/main.go force <ver>")
		}
		var ver int
		_, err := fmt.Sscanf(os.Args[2], "%d", &ver)
		if err != nil {
			log.Fatalf("❌ Versi tidak valid: %v", err)
		}
		if err := m.Force(ver); err != nil {
			log.Fatalf("❌ Gagal memaksa versi migrasi: %v", err)
		}
		fmt.Printf("✅ Versi migrasi dipaksa ke: %d\n", ver)

	default:
		log.Fatalf("❌ Aksi tidak dikenal: %s. Gunakan salah satu dari [up|down|version|force <ver>]", action)
	}
}
