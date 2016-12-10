package subscriptions

import (
	"net/http"
	"fmt"
	"github.com/gorilla/mux"
)

func handleGetSubscriptions(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprintf(w, "Hello %q", vars["userId"]);
}
