package main

import (
	"fmt"
	"log"

	"github.com/dwikikf/agviano-core-api-golang/internal/config"
	"github.com/dwikikf/agviano-core-api-golang/internal/delivery/http"
	"github.com/dwikikf/agviano-core-api-golang/internal/delivery/http/handler"
	"github.com/dwikikf/agviano-core-api-golang/internal/repository"
	"github.com/dwikikf/agviano-core-api-golang/internal/usecase"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatalf("❌ Gagal memuat konfigurasi: %v", err)
	}

	// setup database
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Database.User,
		config.Database.Password,
		config.Database.Host,
		config.Database.Port,
		config.Database.Name,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("❌ Gagal menghubungkan ke database: %v", err)
	}

	//dependency injection
	categoryRepo := repository.NewCategoryGormRepository(db)
	categoryUC := usecase.NewCategoryService(categoryRepo)
	categoryHandler := handler.NewCategoryHandler(categoryUC)

	productRepo := repository.NewProductGormRepository(db)
	productUC := usecase.NewProductService(productRepo, categoryRepo)
	productHandler := handler.NewProductHandler(productUC)

	fmt.Printf("🚀 Memulai %s di lingkungan %s pada port %s, %s\n", config.AppName, config.AppEnv, config.AppPort, config.Message)

	router := http.NewRouter(categoryHandler, productHandler)
	router.Run(":" + config.AppPort)
}
