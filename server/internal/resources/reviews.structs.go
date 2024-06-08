package resources

type CreateReviewBody struct {
	Rating int `json:"rating"`
	Description string `json:"description"`
}
