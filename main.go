package subscriptions

import (
	"github.com/gorilla/mux"
	"net/http"
)

func Run() {
	router := mux.NewRouter()

	router.HandleFunc("/subscriptions", handleGetSubscriptions).
		Methods(http.MethodGet)

	router.HandleFunc("/subscriptions", handlePostSubscriptions).
		Methods(http.MethodPost)

	http.Handle("/", router)
}