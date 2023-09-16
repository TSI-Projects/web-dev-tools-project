package scrapper

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/goccy/go-json"

	"github.com/AndrejsPon00/web-dev-tools/backend/module"
	"github.com/AndrejsPon00/web-dev-tools/backend/scrapper/ss.lv"
	"github.com/gocolly/colly/v2"
)

type Client struct {
	ErrorChan    chan error
	ResultChan   chan *module.PreviewPost
	WG           *sync.WaitGroup
	Collector    *colly.Collector
	TimeoutTimer *time.Timer
	Writer       http.ResponseWriter
	SearchedItem string
}

func NewScraper(searchedItem string, writer http.ResponseWriter) *Client {
	return &Client{
		SearchedItem: searchedItem,
		Writer:       writer,
		TimeoutTimer: time.NewTimer(time.Second * 5),
		ErrorChan:    make(chan error),
		ResultChan:   make(chan *module.PreviewPost),
		WG:           &sync.WaitGroup{},
		Collector: colly.NewCollector(
			colly.MaxDepth(2),
			colly.Async(),
		),
	}
}

func (c *Client) ScrapPosts() []*module.PreviewPost {
	posts := make([]*module.PreviewPost, 0)

	c.WG.Add(1)
	go ss.ScrapPosts(c.SearchedItem, c.WG, c.Collector, c.ResultChan)
	go func() {
		c.WG.Wait()
		close(c.ResultChan)
	}()

	for {
		select {
		case result, ok := <-c.ResultChan:
			if !ok {
				return posts
			}

			post, err := toByteArray(result)
			if err != nil {
				c.ErrorChan <- err
			}

			c.Writer.Write(post)
			if flusher, ok := c.Writer.(http.Flusher); ok {
				flusher.Flush()
			}

			log.Println(result)
		case error := <-c.ErrorChan:
			log.Fatalln(error.Error())
		case <-c.TimeoutTimer.C:
			log.Println("Timeout")
		}
	}
}

func toByteArray(any interface{}) ([]byte, error) {
	output, err := json.Marshal(any)
	if err != nil {
		return nil, err
	}
	return output, nil
}
