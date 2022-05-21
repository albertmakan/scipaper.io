package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	microservices := map[string]int {"/users/":8000, "/papers/":8001, "/library/":8002}
	for path, port := range microservices {
		target, err := url.Parse(fmt.Sprintf("http://localhost:%v", port))
		if err != nil { log.Fatal(err) }
		http.Handle(path, http.StripPrefix(path, httputil.NewSingleHostReverseProxy(target)))
	}
	log.Fatal(http.ListenAndServe(":9095", nil))
}