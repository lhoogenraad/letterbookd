package resources

import (
	"time"
)

type BookData struct {
	Id int
	Title string
	Author string
	Published time.Time
	NumPages int32
	CoverURL string
	Synopsis string
	OnUserReadlist bool
}
