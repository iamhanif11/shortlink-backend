package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/iamhanif11/shortlink-backend.git/internal/config"
	"github.com/iamhanif11/shortlink-backend.git/internal/router"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Printf("Error loading env. \ncause: %s", err.Error())
	}
	// if err := os.MkdirAll(filepath.Join("public", "img"), os.ModePerm); err != nil {
	// 	log.Fatalf("Failed to create upload directory: %s", err.Error())
	// }
	// inisialisasi
	// gin.New()
	app := gin.Default()
	// app.Use(middleware.CORSMiddleware)
	// connect ke db
	db, err := config.ConnectPsql()
	if err != nil {
		log.Fatalf("DB connection error. \ncause: %s", err.Error())
	}
	defer db.Close()
	log.Println("DB Connected")
	// connect ke redis
	rc, err := config.ConnectRedis()
	if err != nil {
		log.Fatalf("Redis connection error. \ncause: %s", err.Error())
	}
	defer rc.Close()
	log.Println("Redis Connected")

	// install router
	router.InitRouter(app, db, rc)
	// run
	// addr := fmt.Sprintf("%s:%s", os.Getenv("APP_HOST"), os.Getenv("APP_PORT"))
	app.Run(fmt.Sprintf("%s:%s", os.Getenv("APP_HOST"), os.Getenv("APP_PORT")))
}
