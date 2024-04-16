package tools

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"

	log "github.com/sirupsen/logrus"
)

var DB *sql.DB

func NewDatabase() (*sql.DB, error) {
	// DB connection config, if you couldn't guess
    cfg := mysql.Config{
        User:   "root",
        Passwd: "letterbookd",
        Net:    "tcp",
        Addr:   "127.0.0.1:3306",
        DBName: "letterbookd",
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


	return newDB, nil
}
