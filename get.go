package subscriptions

import (
	"encoding/json"

	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
)

func handleGetSubscriptions(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	feedId := r.FormValue("feed_id")
	if feedId == "" {
		http.Error(w, "Needs feed_id to find subscriptions", http.StatusBadRequest)
		return
	}

	query := datastore.NewQuery("Subscription").
		Filter("FeedId = ", feedId)

	s := []Subscription{}

	_, err := query.GetAll(ctx, &s)

	if err != nil {
		log.Errorf(ctx, "Error when fetching subscription data %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).
		Encode(s)
}
