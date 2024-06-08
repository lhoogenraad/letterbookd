package resources

type CreateReviewBody struct {
	BookId int `json:"bookId"`
	Rating int `json:"rating"`
	Description string `json:"description"`
}
