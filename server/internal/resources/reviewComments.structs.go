package resources

type ReviewComment struct {
	Id int
	Comment string
	UserId int
	Username string
	ReviewId int
}

type CreateReviewCommentBody struct {
	Comment string
}
