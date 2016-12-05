package api

import (
	"net/http"

	"encoding/json"
	"github.com/eleijonmarck/codeshopping/cart"
)

// CreateCart will create a item in the store
func CreateCart(cr cart.Repository) http.Handler {
	type res struct {
		Carts []byte `json:"carts"`
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		key := r.URL.Query().Get("set")
		if key == "" {
			http.Error(w, "missing name in query string", http.StatusBadRequest)
			return
		}
		// if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		// 	w.WriteHeader(0)
		// 	w.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, err.Error())))
		// }

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
