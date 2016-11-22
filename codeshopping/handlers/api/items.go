package api

import (
	"encoding/json"
	"net/http"

	"github.com/eleijonmarck/codeshopping/cart"
)

func Items(cr *cartRepository) http.Handler {
	type ret struct {
		Items map[string]*CartItem `json:"items"`
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		key := r.URL.Query().Get("key")
		if key == "" {
			// error
			return
		}
		defer r.Body.Close()
		items, err := cr.FindAll()
		if err != nil {
			//error handling
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Wrire(json.UnMarshal(items))
	})
}
