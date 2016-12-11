package api

import (
	"net/http"

	"encoding/json"
	"github.com/eleijonmarck/codeshopping/cart"
)

// CreateCart will create a item in the store
func CreateCart(cr cart.Repository) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		key := r.URL.Query().Get("key")
		if key == "" {
			http.Error(w, "missing key in query string", http.StatusBadRequest)
			return
		}

		newCart := cart.New(key)
		err := cr.Store(newCart)
		if err != nil {
			//error handling
		}
		val := []byte{}
		err2 := json.Unmarshal(val, newCart)
		if err2 != nil {
			//
		}
		w.WriteHeader(http.StatusCreated)
		w.Write(val)
	})
}
