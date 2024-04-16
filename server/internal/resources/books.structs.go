package resources

import (
	"time"
)

type BookData struct {
	Title string
	Author string
	Published time.Time
	NumPages int32
	CoverURL string
}
