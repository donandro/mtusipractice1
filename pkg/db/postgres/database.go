package db

import (
	"fmt"
	"pillsreminder/pkg/db/postgres/config"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DefaultDatabase *Database

type Database struct {
	connection *gorm.DB
}

func InitDB(configuration config.PqConfig) {
	database := configuration.Dbname
	username := configuration.Username
	password := configuration.Password
	host := configuration.Host
	port := configuration.Port

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s",
		host, port, username, password, database)

	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		panic(err)
	}
	sqlDB, err := db.DB()

	sqlDB.SetMaxIdleConns(configuration.MaxIdleConns)
	sqlDB.SetMaxOpenConns(configuration.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Duration(configuration.MaxLifetime) * time.Second)

	DefaultDatabase = &Database{connection: db}
}

func (db *Database) GetDB() *gorm.DB {
	return db.connection
}
