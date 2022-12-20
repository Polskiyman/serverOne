package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

var (
	counter int
	port1   = "8080"
	port2   = "8081"
)

func main() {
	if len(os.Args) > 2 {
		port1, port2 = os.Args[1], os.Args[2]
	}

	http.HandleFunc("/", handleProxy)
	err := http.ListenAndServe("localhost:8090", nil)
	if err != nil {
		fmt.Printf("can't start Proxy: %s\n", err.Error())
	}
}

func handleProxy(w http.ResponseWriter, r *http.Request) {
	port := port1
	if counter%2 == 0 {
		port = port2
	}
	addr, err := url.Parse(fmt.Sprintf("http://localhost:%s", port))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	proxy := httputil.NewSingleHostReverseProxy(addr)

	log.Printf("%s %s%s:%s", r.Method, r.Host, r.URL, port)

	proxy.ServeHTTP(w, r)
	counter++
}
