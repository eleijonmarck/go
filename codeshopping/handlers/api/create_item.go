package api

import (
	"encoding/json"
	"github.com/pborman/uuid"
	"net/http"

	"fmt"
	"github.com/eleijonmarck/codeshopping/cart"
	"strings"
)

// CreateItem will create a item in the store
func CreateItem(cr cart.Repository) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := cart.Cart{}
		if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
			w.WriteHeader(0)
			w.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, err.Error())))
		}

		in.CartID = strings.Split(strings.ToUpper(uuid.New()), "-")[0]
		err := cr.Store(&in)
		if err != nil {
			//error handling
		}
		w.WriteHeader(http.StatusCreated)
	})
}
