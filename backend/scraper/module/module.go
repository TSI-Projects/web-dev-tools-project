package module

type Post struct {
	Description string   `json:"description"`
	Price       string   `json:"price"`
	Id          string   `json:"id"`
	Imgs        []string `json:"imgs"`
}

type PreviewPost struct {
	Id           string `json:"id"`
	Title        string `json:"title"`
	PreviewImage string `json:"preview_img"`
	Description  string `json:"description"`
	Price        string `json:"price"`
	URL          string `json:"url"`
}
