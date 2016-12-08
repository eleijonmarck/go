package api

import (
	"encoding/json"
	"net/http"

	"fmt"
	"github.com/eleijonmarck/codeshopping/cart"
)

func GetAllCarts(cr cart.Repository) http.Handler {
	type ret struct {
		Carts []byte `json:"carts"`
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		allCarts := cr.FindAll()
		byteCart, _ := json.Marshal(&allCarts)
		if err2 := json.NewEncoder(w).Encode(ret{Carts: byteCart}); err2 != nil {
			w.Write([]byte(fmt.Sprintf(`{"error marshal": "%s"}`, err2.Error())))
		}
		w.Write(byteCart)
	})
}
