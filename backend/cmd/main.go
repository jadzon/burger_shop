package main

import (
	"fmt"
	"log"

	"github.com/jadzon/burger_shop/internal/app"
	"github.com/jadzon/burger_shop/internal/config"
	"github.com/jadzon/burger_shop/internal/router"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}
	application := app.NewApplication(cfg)
	_ = application
	// Your application logic here
	log.Printf("Server starting on %s:%s", cfg.ServerHost, cfg.ServerPort)

	srv := router.NewRouter(cfg)
	addr := fmt.Sprintf("%s:%s", cfg.ServerHost, cfg.ServerPort)
	srv.POST("/register")
	srv.Run(addr)
}
