package resources

type OpenLibraryEditionResponse struct {
	NumFound int `json:"numFound"`
	Docs []OpenLibraryEdition `json:"docs"`
}

type OpenLibraryEdition struct {
	Author_Name []string `json:"author_name"`
	AuthorKey []string `json:"author_key"`
	Title string `json:"title"`
	PublishDate []string `json:"publish_date"`
	EditionKey []string `json:"edition_key"`
	NumberOfPages int `json:"number_of_pages"`
	CoverEditionKey string `json:"cover_edition_key"`
	OlID string `json:"key"`
}
