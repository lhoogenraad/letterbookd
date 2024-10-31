package authors

import (
	"fmt"
	"scripts/util"
	"strings"
	"time"
	"scripts/structs"
)

func UploadAuthor(author structs.Author) error {
	nameSplit := strings.Split(author.Name, " ")
	var firstName string
	var lastName string
	if len(nameSplit) < 2 {
		firstName = author.Name
		lastName = ""
	} else {
		firstName = nameSplit[0]
		lastName = strings.Join(nameSplit[1:], " ")
	}

	var insertQuery string = `
		INSERT INTO authors
		(first_name, last_name, date_of_birth, ol_id)
		VALUES
		(?, ?, ?, ?);
	`
	fmt.Println(author.DOB.Format(time.RFC3339), "\n", author.Key, "\n\n")
	_, err := util.DB.Exec(
		insertQuery, 
		firstName, 
		lastName,
		author.DOB.Format("2006-01-02"),
		author.Key,
	)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
