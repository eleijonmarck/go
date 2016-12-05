package api

import (
	"encoding/json"
	"net/http"

	"github.com/garyburd/redigo/redis"
	"strings"
)

func GetACart(pool *redis.Pool) http.Handler {
	type ret struct {
		Carts []byte `json:"carts"`
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c := pool.Get()
		key := r.URL.Query().Get("get")
		if key == "" {
			http.Error(w, "missing name in query string", http.StatusBadRequest)
			return
		}
		keys, _ := redis.Strings(c.Do("KEYS", "test*"))
		val, err := redis.Bytes(c.Do("GET", "test1"))
		if err != nil {
			// err handling
			panic(err)
		}
		if err := json.NewEncoder(w).Encode(ret{Carts: val}); err != nil {
			// errrroror
		}
		w.Write([]byte(strings.Join(keys, ", ")))
		w.Write(val)
	})
}

func getJson(url string, target interface{}) error {
	r, err := http.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}
