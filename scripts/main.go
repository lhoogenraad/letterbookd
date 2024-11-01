package main

import (
	"scripts/util"
	"scripts/books/covers"
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
	err := covers.AddCoversToBooks()
	if err != nil {panic(err)}
}
