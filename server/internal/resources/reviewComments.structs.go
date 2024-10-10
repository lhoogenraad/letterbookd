package resources

import (
	"time"
)

type ReviewComment struct {
	Id int
	Comment string
	UserId int
	Username string
	ReviewId int
	OwnedBy bool
	Edited bool
	Timestamp time.Time
}

type CreateReviewCommentBody struct {
	Comment string
}
