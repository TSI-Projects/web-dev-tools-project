package scrapper

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/AndrejsPon00/web-dev-tools/backend/module"
	"github.com/gocolly/colly/v2"
)

const (
	BASE_SS_URL       = "https://www.ss.lv"
	BASE_SEARCH_QUERY = "ru/search-result/?q="
)

func WebsiteScrapperSS(input string) []*module.PreviewPost {
	encodedQuery := encodeStringToHTML(input)
	completeURL := combineURL(BASE_SS_URL, encodedQuery)

	var posts []*module.PreviewPost

	c := colly.NewCollector()

	c.OnHTML("tr:has(td.msga2):has(td.msg2):has(td.msga2-o.pp6)", func(e *colly.HTMLElement) {
		url := fmt.Sprintf("%s%s", BASE_SS_URL, e.ChildAttr("a", "href"))
		previewImage := e.ChildAttr("img", "src")
		title := strings.TrimSpace(e.ChildText("a.am"))
		description := strings.ReplaceAll(e.ChildText("div.d1"), "\n", " ")
		price := strings.TrimSpace(e.ChildText("td.msga2-o.pp6"))

		if len(description) > 150 {
			return
		}

		posts = append(posts, &module.PreviewPost{
			URL:          url,
			PreviewImage: previewImage,
			Title:        title,
			Description:  description,
			Price:        price,
		})
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})
	c.Visit(completeURL)

	return posts
}

func combineURL(baseURL, query string) string {
	return fmt.Sprintf(baseURL + BASE_SEARCH_QUERY + query)
}

func takInput() (string, error) {
	var input string
	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {
		input = scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		return input, fmt.Errorf("failed to scan input: %v", err)
	}

	return input, nil
}

func encodeStringToHTML(query string) string {
	return strings.ReplaceAll(query, " ", "+")
}
