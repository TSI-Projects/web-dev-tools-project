package ss

import (
	"fmt"
	"strconv"
	"strings"
	"sync"

	"github.com/AndrejsPon00/web-dev-tools/backend/module"
	"github.com/gocolly/colly/v2"
	log "github.com/sirupsen/logrus"
)

const (
	BASE_SS_URL = "https://www.ss.lv/ru/search-result"
)

func ScrapPosts(input string, currentPage uint8, wg *sync.WaitGroup, c *colly.Collector, paginationChan chan *module.Pagination, result chan *module.PreviewPost, errorChan chan error) {
	defer wg.Done()
	encodedQuery := encodeStringToHTML(input)
	completeURL := combineURL(BASE_SS_URL, encodedQuery, currentPage)
	c.OnHTML("tr:has(td.msga2):has(td.msg2):has(td.msga2-o.pp6)", func(e *colly.HTMLElement) {
		url := fmt.Sprintf("%s%s", BASE_SS_URL, e.ChildAttr("a", "href"))
		previewImage := e.ChildAttr("img", "src")
		title := strings.TrimSpace(e.ChildText("a.am"))
		description := strings.ReplaceAll(e.ChildText("div.d1"), "\n", " ")
		price := strings.TrimSpace(e.ChildText("td.msga2-o.pp6"))

		if len(description) > 150 {
			return
		}

		result <- &module.PreviewPost{
			URL:          url,
			PreviewImage: previewImage,
			Title:        title,
			Description:  description,
			Price:        price,
		}
	})

	c.OnHTML(".td2", func(e *colly.HTMLElement) {
		paginationChan <- &module.Pagination{
			Source:  module.SOURCE_SS_LV,
			HasNext: hasNextPage(currentPage, e),
		}
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.OnError(func(response *colly.Response, err error) {
		log.Errorf("Error scraping. With response: %v Error: %v. ", response, err)
		errorChan <- err
	})

	c.Visit(completeURL)
	c.Wait()
}

func hasNextPage(currentPage uint8, e *colly.HTMLElement) bool {
	hasNextPage := false
	for _, page := range e.ChildTexts("a[rel='next']") {
		uintPage, err := strconv.Atoi(page)
		if err != nil {
			continue
		}

		if uint8(uintPage) > currentPage {
			hasNextPage = true
		}
	}
	return hasNextPage
}

func combineURL(baseURL, query string, currentPage uint8) string {
	return fmt.Sprintf("%s/page%d.html?q=%s", baseURL, currentPage, query)
}

func encodeStringToHTML(query string) string {
	return strings.ReplaceAll(query, " ", "+")
}
