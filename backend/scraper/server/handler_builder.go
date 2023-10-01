package server

import (
	"context"
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
	ErrorChan    chan error
	ResultChan   chan *module.PreviewPost
	Scraper      *scrapper.Client
	WaitGroup    *sync.WaitGroup
	Mutex        *sync.Mutex
	TimeoutTimer *time.Timer
	Writer       http.ResponseWriter
}

func NewHandler() IHandler {
	resultChan := make(chan *module.PreviewPost)
	errorChan := make(chan error)
	ctx, cancel := context.WithCancel(context.Background())
	return &Handler{
		ErrorChan:    errorChan,
		ResultChan:   resultChan,
		TimeoutTimer: time.NewTimer(CONNECTION_TIMEOUT),
		WaitGroup:    &sync.WaitGroup{},
		Mutex:        &sync.Mutex{},
		Scraper: &scrapper.Client{
			Done:         cancel,
			Context:      ctx,
			WG:           &sync.WaitGroup{},
			Filter:       &module.Filter{},
			ResultChan:   resultChan,
			ErrorChan:    errorChan,
			SearchedItem: "",
		},
	}
}

func (h *Handler) SetSearchedProduct(searchedProduct string, ppCurentPage uint8) {
	h.Scraper.SearchedItem = searchedProduct
	h.Scraper.PPCurentPage = ppCurentPage
}

func (h *Handler) SetWriter(w http.ResponseWriter) {
	h.Writer = w
}

func (h *Handler) SetFilter(f *module.Filter) {
	h.Scraper.Filter = f
}

func (h *Handler) SetupErrorChannel() {
	defer h.WaitGroup.Done()

	for {
		select {
		case err := <-h.ErrorChan:
			h.Mutex.Lock()
			defer h.Mutex.Unlock()
			log.Errorln(err)
			http.Error(h.Writer, err.Error(), http.StatusInternalServerError)
			return
		case <-h.Scraper.Context.Done():
			return
		}
	}
}

func (h *Handler) SetupResultChannel() {
	defer h.WaitGroup.Done()
	buffer := make([]*module.PreviewPost, 0, POSTS_BUFFER_CAPACITY)

	for {
		select {
		case result, ok := <-h.ResultChan:
			h.Mutex.Lock()
			if !ok {
				if len(buffer) > 0 {
					h.SendResponse(buffer)
				}
				h.Mutex.Unlock()
				return
			}

			if len(buffer) < POSTS_BUFFER_CAPACITY {
				buffer = append(buffer, result)
			} else {
				h.SendResponse(buffer)
				buffer = buffer[:0]
			}
			h.Mutex.Unlock()

		case <-h.Scraper.Context.Done():
			h.SendResponse(buffer)
			return

		case <-h.TimeoutTimer.C:
			h.ErrorChan <- fmt.Errorf("request timed out. Please try again later")
			log.Errorln("Timed out")
			return
		}
	}
}

func (h *Handler) SendResponse(response interface{}) {
	posts, err := toByteArray(response)
	if err != nil {
		h.ErrorChan <- fmt.Errorf("services not responding. Please try again later")
		log.Errorf("Failed to unmarshal, Result: %v: %v", response, err)
	}

	if _, err := h.Writer.Write(posts); err != nil {
		h.ErrorChan <- fmt.Errorf("failed to send response: %v", err)
	}

	if flusher, ok := h.Writer.(http.Flusher); ok {
		flusher.Flush()
		log.Debugln("Response flushed")
	}
}

func (h *Handler) GetScraper() *scrapper.Client {
	return h.Scraper
}

func (h *Handler) AddWaitGroup(amount int) {
	h.WaitGroup.Add(amount)
}

func (h *Handler) Wait() {
	h.WaitGroup.Wait()
	h.Scraper.WG.Wait()
}

func (h *Handler) Clear() {
	close(h.ErrorChan)
	close(h.ResultChan)
	h.TimeoutTimer.Stop()
	h.Writer = nil
	h.Scraper = &scrapper.Client{}
	h.WaitGroup = &sync.WaitGroup{}
	h.Mutex = &sync.Mutex{}
}

func toByteArray(any interface{}) ([]byte, error) {
	output, err := json.Marshal(any)
	if err != nil {
		return nil, err
	}
	return output, nil
}
