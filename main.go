package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	v1 "stock/pkg/api/openfoodwatch/v1"
)

const (
	listen = ""
	port   = 8081
)

func main() {

	AddHandlers()
	StartServer(port, listen)
}

func AddHandlers() {
	http.HandleFunc("/categories", func(w http.ResponseWriter, r *http.Request) {

		ofwClient := v1.NewForOpts()
		resp, err := ofwClient.Categories().Get(context.Background())

		if err != nil {
			log.Fatalln(err)
		}

		data, _ := json.Marshal(resp)
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprint(w, string(data))
	})
}

func StartServer(port int, listen string) {
	base := fmt.Sprintf("%s:%d", listen, port)
	fmt.Println(base)

	if err := http.ListenAndServe(base, nil); err != nil {
		log.Fatalln(err)
	}
}
