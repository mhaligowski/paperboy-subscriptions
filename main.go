package subscriptions

import (
	"github.com/gorilla/mux"
	"net/http"
)

func init() {
	router := mux.NewRouter()

	router.HandleFunc("/users/{userId}/subscriptions", handleGetSubscriptions).Methods(http.MethodGet)

	http.Handle("/", router)
}