package models

import (
	"server/internal/tools"
	"strings"
	"errors"
	"fmt"

	log "github.com/sirupsen/logrus"
)


func handleInsertUserError(err error) (error, int) {
	log.Error(err)
	returnErr := errors.New("Sorry, something went wrong creating your account.")
	status := 500
	if strings.Contains(fmt.Sprint(err), "users.email_unique") {
		returnErr = errors.New("Sorry, an account with that email already exists.")
		status = 400
	}
	return returnErr, status
}

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
		err, code := handleInsertUserError(err)
		return err, code
	}

	return nil, 0
}

