package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/jellydator/ttlcache/v2"
	"log"
	"net/http"
	v1 "stock/pkg/api/openfoodwatch/v1"
	"stock/pkg/rest"
	"strings"
	"time"
)

const (
	listen     = ""
	port       = 8081
	offBaseUrl = "https://de.openfoodfacts.org"
	ttl        = 336 * time.Hour
)

var cache ttlcache.SimpleCache = ttlcache.NewCache()

func main() {
	cache.SetTTL(ttl)

	AddHandlers()
	StartServer(port, listen)
}

func AddHandlers() {
	offClient := v1.NewForOpts(rest.WithBaseURL(offBaseUrl), rest.WithCache(cache))

	http.HandleFunc("/categories", func(w http.ResponseWriter, r *http.Request) {
		resp, err := offClient.Categories().List(context.Background())
		HandleError(w, err)
		WriteJsonSuccess(w, resp)
	})

	http.HandleFunc("/category/", func(w http.ResponseWriter, r *http.Request) {
		id := strings.TrimPrefix(r.URL.Path, "/category/")
		//page := r.URL.Query().Get("page")
		resp, err := offClient.Products().GetForCategory(context.Background(), id)
		HandleError(w, err)
		WriteJsonSuccess(w, resp)
	})

	http.HandleFunc("/products/", func(w http.ResponseWriter, r *http.Request) {
		id := strings.TrimPrefix(r.URL.Path, "/products/")
		resp, err := offClient.Products().Get(context.Background(), id)
		HandleError(w, err)
		WriteJsonSuccess(w, resp)
	})
}

func StartServer(port int, listen string) {
	base := fmt.Sprintf("%s:%d", listen, port)
	fmt.Println(base)

	if err := http.ListenAndServe(base, nil); err != nil {
		log.Fatalln(err)
	}
}

func WriteJsonSuccess(w http.ResponseWriter, response interface{}) {
	WriteJson(w, response)
}

func WriteJson(w http.ResponseWriter, response interface{}) {
	data, _ := json.Marshal(response)
	w.Header().Add("Content-Type", "application/json")
	fmt.Fprint(w, string(data))
}

func HandleError(w http.ResponseWriter, err error) {
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
	}
}
