package app

import (
	"log"

	"github.com/jadzon/burger_shop/internal/config"
	"github.com/jadzon/burger_shop/internal/database"
	"github.com/jadzon/burger_shop/internal/repositories"
	"github.com/jadzon/burger_shop/internal/services"
)

type Application struct {
	userService services.UserService
}

func NewApplication(cfg *config.Config) *Application {
	// Connect to database with GORM
	db, err := database.ConnectDB(
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
		cfg.Debug,
	)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	userRepo := repositories.NewUserRepository(*db)
	userService := services.NewUserService(userRepo)
	// Auto migrate your models
	// if err := db.AutoMigrate(&User{}); err != nil {
	// 	log.Fatal("Failed to migrate database:", err)
	// }
	return &Application{userService: userService}
}
