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

type Error struct {
	Message string `json:"message"`
	Code    uint8  `json:"code"`
}

type Response struct {
	*Error       `json:"error,omitempty"`
	*PreviewPost `json:"preview_post,omitempty"`
	*Post        `json:"post,omitempty"`
}

type Request struct {
	Category     []string `json:"category"`
	Source       []string `json:"source"`
	PriceMax     uint32   `json:"price_max"`
	PriceMin     uint32   `json:"price_min"`
	SearchedItem string   `json:"searched_item"`
}
