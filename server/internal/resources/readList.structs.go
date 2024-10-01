package resources

type ReadListItem struct {
	Id int
	UserId int
	BookId int
	BookName string
	Status string
}

type AddBookToReadlistRequest struct {
	Status string
}
