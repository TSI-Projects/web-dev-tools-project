package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/AndrejsPon00/web-dev-tools/backend/scrapper"
)

func Start() {
	http.HandleFunc("/search", productHandler)
	log.Println("Server is starting...")
	http.ListenAndServe(":8080", nil)
}

func productHandler(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()

	productName, found := values["product"]
	if found {
		products := scrapper.WebsiteScrapperSS(productName[0])
		output, err := json.Marshal(products)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
		w.Write(output)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}
