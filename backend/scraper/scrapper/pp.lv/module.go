package pp

type Response struct {
	Content *Content `json:"content"`
}

type Content struct {
	ItemsCount uint16  `json:"count"`
	Data       []*Data `json:"data"`
}

type Data struct {
	RedirectURL string      `json:"frontUrl"`
	Title       string      `json:"title"`
	FileCount   uint8       `json:"fileCount"`
	Prices      []*Price    `json:"prices"`
	Files       []*File     `json:"files"`
	AdFilters   []*AdFilter `json:"adFilterValues"`
}

type AdFilter struct {
	Filter *Filter `json:"filter"`
}

type Filter struct {
	ID int `json:"id"`
}

type Price struct {
	Value string `json:"value"`
}

type Category struct {
	Name   string    `json:"name"`
	Parent *Category `json:"parent,omitempty"`
}

type File struct {
	File *Thumbnail `json:"file"`
}

type Thumbnail struct {
	ID       int      `json:"id"`
	Versions *Version `json:"versions"`
}

type FileInfo struct {
	Path string `json:"path"`
}

type Version struct {
	OriginalFile *FileInfo `json:"original"`
}
