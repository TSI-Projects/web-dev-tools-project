package module

const (
	SOURCE_SS_LV    = "ss"
	SOURCE_PP_LV    = "pp"
	SOURCE_FACEBOOK = "facebook"
	SOURCE_GELIOS   = "gelios"
	SOURCE_BANKNOTE = "banknote"
)

type Pagination struct {
	Source  string `json:"source"`
	HasNext bool   `json:"has_next"`
}

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

type URLParams struct {
	*Filter
	SearchedItem        string `schema:"product"`
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
