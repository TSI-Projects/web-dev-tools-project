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
	CONNECTION_TIMEOUT    = 20 * time.Second
	POSTS_BUFFER_CAPACITY = 25
)

type ResponseType string

const (
	Close      ResponseType = "close"
	Posts      ResponseType = "posts"
	Pagination ResponseType = "pagination"
)

type Handler struct {
	ErrorChan      chan error
	ResultChan     chan *module.PreviewPost
	PaginationChan chan *module.Pagination
	Scraper        *scrapper.Client
	WaitGroup      *sync.WaitGroup
	Mutex          *sync.Mutex
	TimeoutTimer   *time.Timer
	Writer         http.ResponseWriter
}

func NewHandler() IHandler {
	paginationChan := make(chan *module.Pagination)
	resultChan := make(chan *module.PreviewPost)
	errorChan := make(chan error)
	ctx, cancel := context.WithCancel(context.Background())
	return &Handler{
		PaginationChan: paginationChan,
		ErrorChan:      errorChan,
		ResultChan:     resultChan,
		TimeoutTimer:   time.NewTimer(CONNECTION_TIMEOUT),
		WaitGroup:      &sync.WaitGroup{},
		Mutex:          &sync.Mutex{},
		Scraper: &scrapper.Client{
			Done:           cancel,
			Context:        ctx,
			WG:             &sync.WaitGroup{},
			ResultChan:     resultChan,
			ErrorChan:      errorChan,
			PaginationChan: paginationChan,
		},
	}
}

func (h *Handler) SetParams(params *module.URLParams) {
	h.Scraper.Params = params
}

func (h *Handler) SetWriter(w http.ResponseWriter) {
	h.Writer = w
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
					h.SendResponse(buffer, Posts)
				}
				h.Mutex.Unlock()
				return
			}

			if len(buffer) < POSTS_BUFFER_CAPACITY {
				buffer = append(buffer, result)
			} else {
				h.SendResponse(buffer, Posts)
				buffer = buffer[:0]
			}
			h.Mutex.Unlock()

		case pagination, ok := <-h.PaginationChan:
			if !ok {
				return
			}
			h.Mutex.Lock()
			h.SendResponse(pagination, Pagination)
			h.Mutex.Unlock()

		case <-h.Scraper.Context.Done():
			h.SendResponse(buffer, Posts)
			return

		case <-h.TimeoutTimer.C:
			h.ErrorChan <- fmt.Errorf("request timed out. Please try again later")
			log.Errorln("Timed out")
			return
		}
	}
}

func (h *Handler) SendResponse(response interface{}, resType ResponseType) {
	posts, err := toByteArray(response)
	if err != nil {
		h.ErrorChan <- fmt.Errorf("services not responding. Please try again later")
		log.Errorf("Failed to unmarshal, Result: %v: %v", response, err)
	}

	if _, err := fmt.Fprintf(h.Writer, "event: %s\ndata: %v\n\n", resType, string(posts)); err != nil {
		h.ErrorChan <- fmt.Errorf("failed to send response: %v", err)
	}

	h.Writer.(http.Flusher).Flush()
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
	h.CloseSSEConnection()
	close(h.ErrorChan)
	close(h.ResultChan)
	h.TimeoutTimer.Stop()
	h.Writer = nil
	h.Scraper = &scrapper.Client{}
	h.WaitGroup = &sync.WaitGroup{}
	h.Mutex = &sync.Mutex{}
}

func (h *Handler) CloseSSEConnection() {
	h.SendResponse("Connection closed", Close)
}

func toByteArray(any interface{}) ([]byte, error) {
	output, err := json.Marshal(any)
	if err != nil {
		return nil, err
	}
	return output, nil
}
