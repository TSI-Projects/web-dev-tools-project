package selenium

import (
	"os"

	"github.com/tebeka/selenium"
)

func StartServer(exitChan chan os.Signal) {
	service, err := selenium.NewChromeDriverService("./chrome_driver/chromedriver_mac", 4444)
	if err != nil {
		panic(err)
	}
	defer service.Stop()
	<-exitChan
}
