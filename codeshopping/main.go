package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"

	"github.com/eleijonmarck/codeshopping/cart"
	"github.com/eleijonmarck/codeshopping/handlers"
	"github.com/eleijonmarck/codeshopping/handlers/api"
	"github.com/eleijonmarck/codeshopping/redisdb"
	"github.com/garyburd/redigo/redis"
)

const (
	defaultPort           = "8080"
	defaultRedisURL       = "http://127.0.0.1:6379"
	defaultRedisDBPort    = "6379"
	defaultDBName         = "codeshoppingDB"
	defaultRedisMaxIdle   = 3
	defaultRedisMaxActive = 32
)

func main() {

	// Setup repositories
	var (
		carts cart.Repository
	)

	// Create the logger used by the server
	logger := log.New(os.Stdout, "", 0)

	// Create new Redis Pool
	pool, err := newRedisPool(
		envString("REDISCLOUD_URL", defaultRedisURL),
		envInt("REDIS_MAX_IDLE", defaultRedisMaxIdle),
		envInt("REDIS_MAX_ACTIVE", defaultRedisMaxActive),
	)
	if err != nil {
		logger.Fatal(err)
	}
	defer pool.Close()
	carts, _ = redisdb.NewCartRepository(defaultDBName, pool)

	// creates a http.ServeMux, register handlers to execute in response to routes
	mux := http.NewServeMux()

	// api
	mux.Handle("/carts/", api.GetCart(carts))
	mux.Handle("/carts/create", api.CreateCart(carts))

	// test storage
	storeTestData(carts)

	// handlers
	fmt.Printf("starting server")
	mux.Handle("/", handlers.IndexHandler())
	http.ListenAndServe(":8080", mux)
}

func newRedisPool(addr string, maxIdle, maxActive int) (*redis.Pool, error) {
	url, err := url.Parse(addr)
	if err != nil {
		return nil, err
	}

	return &redis.Pool{
		MaxIdle:     maxIdle,
		MaxActive:   maxActive,
		Wait:        true,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", url.Host)
			if err != nil {
				return nil, err
			}

			if url.User != nil {
				password, _ := url.User.Password()
				_, err := c.Do("AUTH", password)
				if err != nil {
					c.Close()
					return nil, err
				}
			}

			return c, nil
		},
		TestOnBorrow: func(c redis.Conn, _ time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}, nil
}

func envString(env, fallback string) string {
	e := os.Getenv(env)
	if e == "" {
		return fallback
	}
	return e
}

func envInt(env string, fallback int) int {
	e := os.Getenv(env)
	if e == "" {
		return fallback
	}
	eInt, _ := strconv.Atoi(e)
	return eInt
}

func storeTestData(r cart.Repository) {
	test1 := cart.New("test1")
	if err := r.Store(test1); err != nil {
		panic(err)
	}
	fmt.Printf("stored test1")
	log.Print("stored test1")

	//gives error when storing two values at same time?
	test2 := cart.New("test2")
	if err2 := r.Store(test2); err2 != nil {
		panic(err2)
	}
	log.Print("stored test2")
}
