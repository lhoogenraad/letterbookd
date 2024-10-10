package resources

type ReviewComment struct {
	Id int
	Comment string
	UserId int
	Username string
	ReviewId int
	OwnedBy bool
	Edited bool
}

type CreateReviewCommentBody struct {
	Comment string
}
