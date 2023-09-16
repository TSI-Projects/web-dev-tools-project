package ss

import (
	"fmt"
	"strings"
	"sync"

	"github.com/AndrejsPon00/web-dev-tools/backend/module"
	"github.com/gocolly/colly/v2"
)

const (
	BASE_SS_URL       = "https://www.ss.lv/"
	BASE_SEARCH_QUERY = "ru/search-result/?q="
	POSTS_IN_ONE_PAGE = 26
)

func ScrapPosts(input string, wg *sync.WaitGroup, c *colly.Collector, result chan *module.PreviewPost) {
	defer wg.Done()
	encodedQuery := encodeStringToHTML(input)
	completeURL := combineURL(BASE_SS_URL, encodedQuery)

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

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})
	c.Visit(completeURL)
	c.Wait()
}

func ScrapPost(url string, c *colly.Collector) *module.Post {
	post := &module.Post{}
	c.OnHTML("div.msg_div_msg)", func(e *colly.HTMLElement) {
		// post.Imgs := e.ChildAttr("img", "src")
		post.Description = strings.ReplaceAll(e.ChildText("div.d1"), "\n", " ")
		post.Price = strings.TrimSpace(e.ChildText("td.msga2-o.pp6"))
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})
	c.Visit(url)
	return post
}

func combineURL(baseURL, query string) string {
	return fmt.Sprintf(baseURL + BASE_SEARCH_QUERY + query)
}

func encodeStringToHTML(query string) string {
	return strings.ReplaceAll(query, " ", "+")
}
