package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PORT      string
	SecretKey string
	MysqlClient *sql.DB
}

func Load() *Config {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error load env: %v", err)
	}

	dbUser := os.Getenv("DB_USER")
	// dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:@tcp(%s:%s)/%s?parseTime=true", dbUser, dbHost, dbPort, dbName)
	mysqlClient, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	if err := mysqlClient.Ping(); err != nil {
		log.Fatal("DB tidak bisa di-connect:", err)
	}

	fmt.Println("Berhasil connect ke MySQL!")

	return &Config{
		PORT: os.Getenv("PORT_APP"),
		SecretKey: os.Getenv("SECRET_KEY"),
		MysqlClient: mysqlClient,
	}
}