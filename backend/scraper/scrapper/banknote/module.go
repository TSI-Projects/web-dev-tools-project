package banknote

type Response struct {
	Items       []*Item `json:"data"`
	NextPageURL string  `json:"next_page_url,omitempty"`
}

type Item struct {
	Title       string `json:"title"`
	Price       string `json:"price"`
	Image       string `json:"img"`
	RedirectURL string `json:"url"`
}
