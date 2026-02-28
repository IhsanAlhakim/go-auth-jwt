package database

import (
	"database/sql"
	"log"

	"github.com/IhsanAlhakim/go-auth-api/internal/config"
	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

func Connect(cfg *config.Config) (*sql.DB, error) {
	config := mysql.NewConfig()
	config.User = cfg.DBUsername
	config.Passwd = cfg.DBPassword
	config.Net = "tcp"
	config.Addr = cfg.DBAddr
	config.DBName = cfg.DBName

	var err error

	db, err = sql.Open("mysql", config.FormatDSN())
	if err != nil {
		return nil, err
	}

	log.Println("Connected to database")
	return db, nil
}
