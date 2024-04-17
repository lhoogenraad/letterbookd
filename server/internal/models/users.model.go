package models

import (
	"server/internal/tools"
	"server/internal/resources"
)


func AddUser(
	email string,
	passwordHash string,
	firstName string,
	lastName string
) (error, int) {
	var insertQuery string = `
	INSERT into users
	(email, password_hash, first_name, last_name)
	VALUES
	(?, ?, ?, ?)`

	rows, err := tools.DB.Query(
		insertQuery,
		email,
		passwordHash,
		firstName,
		lastName
	)

	
	if err != nil {
		return err, 500
	}

	return nil, 0
}
