package resources

type CreateReviewBody struct {
	Rating int `json:"rating"`
	Description string `json:"description"`
}

type UpdateReviewBody struct {
	Rating int `json:"rating"`
	Description string `json:"description"`
}
