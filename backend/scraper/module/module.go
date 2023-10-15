package module

const (
	SOURCE_SS       Source = "ss"
	SOURCE_PP       Source = "pp"
	SOURCE_FACEBOOK Source = "facebook"
	SOURCE_GELIOS   Source = "gelios"
	SOURCE_BANKNOTE Source = "banknote"
)

const (
	ENV_VAR_PORT = "SERVER_PORT"
	DEFAULT_PORT = ":8080"
)

const (
	MAX_UINT32_SIZE = 4294967295
)

var EVERY_SOURCE = []Source{
	SOURCE_SS,
	SOURCE_PP,
	SOURCE_GELIOS,
	SOURCE_BANKNOTE,
	SOURCE_FACEBOOK,
}

type Source string

type Pagination struct {
	Source  `json:"source"`
	HasNext bool `json:"has_next"`
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
	Sources  []Source `schema:"sources,omitempty"`
	Category []string `schema:"category,omitempty"`
	PriceMax uint32   `schema:"price_max,omitempty"`
	PriceMin uint32   `schema:"price_min,omitempty"`
}
