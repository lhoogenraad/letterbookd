package structs

import "time"

type Author struct {
	Key string `json: "key"`
	Name string `json: "name"`
	Birth_date string `json: "birth_date"`
	DOB time.Time
}
