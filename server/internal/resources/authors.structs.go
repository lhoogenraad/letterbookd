package resources

import "time"

type Author struct {
	FirstName string
	LastName string
	DateOfBirth time.Time
}

type AuthorOL struct {
	Name string `json:"name"`
	Birth_Date string `json:"birth_date"`
}
