package subscriptions

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"

	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
)

func handlePostSubscriptions(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	s := &Subscription{}
	json.NewDecoder(r.Body).Decode(s)

	s.SubscriptionId = buildKey(s)
	log.Debugf(ctx, "Generated key: %d", s.SubscriptionId)
	key := datastore.NewKey(ctx, "Subscription", s.SubscriptionId, 0, nil)

	_, err := datastore.Put(ctx, key, s)

	if err != nil {
		log.Errorf(ctx, "Error when writing to datastore %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// TODO: Write Location header
}

func buildKey(s *Subscription) string {
	b := sha256.New()
	b.Write([]byte(s.UserId))
	b.Write([]byte(s.FeedId))

	return hex.EncodeToString(b.Sum(nil))
}
