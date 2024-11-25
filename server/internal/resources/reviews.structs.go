package resources

/* 
This is the struct of review data that is returned by the API. It
It includes joined user info.
*/
type ReviewData struct {
	Id int
	Description string
	Rating int
	Username string
	UserId int
	OwnedBy bool
	NumComments int
	NumLikes int
	LikedBy bool
	BookId int
}

type CreateReviewBody struct {
	Rating int `json:"rating"`
	Description string `json:"description"`
}

type UpdateReviewBody struct {
	Rating int `json:"rating"`
	Description string `json:"description"`
}
