package resources

type OpenLibraryEditionResponse struct {
	NumFound int `json:"numFound"`
	Docs []OpenLibraryEdition `json:"docs"`
}

type OpenLibraryEdition struct {
	Author_Name []string `json:"author_name"`
	Title string `json:"title"`
	PublishDate []string `json:"publish_date"`
}
