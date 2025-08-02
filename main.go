package main

import (
	config "kumande/configs"
	"kumande/modules"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func initLogging() {
	now := time.Now()
	logFileName := "logs/Kumande-" + now.Format("January-2006") + ".log"

	f, err := os.OpenFile(logFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}

	log.SetOutput(f)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func main() {
	initLogging()
	log.Println("Kumande service is starting...")

	// Load Env
	err := godotenv.Load()
	if err != nil {
		panic("Error loading ENV")
	}

	// Connect DB
	db := config.ConnectDatabase()
	modules.MigrateAll(db)

	// Setup Gin
	router := gin.Default()
	redisClient := config.InitRedis()

	modules.SetUpDependency(router, db, redisClient)

	// Run server
	port := os.Getenv("PORT")
	router.Run(":" + port)

	log.Printf("Kumande is running on port %s\n", port)
}
