package server

import (
	"net/http"

	"github.com/AndrejsPon00/web-dev-tools/backend/module"
	"github.com/AndrejsPon00/web-dev-tools/backend/scrapper"
)

type IHandler interface {
	SetSearchedProduct(searchedProduct string, ppCurentPage uint8)
	SetFilter(f *module.Filter)
	SetWriter(w http.ResponseWriter)
	AddWaitGroup(amount int)
	GetScraper() *scrapper.Client
	SetupErrorChannel()
	SetupResultChannel()
	Wait()
	Clear()
}
