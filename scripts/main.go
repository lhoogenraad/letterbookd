package main

import (
	"scripts/books/bookupload"
	"scripts/util"
)

func setupDatabase() error{
	var dbErr error
	util.DB, dbErr = util.NewDatabase()
	if dbErr != nil{
		panic(dbErr)
	}
	return dbErr
}

func main(){
	setupDatabase()
	err := bookupload.UpdateBooksSynopses()
	if err != nil {panic(err)}
}
