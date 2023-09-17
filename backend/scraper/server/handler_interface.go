package server

import (
	"net/http"

	"github.com/AndrejsPon00/web-dev-tools/backend/module"
)

type IHandler interface {
	SetSearchedProduct(searchedProduct string)
	SetFilter(f *module.Filter)
	SetWriter(w http.ResponseWriter)
	SetupErrorChannel()
	SetupResultChannel()
	AddWaitGroup(amount int)
	Clear()
}
