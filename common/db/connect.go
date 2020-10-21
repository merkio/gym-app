package db

import (
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

// Connect create connection to the database
func Connect() {
	dsn := "user=postgres dbname=gym_app host=127.0.0.1 port=5432 sslmode=disable TimeZone=Europe/Berlin"
	connect, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Can't connect to the database\n%s", err)
	}
	db = connect
	sqlDB, _ := db.DB()

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)
}

// GetDB return db connection
func GetDB() *gorm.DB {
	return db
}
