package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"strconv"
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

func buyHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/buy/"):]
	i, _ := loadproduct(title)
	t, _ := template.ParseFiles("buy.html")
	t.Execute(w, i)
}

func main() {
	fmt.Println("running the server")
	http.HandleFunc("/products", productHandler)
	http.HandleFunc("/buy", buyHandler)
	// http.HandleFunc("/", http.HandlerFunc(indexHandler))
	http.ListenAndServe(":8080", nil)
	// http.ListenAndServe(":8080", nil)
}
