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
	Description struct {
		Value string `json:"value"`
	}
	Covers []int `json:"covers"`
	CoverEditionKey string `json:"cover_edition_key"`
	Cover_Edition struct {
		Key string `json:"key"`
	}
	OlID string `json:"key"`
	Authors []struct {
		Author struct {
			Key string `json:"key"`
		}
	}
}
