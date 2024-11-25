package main

import (
	"scripts/util"
	"scripts/books/bookupload"
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
	err := bookupload.ReadAndUpload()
	if err != nil {panic(err)}
}
