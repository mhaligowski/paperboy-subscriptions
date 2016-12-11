package subscriptions

import (
	"github.com/gorilla/mux"
	"net/http"
)

func init() {
	router := mux.NewRouter()

	router.HandleFunc("/subscriptions", handleGetSubscriptions).
		Methods(http.MethodGet)

	http.Handle("/", router)
}