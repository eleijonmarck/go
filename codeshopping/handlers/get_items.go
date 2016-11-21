package handlers

import (
	"fmt"
	"net/http"

	"github.com/garyburd/redigo/redis"
)

func GetAllItems(db redis.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		key := r.URL.Query().Get("key")
		if key == "" {
			http.Error(w, "missing name in query string", http.StatusBadRequest)
			return
		}
		val, err := db.Do("GET", key)
		w.WriteHeader(http.StatusOK)
		w.Write(val)
	})
}
