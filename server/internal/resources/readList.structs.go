package resources

type ReadListItem struct {
	Id int
	UserId int
	BookId int
	BookName string
	BookSynopsis string
	Status string
	UserReviewed bool
}

type ReadListModReq struct {
	Status string
}
