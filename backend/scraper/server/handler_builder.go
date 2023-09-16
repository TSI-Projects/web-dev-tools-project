package server

type Request struct {
	Category     []string `json:"category"`
	Source       []string `json:"source"`
	PriceMax     uint32   `json:"price_max"`
	PriceMin     uint32   `json:"price_min"`
	SearchedItem string   `json:"searched_item"`
}
