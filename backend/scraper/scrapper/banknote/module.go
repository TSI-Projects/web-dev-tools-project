package banknote

type Response struct {
	CurrentPage  int     `json:"current_page"`
	Items        []*Item `json:"data"`
	FirstPageURL string  `json:"first_page_url"`
	From         int     `json:"from"`
	LastPage     int     `json:"last_page"`
	LastPageURL  string  `json:"last_page_url"`
	NextPageURL  string  `json:"next_page_url"`
	Path         string  `json:"path"`
	PerPage      int     `json:"per_page"`
	PrevPageURL  string  `json:"prev_page_url"`
	To           int     `json:"to"`
	Total        int     `json:"total"`
}

type Item struct {
	Title       string `json:"title"`
	Price       string `json:"actual_price"`
	Image       string `json:"img"`
	RedirectURL string `json:"url"`
}
