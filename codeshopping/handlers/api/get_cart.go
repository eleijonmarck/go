package api

import (
	"encoding/json"
	"net/http"

	"fmt"
	"github.com/eleijonmarck/codeshopping/cart"
	"strings"
)

type ret struct {
	Carts []byte `json:"carts"`
}

func GetCart(cr cart.Repository) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := strings.TrimPrefix(r.URL.Path, "/carts/")

		if id == "" {
			allCarts := cr.FindAll()
			byteCart, _ := json.Marshal(&allCarts)
			if err2 := json.NewEncoder(w).Encode(ret{Carts: byteCart}); err2 != nil {
				w.Write([]byte(fmt.Sprintf(`{"error marshal": "%s"}`, err2.Error())))
			}
			w.Write(byteCart)
		}

		if id != "" {
			foundCart, err := cr.Find(id)
			if err != nil {
				w.Write([]byte(fmt.Sprintf(`{"error finding": "%s"}`, err.Error())))
			}
			byteCart, _ := json.Marshal(&foundCart)
			if err2 := json.NewEncoder(w).Encode(ret{Carts: byteCart}); err2 != nil {
				w.Write([]byte(fmt.Sprintf(`{"error marshal": "%s"}`, err2.Error())))
			}
			w.Write(byteCart)
		}

	})
}
