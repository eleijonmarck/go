package api

import (
	"encoding/json"
	"net/http"

	"fmt"
	"github.com/eleijonmarck/codeshopping/cart"
)

func GetACart(cr cart.Repository) http.Handler {
	type ret struct {
		Carts []byte `json:"carts"`
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		key := r.URL.Query().Get("get")
		if key == "" {
			http.Error(w, "missing name in query string", http.StatusBadRequest)
			return
		}
		fmt.Println(key)
		foundCart, err := cr.Find(key)
		if err != nil {
			w.Write([]byte(fmt.Sprintf(`{"error finding": "%s"}`, err.Error())))
		}
		byteCart, _ := json.Marshal(&foundCart)
		if err2 := json.NewEncoder(w).Encode(ret{Carts: byteCart}); err2 != nil {
			w.Write([]byte(fmt.Sprintf(`{"error marshal": "%s"}`, err2.Error())))
		}
		w.Write(byteCart)
	})
}
