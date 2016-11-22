package api

import (
	"encoding/json"
	"net/http"

	"github.com/eleijonmarck/codeshopping/cart"
	"github.com/pborman/uuid"
)

// CreateItem will create a item in the store
func CreateItem(cr *cartRepository) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := CartItem{}
		if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
			w.WriteHeader()
			w.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, err.Error())))
		}

		in.ID = strings.Split(strings.ToUpper(uuid.New()), "-")[0]
		err := cr.Store(&in)
		if err != nil {
			//error handling
		}
		w.WriteHeader(http.StatusCreated)
	})
}
