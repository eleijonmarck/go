package api

import (
	"encoding/json"
	"net/http"

	"github.com/eleijonmarck/codeshopping/cart"
	"log"
)

func Items(cr cart.Repository) http.Handler {
	type ret struct {
		Items map[string]*cart.CartItem `json:"items"`
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		key := r.URL.Query().Get("key")
		if key == "" {
			// error
			return
		}
		defer r.Body.Close()
		items := cr.FindAll()

		jsonitems, err := json.Marshal(items)
		if err != nil {
			// error handling
		}
		log.Printf("api/items, called")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonitems)
	})
}
