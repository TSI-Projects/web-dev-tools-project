package server

import (
	"log"
	"net/http"
	"os"

	"github.com/AndrejsPon00/web-dev-tools/backend/module"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
)

func Start() {
	r := mux.NewRouter()
	r.HandleFunc("/posts/search", sseMiddleware(productHandler)).Methods(http.MethodGet)
	r.HandleFunc("/sources", sourcesHandler).Methods(http.MethodGet)

	log.Println("Server is starting...")
	log.Fatal(http.ListenAndServe(":8080", getCORSHandler(r)))
}

func sourcesHandler(w http.ResponseWriter, r *http.Request) {
	response := &module.Response{Sources: module.EVERY_SOURCE}
	byteResponse, err := toByteArray(response)
	if err != nil {
		http.Error(w, "failed to decode sources to response", http.StatusInternalServerError)
	}
	w.Write(byteResponse)
}

func productHandler(w http.ResponseWriter, r *http.Request) {
	params := &module.URLParams{Filter: &module.Filter{}}
	decoder := schema.NewDecoder()
	if err := decoder.Decode(params, r.Form); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	handler := NewHandler()
	handler.SetWriter(w)
	handler.SetParams(params)

	handler.GetScraper().WG.Add(1)
	go handler.GetScraper().ScrapPosts()

	handler.AddWaitGroup(2)
	go handler.SetupErrorChannel()
	go handler.SetupResultChannel()

	handler.GetScraper().WG.Wait()
	handler.Wait()
	handler.Clear()
}

func sseMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")

		if err := r.ParseForm(); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		next(w, r)
	}
}

func getCORSHandler(r *mux.Router) http.Handler {
	headersOk := handlers.AllowedHeaders([]string{"*"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"*"})

	return handlers.CORS(originsOk, headersOk, methodsOk)(r)
}

func getPort() string {
	port := os.Getenv(module.ENV_VAR_PORT)
	if port == "" {
		return module.DEFAULT_PORT
	}
	return port
}
