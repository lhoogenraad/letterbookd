package util

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"fmt"
)

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
        fmt.Println(err)
		return nil, err
    }

    pingErr := newDB.Ping()
    if pingErr != nil {
        fmt.Println(pingErr)
		return nil, pingErr
    }
    fmt.Println("Connected to database!")


	return newDB, nil
}
