package main

import (
	"net/http"
	"fmt"
)

type MyMux struct {
}

// 简易的路由器
func (p *MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		sayhelloName1(w, r)
		return
	}
	http.NotFound(w, r)
	return
}

func sayhelloName1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello myroute!")
}

func main() {
	mux := &MyMux{}
	http.ListenAndServe(":9090", mux)
}
