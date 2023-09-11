package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/AndrejsPon00/web-dev-tools/backend/scrapper"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func Start() {
	r := mux.NewRouter()
	r.HandleFunc("/search", productHandler).Methods(http.MethodGet)

	log.Println("Server is starting...")
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{http.MethodPost, http.MethodGet})

	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(originsOk, headersOk, methodsOk)(r)))
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
