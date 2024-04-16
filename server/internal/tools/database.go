package tools

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"

	log "github.com/sirupsen/logrus"
)


func NewDatabase() (error) {
	// DB connection config, if you couldn't guess
    cfg := mysql.Config{
        User:   "root",
        Passwd: "letterbookd",
        Net:    "tcp",
        Addr:   "127.0.0.1:3306",
        DBName: "letterbookd",
    }
	// Get a database handle.
	var db *sql.DB
    var err error
    db, err = sql.Open("mysql", cfg.FormatDSN())
    if err != nil {
        log.Fatal(err)
		return err
    }

    pingErr := db.Ping()
    if pingErr != nil {
        log.Fatal(pingErr)
		return pingErr
    }
    log.Info("Connected to database!")


	return nil
}
