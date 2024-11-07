package resources

import (
	"time"
)

type BookData struct {
	Id int
	Title string
	Author string
	AuthorId int
	Published time.Time
	NumPages int32
	CoverURL string
	Synopsis string
	OnUserReadlist bool
}

type BookDataOL struct {

	Id int
	Title string
	Author string
	AuthorOLId string
	Published time.Time
	NumPages int32
	CoverURL string
	Synopsis string
	OnUserReadlist bool
	OpenLibraryKey string
	CoverEdition string
	OlID string
}
