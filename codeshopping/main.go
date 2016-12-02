package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/eleijonmarck/codeshopping/cart"
	"github.com/eleijonmarck/codeshopping/handlers"
	"github.com/eleijonmarck/codeshopping/redisdb"
	"github.com/garyburd/redigo/redis"
)

const (
	defaultPort        = "8080"
	defaultRedisDBUrl  = "127.0.0.1"
	defaultRedisDBPort = "6379"
	defaultDBName      = "codeshoppingDB"
)

func main() {

	// Setup repositories
	var (
		carts cart.Repository
	)
	conn, err := redis.Dial("tcp", ":6379")
	if err != nil {
		// handle connection error
		panic(err)
	}
	defer conn.Close()
	carts, _ = redisdb.NewCartRepository(defaultDBName, conn)

	// creates a http.ServeMux, used to register handlers to execute in
	// response to routes
	mux := http.NewServeMux()

	// get the items of the database
	mux.Handle("/products/api/", handlers.GetAllItems(conn))

	// test storage
	storeTestData(carts)

	// apis
	fmt.Printf("starting server")
	mux.Handle("/", handlers.IndexHandler())
	// mux.Handle("/products", handlers.ProductHandler())
	http.ListenAndServe(":8080", mux)

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
