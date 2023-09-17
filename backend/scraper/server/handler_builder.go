package server

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/AndrejsPon00/web-dev-tools/backend/module"
	"github.com/AndrejsPon00/web-dev-tools/backend/scrapper"
	"github.com/goccy/go-json"
	log "github.com/sirupsen/logrus"
)

const (
	CONNECTION_TIMEOUT    = 5 * time.Second
	POSTS_BUFFER_CAPACITY = 25
)

type Handler struct {
	ErrorChan       chan error
	ResultChan      chan *module.PreviewPost
	Scraper         *scrapper.Client
	WaitGroup       *sync.WaitGroup
	Mutex           *sync.Mutex
	TimeoutTimer    *time.Timer
	Filter          *module.Filter
	Writer          http.ResponseWriter
	SearchedProduct string
}

func NewHandler() IHandler {
	return &Handler{
		ErrorChan:    make(chan error),
		ResultChan:   make(chan *module.PreviewPost),
		TimeoutTimer: time.NewTimer(CONNECTION_TIMEOUT),
		Scraper:      &scrapper.Client{},
		Filter:       &module.Filter{},
		WaitGroup:    &sync.WaitGroup{},
		Mutex:        &sync.Mutex{},
	}
}

func (h *Handler) SetSearchedProduct(searchedProduct string) {
	h.SearchedProduct = searchedProduct
}

func (h *Handler) SetWriter(w http.ResponseWriter) {
	h.Writer = w
}

func (h *Handler) SetFilter(f *module.Filter) {
	h.Filter = f
}

func (h *Handler) SetupErrorChannel() {
	defer h.WaitGroup.Done()

	err := <-h.ErrorChan
	h.Mutex.Lock()
	defer h.Mutex.Unlock()
	http.Error(h.Writer, err.Error(), http.StatusInternalServerError)
}

func (h *Handler) SetupResultChannel() {
	defer h.WaitGroup.Done()
	buffer := make([]*module.PreviewPost, 0, POSTS_BUFFER_CAPACITY)

	select {
	case result, ok := <-h.ResultChan:
		if !ok {
			return
		}

		if len(buffer) <= POSTS_BUFFER_CAPACITY {
			buffer = append(buffer, result)
		} else {
			posts, err := toByteArray(buffer)
			if err != nil {
				h.ErrorChan <- fmt.Errorf("services not responding. Please try again later")
				log.Errorf("Failed to unmarshal, Result: %v: %v", result, err)
			}

			h.Writer.Write(posts)
			if flusher, ok := h.Writer.(http.Flusher); ok {
				flusher.Flush()
			}
			buffer = buffer[:0]
		}

	case <-h.TimeoutTimer.C:
		h.ErrorChan <- fmt.Errorf("request timed out. Please try again later")
		log.Errorln("Timed out")
	}
}

func (h *Handler) AddWaitGroup(amount int) {
	h.WaitGroup.Add(amount)
}

func (h *Handler) Clear() {
	close(h.ErrorChan)
	close(h.ResultChan)
	h.TimeoutTimer.Stop()
	h.Writer = nil
	h.Scraper = &scrapper.Client{}
	h.WaitGroup = &sync.WaitGroup{}
	h.Mutex = &sync.Mutex{}
	h.Filter = &module.Filter{}
	h.SearchedProduct = ""
}

func toByteArray(any interface{}) ([]byte, error) {
	output, err := json.Marshal(any)
	if err != nil {
		return nil, err
	}
	return output, nil
}
