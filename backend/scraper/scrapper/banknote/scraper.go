package banknote

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"

	"github.com/AndrejsPon00/web-dev-tools/backend/module"
	"github.com/goccy/go-json"
)

const (
	BASE_URL          = "https://veikals.banknote.lv/lv/filter-products?"
	BASE_SEARCH_QUERY = "title="
	BASE_PAGE_QUERY   = "page="
	POSTS_IN_ONE_PAGE = 50
)

func ScrapPosts(input string, pageNumber uint8, wg *sync.WaitGroup, paginationChan chan *module.Pagination, result chan *module.PreviewPost, errorChan chan error) {
	defer wg.Done()

	url := getFullURL(input, pageNumber)
	rawResponse, err := FetchResponse(url)
	if err != nil {
		errorChan <- err
		return
	}

	response, err := DecodeResponse(rawResponse)
	if err != nil {
		errorChan <- err
		return
	}

	SendPaginationPostsToChannel(response, paginationChan)
	SendPreviewPostsToChannel(response, result)
}

func FetchResponse(input string) ([]byte, error) {
	resp, err := http.Get(input)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("out of bounds")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func DecodeResponse(response []byte) (*Response, error) {
	res := &Response{}
	err := json.Unmarshal(response, res)
	return res, err
}

func SendPreviewPostsToChannel(response *Response, resultChan chan *module.PreviewPost) {
	for _, item := range response.Items {
		resultChan <- &module.PreviewPost{
			Title:        item.Title,
			URL:          item.RedirectURL,
			PreviewImage: item.Image,
			Price:        item.Price,
		}
	}
}

func SendPaginationPostsToChannel(response *Response, paginationChan chan *module.Pagination) {
	paginationChan <- &module.Pagination{
		Source:  module.SOURCE_BANKNOTE,
		HasNext: hasHextPage(response),
	}
}
func hasHextPage(response *Response) bool {
	return response.NextPageURL != ""
}

func encodeSpacesForURL(query string) string {
	return strings.ReplaceAll(query, " ", "+")
}

func getFullURL(query string, pageNumber uint8) string {
	return fmt.Sprintf("%s%s%d&%s%s", BASE_URL, BASE_PAGE_QUERY, pageNumber, BASE_SEARCH_QUERY, encodeSpacesForURL(query))
}
