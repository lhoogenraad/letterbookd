package structs

type Book struct{
	Languages []struct  {
		Key string `json: "key"`
	}
	Title string `json:"title"`
	Number_of_pages int16 `json:"number_of_pages`
	Publish_date string `json:"publish_date"`
	Authors []struct {
		Key string `json: "key"`
	}
	Description struct {
		Key string `json: "key"`
		Value string `json: "value"`
	}
	Key string `json: "key"`
	Covers []int `json: "covers"`
}
