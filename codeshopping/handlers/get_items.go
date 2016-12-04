package handlers

import (
	"encoding/json"
	"net/http"

	"fmt"
	"github.com/garyburd/redigo/redis"
	"log"
	"strings"
)

func GetAllItems(pool *redis.Pool) http.Handler {
	type ret struct {
		Carts []byte `json:"carts"`
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c := pool.Get()
		key := r.URL.Query().Get("key")
		if key == "" {
			http.Error(w, "missing name in query string", http.StatusBadRequest)
			return
		}
		keys, _ := redis.Strings(c.Do("KEYS", "test*"))
		fmt.Println(keys)
		log.Println(keys)
		val, err := redis.Bytes(c.Do("GET", "test1"))
		if err != nil {
			// err handling
			panic(err)
		}
		if err := json.NewEncoder(w).Encode(ret{Carts: val}); err != nil {
			// errrroror
		}
		w.Write([]byte(strings.Join(keys, ", ")))
	})
}
