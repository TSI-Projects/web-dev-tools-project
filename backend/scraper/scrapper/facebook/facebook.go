package facebook

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/AndrejsPon00/web-dev-tools/backend/module"
	log "github.com/sirupsen/logrus"
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

const (
	seleniumPath = "http://localhost:4444/wd/hub"
	browser      = "chrome"

	BASE_URL_PATH     = "https://www.facebook.com/marketplace/110934415597150/search?query="
	POSTS_IN_ONE_PAGE = 15
)

func ScrapPosts(input string, pageNumber uint8, wg *sync.WaitGroup, result chan *module.PreviewPost, paginationChan chan *module.Pagination, errorChan chan error) {
	wd, err := newDriver()
	if err != nil {
		errorChan <- fmt.Errorf("error creating driver: %v", err)
		return
	}
	defer wd.Quit()

	fullURL := buildFullUrl(input)
	if err := wd.Get(fullURL); err != nil {
		errorChan <- fmt.Errorf("error connectiong to facebook: %v", err)
		return
	}

	if err := acceptCookies(wd); err != nil {
		errorChan <- fmt.Errorf("error accepting cookies: %v", err)
		return
	}

	time.Sleep(500 * time.Millisecond)

	if err := injectScrollScript(pageNumber, wd); err != nil {
		errorChan <- fmt.Errorf("error injecting scroll script: %v", err)
		return
	}

	pageNumber = getCorrectPagePath(pageNumber)
	getAndSendPosts(pageNumber, wd, result)
	handlePagination(pageNumber, wd, paginationChan)
	wg.Done()
}

func getAndSendPosts(pageNumber uint8, driver selenium.WebDriver, resultChan chan *module.PreviewPost) {
	for i := 1; i < POSTS_IN_ONE_PAGE; i++ {
		path := buildElementPath(pageNumber, i)
		post, err := driver.FindElement(selenium.ByXPATH, path)
		if err != nil {
			log.Warning("Post not found")
			continue
		}

		price, err := getText(PRICE_PATH, driver, post)
		if err != nil {
			log.Warning(err)
			continue
		}

		title, err := getText(TITLE_PATH, driver, post)
		if err != nil {
			log.Warning(err)
			continue
		}

		imgLink, err := getAttribute(IMAGE_PATH, SRC_ATTRIBUTE, driver, post)
		if err != nil {
			log.Warning(err)
			continue
		}

		url, err := getAttribute(URL_PATH, HREF_ATTRIBUTE, driver, post)
		if err != nil {
			log.Warning(err)
			continue
		}

		resultChan <- &module.PreviewPost{
			URL:          url,
			PreviewImage: imgLink,
			Price:        price,
			Title:        title,
		}
	}

}

func handlePagination(pageNumber uint8, driver selenium.WebDriver, paginationChan chan *module.Pagination) {
	pageNumber += 2
	elementPath := buildElementPath(pageNumber, 1)
	post, err := driver.FindElement(selenium.ByXPATH, elementPath)
	if err != nil {
		log.Warning("Post not found")
		writePagination(false, paginationChan)
		return
	}

	_, err = getText(PRICE_PATH, driver, post)
	if err != nil {
		writePagination(false, paginationChan)
		return
	}

	writePagination(true, paginationChan)
}

func writePagination(HasNext bool, paginationChan chan *module.Pagination) {
	paginationChan <- &module.Pagination{
		Source:  module.SOURCE_FACEBOOK,
		HasNext: HasNext,
	}
}

func getAttribute(xPath string, attribute string, driver selenium.WebDriver, parentElement selenium.WebElement) (string, error) {
	element, err := parentElement.FindElement(selenium.ByXPATH, xPath)
	if err != nil {
		return "", fmt.Errorf("failed to find element")
	}
	attributeValue, err := element.GetAttribute(attribute)
	if err != nil {
		return "", fmt.Errorf("failed to find attribute in element (%s)", attribute)
	}

	return attributeValue, nil
}

func getText(xPath string, driver selenium.WebDriver, parentElement selenium.WebElement) (string, error) {
	element, err := parentElement.FindElement(selenium.ByXPATH, xPath)
	if err != nil {
		return "", fmt.Errorf("failed to find element")
	}
	text, err := element.Text()
	if err != nil {
		return "", fmt.Errorf("failed to find text in element")
	}
	return text, nil
}

func buildFullUrl(input string) string {
	return fmt.Sprintf("%s%s", BASE_URL_PATH, input)
}

func buildElementPath(pageNumber uint8, itemNumber int) string {
	return fmt.Sprintf(`/html/body/div[1]/div/div[1]/div/div[3]/div/div/div/div[1]/div[1]/div[2]/div/div/div[3]/div[%d]/div[2]/div[%d]`, pageNumber, itemNumber)
}

func newDriver() (selenium.WebDriver, error) {
	caps := selenium.Capabilities{}
	caps.AddChrome(chrome.Capabilities{Args: []string{
		"window-size=1920x1080",
		"--no-sandbox",
		"--disable-dev-shm-usage",
		"disable-gpu",
		"--headless",
	}})

	return selenium.NewRemote(caps, seleniumPath)
}

func getCorrectPagePath(pageNumber uint8) uint8 {
	if isEvenNumber(pageNumber) {
		pageNumber++
	}
	return pageNumber
}

func isEvenNumber(num uint8) bool {
	return num%2 == 0
}

func injectScrollScript(pageNumber uint8, wd selenium.WebDriver) error {
	scrollingScript := fmt.Sprintf(`
	const scrolls = %d
	let scrollCount = 0
  
	const scrollInterval = setInterval(() => {
	 window.scrollTo(0, document.body.scrollHeight)
	 scrollCount++
  
	 if (scrollCount === scrolls) {
	 clearInterval(scrollInterval)
	 }
	}, 500)
	`, pageNumber)
	if _, err := wd.ExecuteScript(scrollingScript, []interface{}{}); err != nil {
		return err
	}
	sleepDuration := time.Duration(int(pageNumber)*600) * time.Millisecond
	time.Sleep(sleepDuration)
	return nil
}

func acceptCookies(wd selenium.WebDriver) error {
	button, err := waitForElement(COOKIE_ACCEPT_BUTTON_PATH, wd, 3*time.Second)
	if err != nil {
		return err
	}
	button.Click()
	return nil
}

func waitForElement(xPath string, driver selenium.WebDriver, timeout time.Duration) (selenium.WebElement, error) {
	startTime := time.Now()
	var element selenium.WebElement
	for {
		if time.Since(startTime) > timeout {
			return nil, errors.New("timeout while waiting for the element to be rendered")
		}

		var err error
		element, err = driver.FindElement(selenium.ByXPATH, xPath)
		if err != nil || element == nil {
			time.Sleep(200 * time.Millisecond)
			continue
		}

		displayed, err := element.IsDisplayed()
		if err != nil {
			return nil, err
		}

		if displayed {
			return element, nil
		}

		time.Sleep(100 * time.Millisecond)
	}
}
