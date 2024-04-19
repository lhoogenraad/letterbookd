package models

import (
	"server/internal/tools"
	
	log "github.com/sirupsen/logrus"
)


func AddUser(
	email string,
	passwordHash string,
	firstName string,
	lastName string,
) (error, int) {
	var insertQuery string = `
	INSERT into users
	(email, password_hash, first_name, last_name)
	VALUES
	(?, ?, ?, ?)`

	_, err := tools.DB.Exec(
		insertQuery,
		email,
		passwordHash,
		firstName,
		lastName,
	)

	if err != nil {
		log.Error(err)
		return err, 500
	}

	return nil, 0
}
