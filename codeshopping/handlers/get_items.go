package handlers

import (
	"net/http"

	"github.com/garyburd/redigo/redis"
)

func GetAllItems(db redis.Conn) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		key := r.URL.Query().Get("key")
		if key == "" {
			http.Error(w, "missing name in query string", http.StatusBadRequest)
			return
		}
		val, err := redis.Bytes(db.Do("GET", key))
		if err != nil {
			// err handling
		}
		w.WriteHeader(http.StatusOK)
		w.Write(val)
	})
}
