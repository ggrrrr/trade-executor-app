package db

import (
	"database/sql"
	"log"
	"os"

	"github.com/sirupsen/logrus"

	_ "github.com/mattn/go-sqlite3" // Import go-sqlite3 library
)

func connectSqlite3(url string) error {
	var err error

	logrus.Infof("Creating file: %v", url)
	// file, err := os.Exi
	file, err := os.Create(url)
	if err != nil {
		log.Fatal(err.Error())
	}
	file.Close()

	conn, err = sql.Open("sqlite3", url)
	if err != nil {
		log.Fatal(err.Error())
	}
	// defer conn.Close()
	// Create Database Tables
	return createTable()
}
