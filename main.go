package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/tomorrow-code-challenge/backend-wdata-code-challenge.git/config"
	"github.com/tomorrow-code-challenge/backend-wdata-code-challenge.git/core"
	"github.com/tomorrow-code-challenge/backend-wdata-code-challenge.git/definition"
	"github.com/tomorrow-code-challenge/backend-wdata-code-challenge.git/server"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const maxRetries = 5

func main() {
	pgDB := initDB()
	defer pgDB.Close()

	weatherForecast := initWeatherForecast(pgDB)

	server.StartHTTP(&weatherForecast)

	//Handle OS signals for graceful shutdown
	gracefulShutdown := make(chan os.Signal, 1)
	signal.Notify(gracefulShutdown, syscall.SIGINT, syscall.SIGTERM)

	// Block until a signal is received
	_ = <-gracefulShutdown
	shutDown(pgDB)
}

func initWeatherForecast(db *sql.DB) definition.IWeatherForecast {
	return core.NewPgWeatherForecast(db)
}

func initDB() *sql.DB {
	pgInfo := fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable",
		config.Static.PGHost, config.Static.PGPort, config.Static.PGUser, config.Static.PGPassword, config.Static.DbName)
	for retries := 0; retries < maxRetries; retries++ {
		db, err := sql.Open("postgres", pgInfo)
		if err != nil {
			log.Printf("Failed to connect to PostgreSQL: %v. Retrying...", err)
			time.Sleep(time.Second * 2) // Wait few seconds before retrying
			continue
		}

		err = db.Ping()
		if err != nil {
			log.Printf("Failed to ping PostgreSQL: %v. Retrying...", err)
			time.Sleep(time.Second * 2) // Wait few seconds before retrying
			continue
		}

		log.Println("Successfully connected to PostgreSQL!")
		return db
	}

	log.Println("Maximum retries exceeded. Unable to connect to PostgreSQL.")
	os.Exit(1)
	return nil
}

func shutDown(db *sql.DB) {
	log.Println("Shutting down server...")
	server.Shutdown()

	log.Println("Disconnecting postgreSQL client...")
	db.Close()

	log.Println("Server gracefully stopped")
}
