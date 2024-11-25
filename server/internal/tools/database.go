package tools

import (
	"database/sql"
	"time"
	"github.com/go-sql-driver/mysql"

	log "github.com/sirupsen/logrus"
)

var MAX_IDLE_CONNS int = 10
var MAX_OPEN_CONNS int = 40
var MAX_IDLE_MINUTES int = 3
var MAX_IDLE_DURATION time.Duration = time.Duration(MAX_IDLE_MINUTES) * time.Minute

var DB *sql.DB

func NewDatabase() (*sql.DB, error) {
	// DB connection config, if you couldn't guess
    cfg := mysql.Config{
        User:   "root",
        Net:    "tcp",
        Addr:   "127.0.0.1:3306",
        DBName: "letterbookd",
		Passwd: "password",
    }
	// Get a database handle.
    var err error
	newDB, err := sql.Open("mysql", cfg.FormatDSN())
    if err != nil {
        log.Fatal(err)
		return nil, err
    }

    pingErr := newDB.Ping()
    if pingErr != nil {
        log.Fatal(pingErr)
		return nil, pingErr
    }
    log.Info("Connected to database!")

	newDB.SetMaxIdleConns(MAX_IDLE_CONNS)
	newDB.SetMaxOpenConns(MAX_OPEN_CONNS)
	newDB.SetConnMaxIdleTime(MAX_IDLE_DURATION)

	return newDB, nil
}
