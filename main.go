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
)

func main() {
	pgDB, err := initDB()
	if err != nil {
		log.Fatal(err)
	}
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

func initDB() (*sql.DB, error) {
	pgInfo := fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable",
		config.Static.PGHost, config.Static.PGPort, config.Static.PGUser, config.Static.PGPassword, config.Static.DbName)

	db, err := sql.Open("postgres", pgInfo)
	if err != nil {
		log.Printf("Failed to connect to PostgreSQL: %v", err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Printf("Failed to ping PostgreSQL: %v", err)
		return nil, err
	}

	log.Println("Successfully connected to PostgreSQL!")
	return db, nil
}

func shutDown(db *sql.DB) {
	log.Println("Shutting down server...")
	server.Shutdown()

	log.Println("Disconnecting postgreSQL client...")
	db.Close()

	log.Println("Server gracefully stopped")
}
