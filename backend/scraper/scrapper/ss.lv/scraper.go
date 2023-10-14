package ss

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
	"unicode"

	"github.com/AndrejsPon00/web-dev-tools/backend/module"
	"github.com/gocolly/colly/v2"
	log "github.com/sirupsen/logrus"
)

const (
	BASE_SS_URL = "https://www.ss.lv/ru/search-result"
)

func ScrapPosts(input string, currentPage uint8, filter *module.Filter, wg *sync.WaitGroup, c *colly.Collector, paginationChan chan *module.Pagination, result chan *module.PreviewPost, errorChan chan error) {
	defer wg.Done()
	isPaginationSend := false
	encodedQuery := encodeStringToHTML(input)
	completeURL := combineURL(BASE_SS_URL, encodedQuery, currentPage)
	c.OnHTML("tr:has(td.msga2):has(td.msg2):has(td.msga2-o.pp6)", func(e *colly.HTMLElement) {
		url := fmt.Sprintf("%s%s", "https://www.ss.lv", e.ChildAttr("a", "href"))
		previewImage := e.ChildAttr("img", "src")
		title := strings.TrimSpace(e.ChildText("a.am"))
		description := strings.ReplaceAll(e.ChildText("div.d1"), "\n", " ")
		price := strings.TrimSpace(e.ChildText("td.msga2-o.pp6"))

		if len(description) > 150 {
			return
		}

		post := &module.PreviewPost{
			URL:          url,
			PreviewImage: previewImage,
			Title:        title,
			Price:        price,
		}

		response := filterPreviewPostByPrice(post, filter)
		if response != nil {
			result <- response
		}
	})

	c.OnHTML(".td2", func(e *colly.HTMLElement) {
		sendPagination(currentPage, e, paginationChan)
		isPaginationSend = true
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

	if !isPaginationSend {
		sendPagination(currentPage, nil, paginationChan)
	}
}

func filterPreviewPostByPrice(response *module.PreviewPost, filter *module.Filter) *module.PreviewPost {
	if filter == nil {
		return nil
	}

	if filter.PriceMax == 0 {
		filter.PriceMax = module.MAX_UINT32_SIZE
	}

	if !isPriceAboveMinimum(filter.PriceMin, response.Price) {
		return nil
	}

	if !isPriceBelowMaximum(filter.PriceMax, response.Price) {
		return nil
	}

	return response
}

func isPriceBelowMaximum(maxPrice uint32, itemPrice string) bool {
	strPrice := removeNonNumericChar(itemPrice)
	price := strToUint32(strPrice)
	return maxPrice >= price
}

func isPriceAboveMinimum(minPrice uint32, itemPrice string) bool {
	strPrice := removeNonNumericChar(itemPrice)
	price := strToUint32(strPrice)
	return minPrice <= price
}

func removeNonNumericChar(str string) string {
	var numericRunes []rune
	for _, r := range str {
		if unicode.IsDigit(r) {
			numericRunes = append(numericRunes, r)
		}
	}
	return string(numericRunes)
}

func strToUint32(str string) uint32 {
	num, err := strconv.ParseUint(str, 10, 32)
	if err != nil {
		log.Errorf("Failed to convert string '%s' to uint32: %v", str, err)
		return 0
	}

	return uint32(num)
}

func sendPagination(currentPage uint8, e *colly.HTMLElement, channel chan *module.Pagination) {
	channel <- &module.Pagination{
		Source:  module.SOURCE_SS,
		HasNext: hasNextPage(currentPage, e),
	}
}

func hasNextPage(currentPage uint8, e *colly.HTMLElement) bool {
	if e == nil {
		return false
	}

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
