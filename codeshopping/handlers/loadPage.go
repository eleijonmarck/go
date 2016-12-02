package handlers

import (
	"html/template"
	"io/ioutil"
	"net/http"
	"strconv"
)

type page struct {
	Title    string
	Body     []byte
	Products product
}

type product struct {
	Name  string
	Price int
}

func (p *page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &page{Title: title, Body: body}, nil
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

func IndexHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tplValues := map[string]interface{}{"Header": "Home"}
		t, err := template.ParseFiles("index.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		t.Execute(w, tplValues)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("hej"))
	})
}

func ProductHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		product := r.URL.Path[len("/product/"):]
		p, _ := loadPage(product)
		t, _ := template.ParseFiles("product.html")
		t.Execute(w, p)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("hej"))
	})
}
