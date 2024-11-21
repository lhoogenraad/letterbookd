package main

import (
	"scripts/util"
	"scripts/authors"
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
	err := authors.ReadAndUpload()
	if err != nil {panic(err)}
}
