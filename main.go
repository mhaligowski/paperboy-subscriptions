package subscriptions

import (
	"github.com/gorilla/mux"
	"net/http"
	"github.com/rs/cors"
)

func Run() {
	router := mux.NewRouter()

	router.HandleFunc("/subscriptions", handleGetSubscriptions).
		Methods(http.MethodGet)

	router.HandleFunc("/subscriptions", handlePostSubscriptions).
		Methods(http.MethodPost)

	http.Handle("/", cors.Default().Handler(router))
}