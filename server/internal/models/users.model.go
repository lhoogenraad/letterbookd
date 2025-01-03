package models

import (
	"github.com/lhoogenraad/letterbookd/internal/tools"
	"github.com/lhoogenraad/letterbookd/internal/resources"
	"strings"
	"errors"
	"fmt"
	"database/sql"

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
) (int, error, int) {
	var insertQuery string = `
	INSERT into users
	(email, password_hash, first_name, last_name)
	VALUES
	(?, ?, ?, ?)`

	res, err := tools.DB.Exec(
		insertQuery,
		email,
		passwordHash,
		firstName,
		lastName,
	)

	if err != nil {
		err, code := handleInsertUserError(err)
		return -1, err, code
	}

	userId, err := res.LastInsertId()

	if err != nil {
		err, code := handleInsertUserError(err)
		return -1, err, code
	}

	return int(userId), nil, 0
}


func GetUser(email string) (resources.User, error, int) {
	var selectQuery string = `
	SELECT 
	id,
	email,
	password_hash,
	first_name,
	last_name
	FROM users
	WHERE email=?`

	row := tools.DB.QueryRow(selectQuery, email)

	var user resources.User
	switch err := row.Scan(
		&user.Id,
		&user.Email,
		&user.PasswordHash,
		&user.FirstName,
		&user.LastName,
	); err {
	case sql.ErrNoRows:
		return user, errors.New(`Invalid email or password. Please try again.`), 401
	case nil:
		return user, nil, 0
	default:
		return user, errors.New(`Something went wrong on our end. Please try again later`), 500
	}
}
