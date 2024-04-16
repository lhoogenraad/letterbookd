package models

import (
	"time"
)

type BookData struct {
	Title string
	Author string
	Published time.Time
}
