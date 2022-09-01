package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	v1 "stock/pkg/api/openfoodwatch/v1"
	"stock/pkg/rest"
	"strings"
)

const (
	listen     = ""
	port       = 8081
	ofwBaseUrl = "https://de.openfoodfacts.org"
)

func main() {

	AddHandlers()
	StartServer(port, listen)
}

// https://de.openfoodfacts.org/api/v2/product/8076800105988
func AddHandlers() {
	ofwClient := v1.NewForOpts(rest.WithBaseURL(ofwBaseUrl))

	http.HandleFunc("/categories", func(w http.ResponseWriter, r *http.Request) {
		resp, err := ofwClient.Categories().List(context.Background())

		if err != nil {
			log.Fatalln(err)
		}
		WriteJsonSuccess(w, resp)
	})

	http.HandleFunc("/category/", func(w http.ResponseWriter, r *http.Request) {
		id := strings.TrimPrefix(r.URL.Path, "/category/")
		resp, err := ofwClient.Products().GetForCategory(context.Background(), id)

		if err != nil {
			log.Fatalln(err)
		}
		WriteJsonSuccess(w, resp)
	})

	http.HandleFunc("/products/", func(w http.ResponseWriter, r *http.Request) {
		id := strings.TrimPrefix(r.URL.Path, "/products/")
		resp, err := ofwClient.Products().Get(context.Background(), id)

		if err != nil {
			log.Fatalln(err)
		}
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
