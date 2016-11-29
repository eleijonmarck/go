package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/eleijonmarck/codeshopping/cart"
	"github.com/eleijonmarck/codeshopping/handlers"
	"github.com/eleijonmarck/codeshopping/redisdb"
	"github.com/garyburd/redigo/redis"
)

type Page struct {
	Title    string
	Body     []byte
	Products product
}

type product struct {
	Name  string
	Price int
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

// TODO: load from database instead of text files
func loadproduct(name string) (*product, error) {
	filename := "/products/" + name + ".txt"
	price, err := ioutil.ReadFile(filename)
	intPrice, _ := strconv.Atoi(string(price))
	if err != nil {
		return nil, err
	}
	return &product{Name: name, Price: intPrice}, nil
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tplValues := map[string]interface{}{"Header": "Home"}
	t, err := template.ParseFiles("index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, tplValues)
}

func productHandler(w http.ResponseWriter, r *http.Request) {
	product := r.URL.Path[len("/product/"):]
	p, _ := loadPage(product)
	t, _ := template.ParseFiles("product.html")
	t.Execute(w, p)
}

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
	mux.Handle("/products", handlers.GetAllItems(conn))

	// apis
	// mux.Handle("/api/items", api.Items(&carts))
	fmt.Printf("starting server")
	http.ListenAndServe(":8080", mux)

	storeTestData(carts)
}

func storeTestData(r cart.Repository) {
	test1 := cart.New("test1")
	if err := r.Store(test1); err != nil {
		panic(err)
	}
	fmt.Printf("stored test1")
	log.Print("stored test1")

	test2 := cart.New("test2")
	if err := r.Store(test2); err != nil {
		panic(err)
	}
	log.Print("stored test2")
}
