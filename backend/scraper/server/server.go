package server

import (
	"log"
	"net/http"

	"github.com/goccy/go-json"

	"github.com/AndrejsPon00/web-dev-tools/backend/scrapper"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func Start() {
	r := mux.NewRouter()
	r.HandleFunc("/search", productHandler).Methods(http.MethodGet)

	log.Println("Server is starting...")
	headersOk := handlers.AllowedHeaders([]string{"*"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"*"})

	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(originsOk, headersOk, methodsOk)(r)))
}

func productHandler(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	productName, found := values["product"]
	if found {
		client := scrapper.NewScraper(productName[0], w)
		posts := client.ScrapPosts()
		output, err := json.Marshal(posts)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
		w.Write(output)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}
