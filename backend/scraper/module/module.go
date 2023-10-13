package module

const (
	SOURCE_SS       = "ss"
	SOURCE_PP       = "pp"
	SOURCE_FACEBOOK = "facebook"
	SOURCE_GELIOS   = "gelios"
	SOURCE_BANKNOTE = "banknote"
)

const (
	MAX_UINT32_SIZE = 4294967295
)

type Pagination struct {
	Source  string `json:"source"`
	HasNext bool   `json:"has_next"`
}

type PreviewPost struct {
	Title        string `json:"title"`
	PreviewImage string `json:"preview_img"`
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
}

type URLParams struct {
	*Filter
	SearchedItem        string `schema:"query"`
	PPCurrentPage       uint8  `schema:"pp_page"`
	SSCurrentPage       uint8  `schema:"ss_page"`
	BanknoteCurrentPage uint8  `schema:"banknote_page"`
}

type Filter struct {
	Sources  []string `schema:"sources,omitempty"`
	Category []string `schema:"category,omitempty"`
	PriceMax uint32   `schema:"price_max,omitempty"`
	PriceMin uint32   `schema:"price_min,omitempty"`
}
