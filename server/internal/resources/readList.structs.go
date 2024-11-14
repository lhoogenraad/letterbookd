package resources

type ReadListItem struct {
	Id int
	UserId int
	BookId int
	BookName string
	BookSynopsis string
	Status string
	UserReviewed bool
	CoverURL string
}

type ReadListModReq struct {
	Status string
}
