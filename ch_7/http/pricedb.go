package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

func main() {
	var db database
	db.store = make(map[string]dollars)
	db.store["shoes"] = 50
	db.store["socks"] = 5

	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/update", db.update)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database struct {
	store map[string]dollars
	sync.Mutex
}

func (db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db.store {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	db.Lock()
	price, ok := db.store[item]
	db.Unlock()
	if !ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}

func (db database) update(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")

	f, err := strconv.ParseFloat(price, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, "bad request: %q\n", price)
		return
	}
	db.Lock()
	db.store[item] = dollars(f)
	db.Unlock()

	fmt.Fprintf(w, "%s -> %s\n", item, dollars(f))
}
