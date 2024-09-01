package resources

type ReviewComment struct {
	Id int
	Comment string
	UserId int
	ReviewId int
}

type CreateReviewCommentBody struct {
	Comment string
}
