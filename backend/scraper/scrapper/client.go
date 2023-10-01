package scrapper

import (
	"context"
	"sync"

	"github.com/AndrejsPon00/web-dev-tools/backend/module"
	"github.com/AndrejsPon00/web-dev-tools/backend/scrapper/pp.lv"
	"github.com/AndrejsPon00/web-dev-tools/backend/scrapper/ss.lv"
	"github.com/gocolly/colly/v2"
	log "github.com/sirupsen/logrus"
)

type Client struct {
	ResultChan   chan *module.PreviewPost
	ErrorChan    chan error
	WG           *sync.WaitGroup
	Filter       *module.Filter
	Done         context.CancelFunc
	Context      context.Context
	SearchedItem string
	PPCurentPage uint8
}

func (c *Client) ScrapPosts() {
	for _, source := range c.Filter.Sources {
		c.WG.Add(1)
		collector := colly.NewCollector()
		switch source {
		case module.SOURCE_SS_LV:
			go ss.ScrapPosts(c.SearchedItem, c.WG, collector, c.ResultChan, c.ErrorChan)
		case module.SOURCE_FACEBOOK:
			//add scrap facebook
		case module.SOURCE_GELIOS:
			//add scrap gelios
		case module.SOURCE_PP_LV:
			go pp.ScrapPosts(c.SearchedItem, c.PPCurentPage, c.WG, c.ResultChan, c.ErrorChan)
		default:
			log.Errorln("Unknown source: ", source)
			c.WG.Done()
		}
	}

	c.WG.Done()
	c.WG.Wait()
	c.Done()
}
